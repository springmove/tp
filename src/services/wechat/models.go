package wechat

import (
	"fmt"

	"github.com/springmove/tp/src/base"
)

type RespAccessToken struct {
	base.Resp
	Token  string `json:"access_token"`
	Expiry int    `json:"expires_in"`
}

type ReqMiniProgramAuthMobile struct {
	Code string `json:"code"`
}

type RespMiniProgramAuthMobile struct {
	base.Resp

	Mobile *MiniProgramAuthMobile `json:"phone_info"`
}

type MiniProgramAuthMobile struct {
	FullMobile  string `json:"phoneNumber"`
	Mobile      string `json:"purePhoneNumber"`
	CountryCode string `json:"countryCode"`
}

func (s *MiniProgramAuthMobile) ToValidMobile() string {
	return fmt.Sprintf("+%s-%s", fmt.Sprintf("%03s", s.CountryCode), s.Mobile)
}
