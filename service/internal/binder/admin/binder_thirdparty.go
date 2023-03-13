package admin

import (
	"personal-website-golang/service/internal/thirdparty/logger"

	"go.uber.org/dig"
)

func provideThirdParty(binder *dig.Container) {
	if err := binder.Provide(logger.NewAppLogger, dig.Name("appLogger")); err != nil {
		panic(err)
	}

	if err := binder.Provide(logger.NewSysLogger, dig.Name("sysLogger")); err != nil {
		panic(err)
	}
}
