package job

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/qycms_blog/conf"
	"github.com/iwinder/qingyucms/internal/qycms_blog/job/jbiz"
	"github.com/robfig/cron/v3"
)

type SiteMapJob struct {
	cj   *jbiz.SiteMapJobRepo
	cron *cron.Cron
	conf *conf.Qycms
	log  *log.Helper
}

func NewSiteMapJob(cj *jbiz.SiteMapJobRepo, conf *conf.Qycms, logger log.Logger) *SiteMapJob {
	Conrs := NewWithSeconds() // 定时任务
	return &SiteMapJob{
		cj:   cj,
		cron: Conrs,
		conf: conf,
		log:  log.NewHelper(logger),
	}
}

func (c *SiteMapJob) Handle() {
	if len(c.conf.Jobs.SiteMapJobCron) == 0 {
		c.log.Error(fmt.Errorf("更新评论数据定时任务未开启"))
		return
	}
	id, err := c.cron.AddJob(c.conf.Jobs.SiteMapJobCron, c.cj)
	if err != nil {
		c.log.Error(fmt.Errorf("更新评论数据定时任务注册失败: %w", err))
		return
	}
	c.log.Error(fmt.Errorf("更新评论数据定时任务注册成功,ID为 : %d", id))
	c.cron.Start()
}

func (c *SiteMapJob) Stop() {
	c.cron.Stop()
}
