package service

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"personal-website-golang/service/internal/app/admin"
	binder "personal-website-golang/service/internal/binder/admin"
	config "personal-website-golang/service/internal/config/admin"
	"personal-website-golang/service/internal/thirdparty/logger"
)

// admin
var (
	adminApp     *AdminApp
	adminSetOnce sync.Once
)

type AdminApp struct {
	in adminAppDigIn
}

type adminAppDigIn struct {
	dig.In

	AppConf   config.IAppConfig
	SysLogger logger.ILogger `name:"sysLogger"`

	AdminRestService admin.IService
}

func RunAdminServer() {
	binder := binder.New()
	if err := binder.Invoke(runAdminServer); err != nil {
		panic(err)
	}

	select {}
}

func runAdminServer(in AdminApp) {
	adminSetOnce.Do(func() {
		adminApp = &in
	})

	ctx := context.Background()

	adminApp.serverInterrupt(ctx)
	adminApp.ginMode()

	adminApp.in.SysLogger.Info(ctx, fmt.Sprintf("[Build Info] %s", getBuildInfo()))

	go adminApp.in.AdminRestService.Run(ctx)
}

func (app *AdminApp) ginMode() {
	gin.DisableConsoleColor()
	if app.in.AppConf.GetGinConfig().DebugMode == false {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

func (app *AdminApp) serverInterrupt(ctx context.Context) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	go func() {
		select {
		case c := <-interrupt:
			app.in.SysLogger.Warn(ctx, fmt.Sprintf("Server Shutdown, osSignal: %v", c))
			os.Exit(0)
		}
	}()
}
