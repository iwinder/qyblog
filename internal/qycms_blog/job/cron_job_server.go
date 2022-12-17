package job

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/qycms_blog/conf"
	"github.com/iwinder/qingyucms/internal/qycms_blog/job/jbiz"
	"github.com/robfig/cron/v3"
)

type CronJobServer struct {
	cron *cron.Cron
	conf *conf.Qycms
	log  *log.Helper
	cj   *jbiz.CountCommentJobRepo
	pcj  *jbiz.PostVCountJobRepo
	smj  *jbiz.SiteMapJobRepo
	emj  *jbiz.EmailSendJobRepo
}

func NewCronJobServer(conf *conf.Qycms, logger log.Logger,
	cj *jbiz.CountCommentJobRepo,
	pcj *jbiz.PostVCountJobRepo,
	smj *jbiz.SiteMapJobRepo,
	emj *jbiz.EmailSendJobRepo,
) *CronJobServer {
	Conrs := NewWithSeconds() // 定时任务
	return &CronJobServer{
		cron: Conrs,
		conf: conf,
		log:  log.NewHelper(logger),
		cj:   cj,
		pcj:  pcj,
		smj:  smj,
		emj:  emj,
	}
}

func (c *CronJobServer) Handle() {
	c.registerCommentJob()
	c.registerPostVCountJob()
	c.registerSiteMapJob()
	c.registerEmailEndJob()
	c.cron.Start()
}

func (c *CronJobServer) Stop() {
	c.cron.Stop()
}

func (c *CronJobServer) registerCommentJob() {
	if len(c.conf.Jobs.PostViewCountJobCron) == 0 {
		c.log.Error(fmt.Errorf("更新评论数据定时任务未开启"))
		return
	}
	id, err := c.cron.AddJob(c.conf.Jobs.CommentJobCron, c.cj)
	if err != nil {
		c.log.Error(fmt.Errorf("更新评论数据定时任务注册失败: %w", err))
		return
	}
	c.log.Error(fmt.Errorf("更新评论数据定时任务注册成功,ID为 : %d", id))
}

func (c *CronJobServer) registerPostVCountJob() {
	if len(c.conf.Jobs.PostViewCountJobCron) == 0 {
		c.log.Error(fmt.Errorf("更新文章浏览数据定时任务未开启"))
		return
	}
	id, err := c.cron.AddJob(c.conf.Jobs.PostViewCountJobCron, c.pcj)
	if err != nil {
		c.log.Error(fmt.Errorf("更新文章浏览数据定时任务注册失败: %w", err))
		return
	}
	c.log.Error(fmt.Errorf("更新文章浏览数据定时任务注册成功,ID为 : %d", id))
}

func (c *CronJobServer) registerSiteMapJob() {
	if len(c.conf.Jobs.SiteMapJobCron) == 0 {
		c.log.Error(fmt.Errorf("更新网站地图数据定时任务未开启"))
		return
	}
	id, err := c.cron.AddJob(c.conf.Jobs.SiteMapJobCron, c.smj)
	if err != nil {
		c.log.Error(fmt.Errorf("更新网站地图数据定时任务注册失败: %w", err))
		return
	}
	c.log.Error(fmt.Errorf("更新网站地图数据定时任务注册成功,ID为 : %d", id))
}

func (c *CronJobServer) registerEmailEndJob() {
	if len(c.conf.Jobs.EmailSendJobCron) == 0 {
		c.log.Error(fmt.Errorf("发送邮件回复提醒定时任务未开启"))
		return
	}
	id, err := c.cron.AddJob(c.conf.Jobs.EmailSendJobCron, c.emj)
	if err != nil {
		c.log.Error(fmt.Errorf("发送邮件回复提醒定时任务注册失败: %w", err))
		return
	}
	c.log.Error(fmt.Errorf("发送邮件回复提醒定时任务注册成功,ID为 : %d", id))
}
