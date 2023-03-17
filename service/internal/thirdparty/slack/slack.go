package slack

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/dig"

	"personal-website-golang/service/internal/thirdparty/logger"
)

const slackPostUrl = "https://slack.com/api/chat.postMessage"

type slackMessage struct {
	Text    string `json:"text"`
	Channel string `json:"channel"`
	BotName string `json:"bot_name"`
}

type ISlackNotification interface {
	Send(s string) error
}

type digIn struct {
	dig.In

	AppConf config.IAppConfig
	Logger  logger.ILogger `name:"appLogger"`
}

func NewSlackNotification(in digIn) ISlackNotification {
	return &slackNotification{digIn: in}
}

type slackNotification struct {
	digIn
}

func (sn *slackNotification) Send(s string) error {
	message := &slackMessage{
		Text:    s,
		Channel: sn.AppConf.GetSlackConfig().ChannelID,
		BotName: sn.AppConf.GetSlackConfig().BotName,
	}

	b, err := json.Marshal(message)
	if err != nil {
		sn.digIn.Logger.Error(context.Background(), err)
		return err
	}

	req, err := http.NewRequest(http.MethodPost, slackPostUrl, bytes.NewBuffer(b))
	if err != nil {
		sn.digIn.Logger.Error(context.Background(), err)
		return err
	}

	req.Header.Set("Authorization", "Bearer "+sn.AppConf.GetSlackConfig().Token)
	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		sn.digIn.Logger.Error(context.Background(), err)
		return err
	}

	defer response.Body.Close()

	return nil
}
