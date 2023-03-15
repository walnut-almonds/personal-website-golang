package admin

import (
	"personal-website-golang/service/internal/core/auth"
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
}

type digOut struct {
	dig.Out

	AuthCtrl IAuthCtrl

	AuthMiddleware IAuthMiddleware
}

type authUseCaseIn struct {
	dig.In

	InsertSession auth.IInsertSession
	DeleteSession auth.IDeleteSession

	ValidateToken auth.IValidateToken
}
