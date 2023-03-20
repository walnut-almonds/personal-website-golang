package redis

import (
	"go.uber.org/dig"

	adminConfig "personal-website-golang/service/internal/config/admin"
	webConfig "personal-website-golang/service/internal/config/web"
	"personal-website-golang/service/internal/thirdparty/logger"
)

func NewAdminClient(in digInAdmin) IRedisClient {
	opsConf := in.OpsConf.GetOpsRedisConfig()
	return Init(in.SysLogger, opsConf.Address, opsConf.Password, opsConf.Database)
}

type digInAdmin struct {
	dig.In

	OpsConf   adminConfig.IOpsConfig
	SysLogger logger.ILogger `name:"sysLogger"`
}

func NewWebClient(in digInWeb) IRedisClient {
	opsConf := in.OpsConf.GetOpsRedisConfig()
	return Init(in.SysLogger, opsConf.Address, opsConf.Password, opsConf.Database)
}

type digInWeb struct {
	dig.In

	OpsConf   webConfig.IOpsConfig
	SysLogger logger.ILogger `name:"sysLogger"`
}
