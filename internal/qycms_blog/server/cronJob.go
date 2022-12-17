package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/qycms_blog/job"
)

func NewQyCronJob(logger log.Logger, cjs *job.CronJobServer) *job.QyCronJob {
	return job.NewQyCronJob(logger, cjs)
}
