package rbac

import (
	"go.uber.org/dig"
	"personal-website-golang/service/internal/thirdparty/logger"
)

func NewAdminAuth(in digInAdmin) IAuthClient {
	return newAuth(in.SysLogger, )
}

type digInAdmin struct {
	dig.In

	SysLogger logger.ILogger `name:"sysLogger"`
}

func NewWebAuth(in digInWeb) IAuthClient {

}

type digInWeb struct {
	dig.In

	SysLogger logger.ILogger `name:"sysLogger"`
}
