package crontab

import (
	"log"

	"github.com/robfig/cron/v3"
	"go.uber.org/dig"
)

type ICronJob interface {
	Start()
	AddFunc(spec string, cmd func()) (cron.EntryID, error)
}

type cronJob struct {
	cron *cron.Cron
}

type digIn struct {
	dig.In
}

func NewCronJob(in digIn) ICronJob {
	c := cron.New()
	log.Print("new crontab job")
	return &cronJob{cron: c}
}

func (job *cronJob) Start() {
	job.cron.Start()
}

func (job *cronJob) AddFunc(spec string, cmd func()) (cron.EntryID, error) {
	return job.cron.AddFunc(spec, cmd)
}
