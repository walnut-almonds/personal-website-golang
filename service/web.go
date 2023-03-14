package service

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"personal-website-golang/service/internal/app/web"
	binder "personal-website-golang/service/internal/binder/web"
	config "personal-website-golang/service/internal/config/web"
	"personal-website-golang/service/internal/thirdparty/logger"
)

// web
var (
	webApp     *WebApp
	webSetOnce sync.Once
)

type WebApp struct {
	in webAppDigIn
}

type webAppDigIn struct {
	dig.In

	AppConf   config.IAppConfig
	SysLogger logger.ILogger `name:"sysLogger"`

	WebRestService web.IService
}

func RunWebServer() {
	binder := binder.New()
	if err := binder.Invoke(runWebServer); err != nil {
		panic(err)
	}

	select {}
}

func runWebServer(in WebApp) {
	webSetOnce.Do(func() {
		webApp = &in
	})

	ctx := context.Background()

	webApp.serverInterrupt(ctx)
	webApp.ginMode()

	webApp.in.SysLogger.Info(ctx, fmt.Sprintf("[Build Info] %s", getBuildInfo()))

	go webApp.in.WebRestService.Run(ctx)
}

func (app *WebApp) ginMode() {
	gin.DisableConsoleColor()
	if app.in.AppConf.GetGinConfig().DebugMode == false {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

func (app *WebApp) serverInterrupt(ctx context.Context) {
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
