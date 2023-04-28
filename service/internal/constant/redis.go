package constant

import "time"

const (
	RedisKey_MemberToken    = "memberToken"
	RedisKey_UserAdminToken = "userAdminToken"
)

var (
	RedisTTL_MemberToken    = 24 * time.Hour
	RedisTTL_UserAdminToken = 4 * time.Hour
)
