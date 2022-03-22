package notification

import (
	"github.com/vincentob/hydra/httpclient"
	"github.com/pkg/errors"
)

const (
	MsgTypeLink = "link"
	MsgTypeText = "text"
	MsgTypeMD   = "markdown"
)

// DDMsg dingding 告警消息体，默认是 Link 类型
type DDMsg struct {
	Msgtype  string        `json:"msgtype"`
	Link     DDMsgLink     `json:"link,omitempty"`
	Text     DDMsgText     `json:"text,omitempty"`
	Markdown DDMsgMarkDown `json:"markdown,omitempty"`
	At       DDAt          `json:"at,omitempty"`
}

// DDMsgLink link 类型的 msg
type DDMsgLink struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicURL     string `json:"picUrl"`
	MessageURL string `json:"messageUrl"`
}

type DDMsgText struct {
	Content string `json:"content"`
}

type DDMsgMarkDown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type DDAt struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	IsAtAll   bool     `json:"isAtAll,omitempty"`
}

type DDResponse struct {
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

func SendDD(msg *DDMsg, url string) (*DDResponse, error) {
	var dingDingResp DDResponse

	if _, err := httpclient.DoPostJson(url, nil, msg, &dingDingResp); err != nil {
		return nil, errors.Wrap(err, "SendDD failed")
	}

	return &dingDingResp, nil
}
