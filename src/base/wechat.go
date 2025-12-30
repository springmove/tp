package base

const (
	ServiceWechat = "wechat"

	ErrNotSupported = "ErrNotSupported"
)

const (
	WxOK = 0
)

var IWechat IServiceWechat

type IServiceWechat interface {
	Client(index ...int) IWechatClient
}

type IWechatClient interface {
	Code2Session(req *ReqCode2Session) (*RespCode2Session, error)
}

type Resp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type ReqCode2Session struct {
	Code string
}

type RespCode2Session struct {
	Resp

	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
}
