package user_admin

import (
	"go.uber.org/dig"
	adminConfig "personal-website-golang/service/internal/config/admin"
	"personal-website-golang/service/internal/thirdparty/mysql"
)

func NewUserAdmin(in digIn) digOut {
	return digOut{}
}

type digIn struct {
	dig.In

	OpsConf adminConfig.IOpsConfig

	Mysql mysql.IMySQLClient
}

type digOut struct {
	dig.Out
}
