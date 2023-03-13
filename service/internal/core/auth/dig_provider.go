package auth

import (
	"go.uber.org/dig"

	"personal-website-golang/service/internal/thirdparty/logger"
)

func newAuth(in digIn) digOut {
	return digOut{
		User: newAuth(in),
	}
}

type digIn struct {
	dig.In

	AppLogger logger.ILogger `name:"appLogger"`
}

type digOut struct {
	dig.Out

	User IUser
}
