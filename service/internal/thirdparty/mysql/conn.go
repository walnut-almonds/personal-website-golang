package mysql

import (
	"context"
	"fmt"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
	self *DBClient
)

type IMySQLClient interface {
	Session() *gorm.DB
}

type DBClient struct {
	client *gorm.DB
}

func initWithConfig(in digIn) IMySQLClient {
	connect := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		in.OpsConf.GetOpsMySQLConfig().Username,
		in.OpsConf.GetOpsMySQLConfig().Password,
		in.OpsConf.GetOpsMySQLConfig().Address,
		in.OpsConf.GetOpsMySQLConfig().Database,
	)

	var err error
	db, err = gorm.Open(mysql.Open(connect))
	if err != nil {
		panic(fmt.Sprintf("conn: %s err: %v", connect, err))
	}

	if in.AppConf.GetMySQLConfig().LogMode {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	in.SysLogger.Info(context.Background(), fmt.Sprintf("mysql [%s] connect success", in.OpsConf.GetOpsMySQLConfig().Database))
	// SetMaxIdleConn sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(in.AppConf.GetMySQLConfig().MaxIdle)
	// SetMaxOpenConn sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(in.AppConf.GetMySQLConfig().MaxOpen)
	// SetConnMaxLifetime sets the maximum amount of timeUtil a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(in.AppConf.GetMySQLConfig().ConnMaxLifeMin) * time.Minute)

	return &DBClient{db}
}

// Session creates an original gorm.DB session.
func (*DBClient) Session() *gorm.DB {
	return db.Session(&gorm.Session{})
}
