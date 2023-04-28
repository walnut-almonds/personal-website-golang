package context

import (
	"context"

	"github.com/gin-gonic/gin"

	"personal-website-golang/service/internal/model/bo"
)

type keyName = string

const (
	keyName_AuthInfo keyName = "keyName_AuthInfo"
)

func SetAuthInfo(ctx *gin.Context, authInfo *bo.UserAdminInfo) {
	ctx.Set(keyName_AuthInfo, authInfo)
}

func GetAuthInfo(ctx context.Context) (*bo.UserAdminInfo, bool) {
	val := ctx.Value(keyName_AuthInfo)

	v, ok := val.(*bo.UserAdminInfo)
	if !ok {
		return nil, false
	}

	return v, true
}
