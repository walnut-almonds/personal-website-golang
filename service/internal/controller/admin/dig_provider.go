package admin

import (
	"personal-website-golang/service/internal/core/user_admin_auth"
	"personal-website-golang/service/internal/thirdparty/rbac"
	"sync"

	"go.uber.org/dig"
)

var (
	once sync.Once
	self *packet
)

func NewRestCtl(in digIn) digOut {

}

type packet struct {
	in digIn

	digOut
}

type digIn struct {
	dig.In

	AuthUseCase authUseCaseIn

	RBACAuthClient rbac.IAuthClient
}

type digOut struct {
	dig.Out

	AuthCtrl IAuthCtrl

	AuthMiddleware IAuthMiddleware
}

type authUseCaseIn struct {
	dig.In

	InsertSession user_admin_auth.IInsertSession
	DeleteSession user_admin_auth.IDeleteSession

	ValidateToken user_admin_auth.IValidateToken
}
