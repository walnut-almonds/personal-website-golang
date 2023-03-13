package service

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"personal-website-golang/service/internal/app/admin"
	binder "personal-website-golang/service/internal/binder/admin"
	config "personal-website-golang/service/internal/config/admin"
	"personal-website-golang/service/internal/thirdparty/logger"
)

func RunAdmin() {
	binder := binder.New()
	if err := binder.Invoke(initServer); err != nil {
		panic(err)
	}

	select {}
}

type adminAppDigIn struct {
	dig.In

	AppConf   config.IAppConfig
	SysLogger logger.ILogger `name:"sysLogger"`

	AdminRestService admin.IService
}

func initServer(in digIn) {
	ctx := context.Background()

	serverInterrupt(ctx, in)
	ginMode(in)
	in.SysLogger.Info(ctx, fmt.Sprintf("[Build Info] %s", getBuildInfo()))

	go in.AdminRestService.Run(ctx)
}

func ginMode(in digIn) {
	gin.DisableConsoleColor()
	if in.AppConf.GetGinConfig().DebugMode == false {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

func serverInterrupt(ctx context.Context, in digIn) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	go func() {
		select {
		case c := <-interrupt:
			in.SysLogger.Warn(ctx, fmt.Sprintf("Server Shutdown, osSignal: %v", c))
			os.Exit(0)
		}
	}()
}
