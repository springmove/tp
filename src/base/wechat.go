package base

const (
	ServiceWechat = "wechat"

	ErrNotSupported = "ErrNotSupported"
)

const (
	WxOK = 0
)

const (
	BotMsgTypeText     = "text"
	BotMsgTypeMarkdown = "markdown"
)

var IWechat IServiceWechat

type IServiceWechat interface {
	Client(index ...int) IWechatClient
}

type IWechatClient interface {
	Code2Session(req *ReqCode2Session) (*RespCode2Session, error)
	GetMobileByAuthCode(authCodeMobile string) (string, error)
	PostGroupBotMsg(req *ReqBotMsg) error
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

type Content struct {
	Content string `json:"content"`
}

type BotMsgContent struct {
	MsgType  string   `json:"msgtype"`
	Text     *Content `json:"text,omitempty"`
	Markdown *Content `json:"markdown,omitempty"`
}

type ReqBotMsg struct {
	BotKey string
	BotMsgContent
}
