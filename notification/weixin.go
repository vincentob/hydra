package notification

import (
	"fmt"

	"github.com/dantin-s/hydra/httpclient"
	"github.com/pkg/errors"
)

const (
	WXTokenUrl = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
	WXMsgUrl   = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s"
)

// WXAccessTokenBody response body for get access token
type WXAccessTokenBody struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// WXMsgBody the request body we send to weixin api gateway
type WXMsgBody struct {
	ToTag   string    `json:"totag"`
	MsgType string    `json:"msgtype"`
	AgentId string    `json:"agentid"`
	Text    WXMsgText `json:"text"`
	News    WXMsgNews `json:"news"`
}

// WXMsgText for wx text body
type WXMsgText struct {
	Content string `json:"content"`
}

// WXMsgNews for wx news type body
type WXMsgNews struct {
	Articles []WXMsgArticle `json:"articles"`
}

// WXMsgArticle for article type body
type WXMsgArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"picurl"`
}

type WXResponse struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	Invaliduser string `json:"invaliduser"`
}

// SendWX
func SendWX(content *WXMsgBody, agentID string, agentSecret string) error {
	// get accesstoken first
	accessToken := &WXAccessTokenBody{}
	tokenURL := fmt.Sprintf(WXTokenUrl, agentID, agentSecret)
	if _, err := httpclient.DoGetJson(tokenURL, nil, accessToken); err != nil {
		return errors.Wrap(err, "get accesstoken failed")
	}

	if accessToken.ErrCode != 0 {
		return errors.Wrap(errors.New(accessToken.ErrMsg), "get accessToken failed")
	}

	msgURL := fmt.Sprintf(WXMsgUrl, accessToken.AccessToken)
	wxResult := &WXResponse{}
	if _, err := httpclient.DoPostJson(msgURL, nil, content, wxResult); err != nil {
		return errors.Wrap(errors.New(wxResult.ErrMsg), "send wx failed")
	}

	return nil
}
