package ftp

import (
	"context"

	"go.uber.org/dig"

	"personal-website-golang/service/internal/thirdparty/logger"
)

type digIn struct {
	dig.In

	SysLogger logger.ILogger `name:"sysLogger"`
}

func NewFtpCli(in digIn) IFtp {
	in.SysLogger.Info(context.Background(), "new ftp cli")
	return newFtpCli()
}
