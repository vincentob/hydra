package notification_test

import (
	"testing"

	"github.com/vincentob/hydra/notification"
	"github.com/sirupsen/logrus"
)

func TestSendDD(t *testing.T) {
	ddMsg := &notification.DDMsg{
		Msgtype: notification.MsgTypeMD,
		Markdown: notification.DDMsgMarkDown{
			Title: "Warning: gitlab project name conflict!",
			Text:  "### xxx",
		},
	}

	robot := "https://oapi.dingtalk.com/robot/send?access_token=422e0abff6d8bb9b049d9e7410e2d3043486d7639004ad9c071f76e010e48748"

	resp, err := notification.SendDD(ddMsg, robot)
	if err != nil {
		logrus.Error(err.Error())
	}
	logrus.Info(resp.Errcode)
	logrus.Info(resp.Errmsg)

}
