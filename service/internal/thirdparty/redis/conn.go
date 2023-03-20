package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"personal-website-golang/service/internal/thirdparty/logger"
)

func Init(sysLogger logger.ILogger, address, password string, database int) IRedisClient {
	client := &redisClient{
		client: redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password,
			DB:       database,
		}),
	}

	sysLogger.Info(context.Background(), fmt.Sprintf("redis [%d] init success", database))

	return client
}
