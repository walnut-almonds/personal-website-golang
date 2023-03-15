package admin

import "github.com/gin-gonic/gin"

type IAuthMiddleware interface {
	MiddlewareAuth(ctx *gin.Context)
}

func newAuthMiddleware(in digIn) IAuthMiddleware {
	return &authMiddlewareCtrl{
		in: in,
	}
}

type authMiddlewareCtrl struct {
	in digIn
}

func (mid *authMiddlewareCtrl) MiddlewareAuth(ctx *gin.Context) {

}
