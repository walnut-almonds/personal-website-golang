package crontab

import (
	"github.com/robfig/cron/v3"
)

type ICronJob interface {
	Start()
	AddFunc(spec string, cmd func()) (cron.EntryID, error)
}

type cronJob struct {
	cron *cron.Cron
}

func (job *cronJob) Start() {
	job.cron.Start()
}

func (job *cronJob) AddFunc(spec string, cmd func()) (cron.EntryID, error) {
	return job.cron.AddFunc(spec, cmd)
}
