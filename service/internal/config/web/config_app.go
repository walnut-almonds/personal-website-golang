package admin

import (
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func newAppConfig() IAppConfig {
	cfg := &AppConfigSetup{
		v:              viper.New(),
		lastChangeTime: time.Time{},
	}

	cfg.Load()

	return cfg
}

type IAppConfig interface {
	GetServerConfig() ServerConfig
	GetLogConfig() LogConfig
	GetGinConfig() GinConfig
	GetMySQLConfig() MySQLConfig
	GetLocalCacheConfig() LocalCacheConfig
	GetSlackConfig() SlackConfig
	GetMongoConfig() MongoConfig
	GetTelegramConfig() TelegramConfig
}

type AppConfigSetup struct {
	v              *viper.Viper
	lastChangeTime time.Time

	AppConfig AppConfig `mapstructure:"app_config"`
}

type AppConfig struct {
	ServerConfig     ServerConfig     `mapstructure:"server_config"`
	LogConfig        LogConfig        `mapstructure:"log_config"`
	GinConfig        GinConfig        `mapstructure:"gin_config"`
	MySQLConfig      MySQLConfig      `mapstructure:"mysql_config"`
	LocalCacheConfig LocalCacheConfig `mapstructure:"local_cache_config"`
	SlackConfig      SlackConfig      `mapstructure:"slack_config"`
	MongoConfig      MongoConfig      `mapstructure:"mongo_config"`
	TelegramConfig   TelegramConfig   `mapstructure:"telegram_config"`
}

func (c *AppConfigSetup) GetLastChangeTime() time.Time {
	return c.lastChangeTime
}

func (c *AppConfigSetup) Load() {
	c.loadYaml()
}

func (c *AppConfigSetup) loadYaml() {
	path, err := filepath.Abs("conf.d")
	if err != nil {
		panic(err)
	}
	c.v.SetConfigName("config.yaml")
	c.v.SetConfigType("yaml")
	c.v.AddConfigPath(path)
	if err := c.v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := c.v.Unmarshal(c); err != nil {
		panic(err)
	}
	c.v.OnConfigChange(func(in fsnotify.Event) {
		c.v.Unmarshal(c)
		c.lastChangeTime = time.Now()
	})
	c.v.WatchConfig()
}

func (cfg *AppConfigSetup) GetServerConfig() ServerConfig {
	return cfg.AppConfig.ServerConfig
}

func (cfg *AppConfigSetup) GetLogConfig() LogConfig {
	return cfg.AppConfig.LogConfig
}

func (cfg *AppConfigSetup) GetGinConfig() GinConfig {
	return cfg.AppConfig.GinConfig
}

func (cfg *AppConfigSetup) GetMySQLConfig() MySQLConfig {
	return cfg.AppConfig.MySQLConfig
}

func (cfg *AppConfigSetup) GetLocalCacheConfig() LocalCacheConfig {
	return cfg.AppConfig.LocalCacheConfig
}

func (cfg *AppConfigSetup) GetSlackConfig() SlackConfig {
	return cfg.AppConfig.SlackConfig
}

func (cfg *AppConfigSetup) GetMongoConfig() MongoConfig {
	return cfg.AppConfig.MongoConfig
}

func (cfg *AppConfigSetup) GetTelegramConfig() TelegramConfig {
	return cfg.AppConfig.TelegramConfig
}
