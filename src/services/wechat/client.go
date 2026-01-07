package wechat

import (
	"encoding/json"
	"fmt"

	"github.com/springmove/sptty"
	"github.com/springmove/tp/src/base"
	"gopkg.in/resty.v1"
)

func createWechatClient(cfg *WechatConfig) *wechatClient {
	return &wechatClient{
		cfg: cfg,
	}
}

type wechatClient struct {
	base.IWechatClient

	http *resty.Client

	cfg *WechatConfig
}

func (s *wechatClient) init() error {

	s.http = sptty.CreateHttpClient()
	return nil
}

func (s *wechatClient) Code2Session(req *base.ReqCode2Session) (*base.RespCode2Session, error) {

	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appID=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		s.cfg.AppID,
		s.cfg.Secret,
		req.Code)

	resp, err := s.http.R().Get(url)
	if err != nil {
		return nil, err
	}

	respBody := base.RespCode2Session{}
	if err := json.Unmarshal(resp.Body(), &respBody); err != nil {
		return nil, err
	}

	if respBody.ErrCode != base.WxOK {
		return nil, fmt.Errorf("ErrCode: %d, ErrMsg: %s", respBody.ErrCode, respBody.ErrMsg)
	}

	return &respBody, nil
}

func (s *wechatClient) GetMobileByAuthCode(authCodeMobile string) (string, error) {
	if authCodeMobile == "" {
		return "", nil
	}

	accessToken, err := s.getAccessToken()
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s", accessToken.Token)
	resp, err := s.http.R().SetBody(&ReqMiniProgramAuthMobile{
		Code: authCodeMobile,
	}).Post(url)
	if err != nil {
		return "", err
	}

	respBody := RespMiniProgramAuthMobile{}
	if err = json.Unmarshal(resp.Body(), &respBody); err != nil {
		return "", err
	}

	if respBody.ErrCode != base.WxOK {
		return "", fmt.Errorf("ErrCode: %d, ErrMsg: %s", respBody.ErrCode, respBody.ErrMsg)
	}

	return respBody.Mobile.Mobile, nil
}

func (s *wechatClient) getAccessToken() (*RespAccessToken, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
		s.cfg.AppID,
		s.cfg.Secret)

	resp, err := s.http.R().Get(url)
	if err != nil {
		return nil, err
	}

	rt := RespAccessToken{}
	if err = json.Unmarshal(resp.Body(), &rt); err != nil {
		return nil, err
	}

	if rt.ErrCode != base.WxOK {
		return nil, fmt.Errorf("ErrCode: %d, ErrMsg: %s", rt.ErrCode, rt.ErrMsg)
	}

	return &rt, nil
}
