package web

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"personal-website-golang/service/internal/thirdparty/logger"
)

func NewWebRestService(in restServiceIn) IService {
	self := &restService{
		in: in,
	}

	return self
}

type IService interface {
	Run(ctx context.Context)
}

type restService struct {
	in restServiceIn
}

type restServiceIn struct {
	dig.In

	AppConf   config.IAppConfig
	SysLogger logger.ILogger `name:"sysLogger"`

	// web api
	LoginCtrl web.IAuthCtrl
}

func (s *restService) Run(ctx context.Context) {
	engine := gin.New()

	s.setRoutes(engine)

	addr := s.in.AppConf.GetGinConfig().Address
	if err := engine.Run(addr); err != nil {
		s.in.SysLogger.Panic(nil, err)
	}
}

func (s *restService) setRoutes(engine *gin.Engine) {
	engine.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	// 設定路由
	s.setPublicRoutes(engine)  // 如：DepositUseCase API, Web API
	s.setPrivateRoutes(engine) // 如：pprof, health
}

func (s *restService) setPublicRoutes(engine *gin.Engine) {
	privateRouteGroup := engine.Group("")

	// 設定路由
	s.setApiRouters(privateRouteGroup)
}

func (s *restService) setPrivateRoutes(engine *gin.Engine) {
	privateRouteGroup := engine.Group("/_")

	// health check
	privateRouteGroup.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
}
