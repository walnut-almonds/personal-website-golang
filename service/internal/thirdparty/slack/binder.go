package slack

import (
	"go.uber.org/dig"

	adminConfig "personal-website-golang/service/internal/config/admin"
	webConfig "personal-website-golang/service/internal/config/web"
)

func NewAdmin(in digInAdmin) ISlackNotification {
	opsConf := in.OpsConf.GetOpsSlackConfig()
	return &slackNotification{
		channelID: opsConf.ChannelID,
		botName:   opsConf.BotName,
		token:     opsConf.Token,
	}
}

type digInAdmin struct {
	dig.In

	OpsConf adminConfig.IOpsConfig
}

func NewWeb(in digInWeb) ISlackNotification {
	opsConf := in.OpsConf.GetOpsSlackConfig()
	return &slackNotification{
		channelID: opsConf.ChannelID,
		botName:   opsConf.BotName,
		token:     opsConf.Token,
	}
}

type digInWeb struct {
	dig.In

	OpsConf webConfig.IOpsConfig
}
