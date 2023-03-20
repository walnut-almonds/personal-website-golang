package crontab

import (
	"context"

	"github.com/robfig/cron/v3"
	"go.uber.org/dig"

	"personal-website-golang/service/internal/thirdparty/logger"
)

type digIn struct {
	dig.In

	SysLogger logger.ILogger `name:"sysLogger"`
}

func NewCronJob(in digIn) ICronJob {
	c := cron.New()
	in.SysLogger.Info(context.Background(), "new crontab job")
	return &cronJob{cron: c}
}
