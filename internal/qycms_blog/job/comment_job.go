package job

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/qycms_blog/conf"
	"github.com/iwinder/qingyucms/internal/qycms_blog/job/jbiz"
	"github.com/robfig/cron/v3"
)

// CountCommentJob 评论数据更新定时 每天1点
type CountCommentJob struct {
	cj   *jbiz.CountCommentJobRepo
	cron *cron.Cron
	conf *conf.Qycms
	log  *log.Helper
}

func NewCommentJob(cj *jbiz.CountCommentJobRepo, conf *conf.Qycms, logger log.Logger) *CountCommentJob {
	Conrs := NewWithSeconds() // 定时任务
	return &CountCommentJob{
		cj:   cj,
		cron: Conrs,
		conf: conf,
		log:  log.NewHelper(logger),
	}
}

func (c *CountCommentJob) Handle() {
	id, err := c.cron.AddJob(c.conf.Jobs.CommentJobCron, c.cj)
	if err != nil {
		c.log.Error(fmt.Errorf("更新评论数据定时任务注册失败: %w", err))
		return
	}
	c.log.Error(fmt.Errorf("更新评论数据定时任务注册成功,ID为 : %d", id))
	c.cron.Start()
}

func (c *CountCommentJob) Stop() {
	c.cron.Stop()
}
