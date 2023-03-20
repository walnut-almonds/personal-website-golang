package localcache

import (
	"context"
	"time"

	"github.com/patrickmn/go-cache"
	"go.uber.org/dig"

	adminConfig "personal-website-golang/service/internal/config/admin"
	webConfig "personal-website-golang/service/internal/config/web"
	"personal-website-golang/service/internal/thirdparty/logger"
)

func NewAdminDefault(in digInAdmin) ILocalCache {
	in.SysLogger.Info(context.Background(), "new local cache")
	return &localCache{
		c: cache.New(in.AppConf.GetLocalCacheConfig().DefaultExpirationSec*time.Second, 10*time.Second),
	}
}

type digInAdmin struct {
	dig.In

	AppConf   adminConfig.IAppConfig
	SysLogger logger.ILogger `name:"sysLogger"`
}

func NewWebDefault(in digInWeb) ILocalCache {
	in.SysLogger.Info(context.Background(), "new local cache")
	return &localCache{
		c: cache.New(in.AppConf.GetLocalCacheConfig().DefaultExpirationSec*time.Second, 10*time.Second),
	}
}

type digInWeb struct {
	dig.In

	AppConf   webConfig.IAppConfig
	OpsConf   webConfig.IOpsConfig
	SysLogger logger.ILogger `name:"sysLogger"`
}
