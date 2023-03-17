package mysql

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IMySQLClient interface {
	Session() *gorm.DB
}

type DBClient struct {
	client *gorm.DB
}

type Config struct {
	Username string
	Password string
	Address  string
	Database string

	LogMode        bool
	MaxIdle        int
	MaxOpen        int
	ConnMaxLifeMin int
}

func initWithConfig(config Config) IMySQLClient {
	connect := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		config.Username,
		config.Password,
		config.Address,
		config.Database,
	)

	db, err := gorm.Open(mysql.Open(connect))
	if err != nil {
		panic(fmt.Sprintf("conn: %s err: %v", connect, err))
	}

	if config.LogMode {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	in.SysLogger.Info(context.Background(), fmt.Sprintf("mysql [%s] connect success", config.Database))

	sqlDB.SetMaxIdleConns(config.MaxIdle)
	sqlDB.SetMaxOpenConns(config.MaxOpen)
	sqlDB.SetConnMaxLifetime(time.Duration(config.ConnMaxLifeMin) * time.Minute)

	return &DBClient{db}
}

func (c *DBClient) Session() *gorm.DB {
	return c.client.Session(&gorm.Session{})
}
