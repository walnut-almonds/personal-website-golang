package user_admin_auth

import (
	"github.com/go-redis/redis/v8"
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

	Redis *redis.Client
}

type digOut struct {
	dig.Out

	User IUser
}
