package logger

import (
	"go.uber.org/dig"

	adminConfig "personal-website-golang/service/internal/config/admin"
	webConfig "personal-website-golang/service/internal/config/web"
)

func NewAdminAppLogger(in digInAdmin) ILogger {
	return NewAppLogger(in.AppConf.GetLogConfig().Level, in.AppConf.GetLogConfig().Name, in.AppConf.GetLogConfig().Env)
}

func NewAdminSysLogger(in digInAdmin) ILogger {
	return NewSysLogger(in.AppConf.GetLogConfig().Level, in.AppConf.GetLogConfig().Name, in.AppConf.GetLogConfig().Env)
}

type digInAdmin struct {
	dig.In

	AppConf adminConfig.IAppConfig
}

func NewWebAppLogger(in digInWeb) ILogger {
	return NewAppLogger(in.AppConf.GetLogConfig().Level, in.AppConf.GetLogConfig().Name, in.AppConf.GetLogConfig().Env)
}

func NewWebSysLogger(in digInWeb) ILogger {
	return NewSysLogger(in.AppConf.GetLogConfig().Level, in.AppConf.GetLogConfig().Name, in.AppConf.GetLogConfig().Env)
}

type digInWeb struct {
	dig.In

	AppConf webConfig.IAppConfig
}
