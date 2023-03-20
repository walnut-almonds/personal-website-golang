package mongo

import (
	"go.uber.org/dig"

	adminConfig "personal-website-golang/service/internal/config/admin"
	webConfig "personal-website-golang/service/internal/config/web"
	"personal-website-golang/service/internal/thirdparty/logger"
)

func NewAdminClient(in digInAdmin) IMongoDBClient {
	appConf := in.AppConf.GetMongoConfig()
	opsConf := in.OpsConf.GetOpsMongoConfig()
	return initWithConfig(in.SysLogger, Config{
		Host:     opsConf.Host,
		User:     opsConf.User,
		Password: opsConf.Password,
		DbName:   opsConf.Dbname,

		ReplicaName:     appConf.ReplicaName,
		ReadPreference:  appConf.ReadPreference,
		MaxPoolSize:     appConf.MaxPoolSize,
		MinPoolSize:     appConf.MinPoolSize,
		MaxConnIdleTime: appConf.MaxConnIdleTime,
		MaxStaleness:    appConf.MaxStaleness,
		SlowLogEnable:   appConf.SlowLogEnable,
		SlowLogJudgment: appConf.SlowLogJudgment,
	})
}

type digInAdmin struct {
	dig.In

	AppConf   adminConfig.IAppConfig
	OpsConf   adminConfig.IOpsConfig
	SysLogger logger.ILogger `name:"sysLogger"`
}

func NewWebClient(in digInWeb) IMongoDBClient {
	appConf := in.AppConf.GetMongoConfig()
	opsConf := in.OpsConf.GetOpsMongoConfig()
	return initWithConfig(in.SysLogger, Config{
		Host:     opsConf.Host,
		User:     opsConf.User,
		Password: opsConf.Password,
		DbName:   opsConf.Dbname,

		ReplicaName:     appConf.ReplicaName,
		ReadPreference:  appConf.ReadPreference,
		MaxPoolSize:     appConf.MaxPoolSize,
		MinPoolSize:     appConf.MinPoolSize,
		MaxConnIdleTime: appConf.MaxConnIdleTime,
		MaxStaleness:    appConf.MaxStaleness,
		SlowLogEnable:   appConf.SlowLogEnable,
		SlowLogJudgment: appConf.SlowLogJudgment,
	})
}

type digInWeb struct {
	dig.In

	AppConf   webConfig.IAppConfig
	OpsConf   webConfig.IOpsConfig
	SysLogger logger.ILogger `name:"sysLogger"`
}
