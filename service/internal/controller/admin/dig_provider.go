package admin

import (
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

	AuthHttpMiddleware IAuthHttpMiddleware
}

type authUseCaseIn struct {
	dig.In

	AddSession    auth.IAddSession
	DeleteSession auth.IDeleteSession

	ValidateToken auth.IValidateToken
}
