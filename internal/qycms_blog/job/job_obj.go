package job

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/robfig/cron/v3"
)

var ProviderSet = wire.NewSet(NewCronJobServer)

type QyJob interface {
	Handle()
	Stop()
}

type QyCronJob struct {
	cjs *CronJobServer
	log *log.Helper
}

func NewQyCronJob(logger log.Logger,
	cjs *CronJobServer) *QyCronJob {
	return &QyCronJob{log: log.NewHelper(logger), cjs: cjs}
}

func (s *QyCronJob) Start(ctx context.Context) error {
	log.Info("定时任务启动>>>>>>>>>>>>>")
	s.cjs.Handle()
	return nil
}
func (s *QyCronJob) Stop(ctx context.Context) error {
	s.cjs.Stop()
	log.Info("定时任务结束>>>>>>>>>>>>>")
	return nil
}

func NewWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}
