package telegram

import (
	"go.uber.org/dig"

	adminConfig "personal-website-golang/service/internal/config/admin"
	webConfig "personal-website-golang/service/internal/config/web"
)

func NewAdmin(in digInAdmin) ITelegram {
	return NewTelegram(in.OpsConf.GetOpsTelegramConfig().Token)
}

type digInAdmin struct {
	dig.In

	OpsConf adminConfig.IOpsConfig
}

func NewWeb(in digInWeb) ITelegram {
	return NewTelegram(in.OpsConf.GetOpsTelegramConfig().Token)
}

type digInWeb struct {
	dig.In

	OpsConf webConfig.IOpsConfig
}
