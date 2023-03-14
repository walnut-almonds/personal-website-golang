package admin

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func newOpsConfig() IOpsConfig {
	ojb := &OpsConfigSetup{
		v:              viper.New(),
		lastChangeTime: time.Now(),
	}

	ojb.Load()

	return ojb
}

type IOpsConfig interface {
	GetOpsMySQLConfig() MySQLOps
	GetOpsFileServerConfig() FileServerOps
	GetDWHServerConfig() DWHOps
	GetOpsMongoConfig() MongoOps
}

type OpsConfigSetup struct {
	v              *viper.Viper
	lastChangeTime time.Time

	OpsConfig OpsConfig `mapstructure:"ops_config"`
}

type OpsConfig struct {
	MySQLOps      MySQLOps      `mapstructure:"mysql_ops"`
	FileServerOps FileServerOps `mapstructure:"file_server_ops"`
	DWHOps        DWHOps        `mapstructure:"dwh_ops"`
	MongoOps      MongoOps      `mapstructure:"mongo_ops"`
}

func (c *OpsConfigSetup) GetLastChangeTime() time.Time {
	return c.lastChangeTime
}

func (c *OpsConfigSetup) Load() {
	c.loadYaml()
}

func (c *OpsConfigSetup) loadYaml() {
	path, err := filepath.Abs("conf.d")
	if err != nil {
		panic(err)
	}
	c.v.SetConfigName("ops.yaml")
	c.v.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	fmt.Println(path)
	c.v.AddConfigPath(path)
	if err := c.v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := c.v.Unmarshal(c); err != nil {
		panic(err)
	}
	c.v.OnConfigChange(func(in fsnotify.Event) {
		if err := c.v.Unmarshal(c); err != nil {
			panic(err)
		}
		c.lastChangeTime = time.Now()
	})
	c.v.WatchConfig()
}

func (cfg *OpsConfigSetup) GetOpsMySQLConfig() MySQLOps {
	return cfg.OpsConfig.MySQLOps
}

func (cfg *OpsConfigSetup) GetOpsFileServerConfig() FileServerOps {
	return cfg.OpsConfig.FileServerOps
}

func (cfg *OpsConfigSetup) GetDWHServerConfig() DWHOps {
	return cfg.OpsConfig.DWHOps
}

func (cfg *OpsConfigSetup) GetOpsMongoConfig() MongoOps {
	return cfg.OpsConfig.MongoOps
}
