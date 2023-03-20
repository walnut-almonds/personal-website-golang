package slack

import (
	"bytes"
	"encoding/json"
	"net/http"
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

type slackNotification struct {
	channelID string
	botName   string
	token     string
}

func (sn *slackNotification) Send(s string) error {
	message := &slackMessage{
		Text:    s,
		Channel: sn.channelID,
		BotName: sn.botName,
	}

	b, err := json.Marshal(message)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, slackPostUrl, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+sn.token)
	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	return nil
}
