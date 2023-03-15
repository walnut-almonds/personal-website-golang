package mysql

import (
	"go.uber.org/dig"

	"personal-website-golang/service/internal/thirdparty/logger"
)

func NewDBClient(in digIn) IMySQLClient {
	return initWithConfig(in)
}

type digIn struct {
	dig.In

	AppConf   config.IAppConfig
	OpsConf   config.IOpsConfig
	SysLogger logger.ILogger `name:"sysLogger"`
}
