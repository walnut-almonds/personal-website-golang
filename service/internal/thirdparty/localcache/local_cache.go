package localcache

import (
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
	"go.uber.org/dig"

	config "personal-website-golang/service/internal/config/admin"
)

var local *localCache

type ILocalCache interface {
	Save(key string, value interface{})
	Get(key string) (interface{}, bool)
	Delete(key string)
	Increment(key string, n int)
	DeleteSet(key string)
}

type digIn struct {
	dig.In

	AppConf config.IAppConfig
}

func NewDefault(in digIn) ILocalCache {
	return &localCache{
		c: cache.New(in.AppConf.GetLocalCacheConfig().DefaultExpirationSec*time.Second, 10*time.Second),
	}
}

func New(defaultExpirationSec time.Duration) ILocalCache {
	return &localCache{
		c: cache.New(defaultExpirationSec*time.Second, 10*time.Second),
	}
}

type localCache struct {
	c *cache.Cache
}

func (lc *localCache) Save(key string, value interface{}) {
	lc.c.SetDefault(key, value)
}

func (lc *localCache) Get(key string) (interface{}, bool) {
	value, existed := lc.c.Get(key)
	return value, existed
}

func (lc *localCache) Delete(key string) {
	lc.c.Delete(key)
}

func (lc *localCache) DeleteSet(key string) {
	cacheSet := lc.c.Items()

	for k := range cacheSet {
		if strings.Contains(k, key) {
			lc.Delete(k)
		}
	}
}

func (lc *localCache) Increment(key string, n int) {
	lc.c.IncrementInt(key, n)
}
