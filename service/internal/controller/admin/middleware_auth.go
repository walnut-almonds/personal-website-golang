package admin

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
	"net/http"
	utilCtx "personal-website-golang/service/internal/util/context"
)

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
	token := ctx.GetHeader("Authorization")

	adminInfo, err := mid.in.AuthUseCase.ValidateToken.Handle(ctx, token)
	if err != nil {
		ctx.JSON(http.StatusForbidden, err)
		ctx.Abort()
		return
	}

	utilCtx.SetAuthInfo(ctx, adminInfo)

	ctx.Next()
}

func (mid *authMiddlewareCtrl) MiddlewareRBAC(ctx *gin.Context) {
	authInfo, exist := utilCtx.GetAuthInfo(ctx)
	if !exist {
		ctx.JSON(http.StatusForbidden, xerrors.Errorf("auth info not exist"))
		ctx.Abort()
		return
	}

	approve, err := mid.in.RBACAuthClient.GetAdminAuthInstance().Enforce(authInfo.ID, ctx.Request.URL.Path, ctx.Request.Method)
	if err != nil {
		ctx.JSON(http.StatusForbidden, err)
		ctx.Abort()
		return
	}
	if approve {
		ctx.Next()
	} else {
		ctx.JSON(http.StatusForbidden, xerrors.Errorf("forbidden"))
		ctx.Abort()
		return
	}
}
