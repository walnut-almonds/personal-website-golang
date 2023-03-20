package mysql

import (
	"go.uber.org/dig"

	adminConfig "personal-website-golang/service/internal/config/admin"
	webConfig "personal-website-golang/service/internal/config/web"
	"personal-website-golang/service/internal/thirdparty/logger"
)

func NewAdminClient(in digInAdmin) IMySQLClient {
	appConf := in.AppConf.GetMySQLConfig()
	opsConf := in.OpsConf.GetOpsMySQLConfig()
	return initWithConfig(in.SysLogger, Config{
		Username: opsConf.Username,
		Password: opsConf.Password,
		Address:  opsConf.Address,
		Database: opsConf.Database,

		LogMode:        appConf.LogMode,
		MaxIdle:        appConf.MaxIdle,
		MaxOpen:        appConf.MaxOpen,
		ConnMaxLifeMin: appConf.ConnMaxLifeMin,
	})
}

type digInAdmin struct {
	dig.In

	AppConf   adminConfig.IAppConfig
	OpsConf   adminConfig.IOpsConfig
	SysLogger logger.ILogger `name:"sysLogger"`
}

func NewWebClient(in digInWeb) IMySQLClient {
	appConf := in.AppConf.GetMySQLConfig()
	opsConf := in.OpsConf.GetOpsMySQLConfig()
	return initWithConfig(in.SysLogger, Config{
		Username: opsConf.Username,
		Password: opsConf.Password,
		Address:  opsConf.Address,
		Database: opsConf.Database,

		LogMode:        appConf.LogMode,
		MaxIdle:        appConf.MaxIdle,
		MaxOpen:        appConf.MaxOpen,
		ConnMaxLifeMin: appConf.ConnMaxLifeMin,
	})
}

type digInWeb struct {
	dig.In

	AppConf   webConfig.IAppConfig
	OpsConf   webConfig.IOpsConfig
	SysLogger logger.ILogger `name:"sysLogger"`
}
