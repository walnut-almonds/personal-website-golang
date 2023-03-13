package service

import (
	"personal-website-golang/service/internal/app/admin"
	binder "personal-website-golang/service/internal/binder/admin"
	config "personal-website-golang/service/internal/config/admin"
	"personal-website-golang/service/internal/thirdparty/logger"

	"go.uber.org/dig"
)

func RunWeb() {
	binder := binder.New()
	if err := binder.Invoke(initServer); err != nil {
		panic(err)
	}

	select {}
}

type WebAppDigIn struct {
	dig.In

	AppConf   config.IAppConfig
	SysLogger logger.ILogger `name:"sysLogger"`

	AdminRestService admin.IService
}
