package user_admin_auth

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"

	"personal-website-golang/service/internal/constant"
	"personal-website-golang/service/internal/model/bo"
)

type IValidateToken interface {
	Handle(ctx context.Context, token string) (*bo.UserAdminInfo, error)
}

func newValidateToken(in digIn) IValidateToken {
	return &validateToken{
		in: in,
	}
}

type validateToken struct {
	in digIn
}

func (uc *validateToken) Handle(ctx context.Context, token string) (*bo.UserAdminInfo, error) {
	if token == "" {
		return nil, constant.Error_AuthTokenUnauthorized
	}

	keyByToken := constant.RedisKey_UserAdminToken + "_" + token
	idStr, err := uc.in.Redis.Get(ctx, keyByToken).Result()
	if err != nil {
		if err != redis.Nil {
			uc.in.AppLogger.Warn(ctx, fmt.Sprintf("[validateToken] redis get error, key:[%s] err:[%v]", keyByToken, err))
		}
		return nil, constant.Error_AuthTokenUnauthorized
	}

	keyByID := constant.RedisKey_UserAdminToken + "_" + idStr
	if err = uc.in.Redis.Set(ctx, keyByToken, idStr, constant.RedisTTL_UserAdminToken).Err(); err != nil {
		uc.in.AppLogger.Warn(ctx, fmt.Sprintf("[validateToken] redis set error, key:[%s] err:[%v]", keyByToken, err))
		return nil, constant.Error_AuthTokenCreateFailed
	}
	if err = uc.in.Redis.Set(ctx, keyByID, token, constant.RedisTTL_UserAdminToken).Err(); err != nil {
		uc.in.AppLogger.Warn(ctx, fmt.Sprintf("[validateToken] redis set error, key:[%s] err:[%v]", keyByID, err))
		return nil, constant.Error_AuthTokenCreateFailed
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		uc.in.AppLogger.Warn(ctx, fmt.Sprintf("[validateToken] id parse error, key:[%s] err:[%v]", keyByID, err))
		return nil, constant.Error_AuthTokenUnauthorized
	}

	return &bo.UserAdminInfo{
		ID:    id,
		Token: token,
	}, nil
}
