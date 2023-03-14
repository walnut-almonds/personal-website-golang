package admin

import "github.com/gin-gonic/gin"

func (s *restService) setApiRouters(parentRouteGroup *gin.RouterGroup) {
	privateRouteGroup := parentRouteGroup.Group("/admin/v1")

	s.setLoginAPIRoutes(privateRouteGroup)
}

func (s *restService) setLoginAPIRoutes(parentRouteGroup *gin.RouterGroup) {
	anonymous := parentRouteGroup.Group("")
	anonymous.POST("/login", s.in.LoginCtrl.AddSession)

	authRouteGroup := parentRouteGroup.Group("")
	authRouteGroup.Use(s.in.AuthTokenHttpMiddleware.MiddlewareAuth)

	authRouteGroup.POST("/logout", s.in.LoginCtrl.DeleteSession)
}
