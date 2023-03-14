package admin

import "time"

type ServerConfig struct {
	TimeZone int `mapstructure:"time_zone"`
}

type LogConfig struct {
	Name  string `mapstructure:"name"`
	Env   string `mapstructure:"env"`
	Level string `mapstructure:"level"`
}

type GinConfig struct {
	AdminAddress string `mapstructure:"admin_address"`
	DebugMode    bool   `mapstructure:"debug_mode"`
}

type MySQLConfig struct {
	LogMode        bool `mapstructure:"log_mode"`
	MaxIdle        int  `mapstructure:"max_idle"`
	MaxOpen        int  `mapstructure:"max_open"`
	ConnMaxLifeMin int  `mapstructure:"conn_max_life_min"`
}

type LocalCacheConfig struct {
	DefaultExpirationSec time.Duration `mapstructure:"default_expiration_sec"`
}

type SlackConfig struct {
	Token     string `mapstructure:"token"`
	ChannelID string `mapstructure:"channel_id"`
	BotName   string `mapstructure:"bot_name"`
}

type MongoConfig struct {
	ReplicaName     string        `mapstructure:"replica_name"`
	ReadPreference  string        `mapstructure:"read_preference"`
	MaxPoolSize     int           `mapstructure:"max_pool_size"`
	MinPoolSize     int           `mapstructure:"min_pool_size"`
	MaxConnIdleTime time.Duration `mapstructure:"max_conn_idle_time"`
	MaxStaleness    time.Duration `mapstructure:"max_staleness"`
	SlowLogEnable   bool          `mapstructure:"slow_log_enable"`
	SlowLogJudgment time.Duration `mapstructure:"slow_log_judgment"`
}

type TelegramConfig struct {
	Token     string `mapstructure:"token"`
	AM6ChatID int64  `mapstructure:"am6_chat_id"`
	HK6ChatID int64  `mapstructure:"hk6_chat_id"`
}
