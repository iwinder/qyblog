package job

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/qycms_blog/conf"
	"github.com/iwinder/qingyucms/internal/qycms_blog/job/jbiz"
	"github.com/robfig/cron/v3"
)

type PostVCountJob struct {
	cj   *jbiz.PostVCountJobRepo
	cron *cron.Cron
	conf *conf.Qycms
	log  *log.Helper
}

func NewPostVCountJob(cj *jbiz.PostVCountJobRepo, conf *conf.Qycms, logger log.Logger) *PostVCountJob {
	Conrs := NewWithSeconds() // 定时任务
	return &PostVCountJob{
		cj:   cj,
		cron: Conrs,
		conf: conf,
		log:  log.NewHelper(logger),
	}
}

func (c *PostVCountJob) Handle() {
	if len(c.conf.Jobs.PostViewCountJobCron) == 0 {
		c.log.Error(fmt.Errorf("更新文章浏览数据定时任务未开启"))
		return
	}
	id, err := c.cron.AddJob(c.conf.Jobs.PostViewCountJobCron, c.cj)
	if err != nil {
		c.log.Error(fmt.Errorf("更新文章浏览数据定时任务注册失败: %w", err))
		return
	}
	c.log.Error(fmt.Errorf("更新文章浏览数据定时任务注册成功,ID为 : %d", id))
	c.cron.Start()
}

func (c *PostVCountJob) Stop() {
	c.cron.Stop()
}
