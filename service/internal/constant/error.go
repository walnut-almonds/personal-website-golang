package constant

import "github.com/pkg/errors"

var (
	Error_AuthTokenCreateFailed = errors.Errorf("token 建立失敗")
	Error_AuthTokenUnauthorized = errors.Errorf("請重新登入")
)
