package admin

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	config "personal-website-golang/service/internal/config/admin"
	"personal-website-golang/service/internal/thirdparty/logger"
)

func NewAdminRestService(in restServiceIn) IService {
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

	// admin api
	LoginCtrl admin.IAuthCtrl
}

func (s *restService) Run(ctx context.Context) {
	engine := gin.New()

	s.setRoutes(engine)

	addr := s.in.AppConf.GetGinConfig().AdminAddress
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
