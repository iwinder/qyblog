package jbiz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qyblog/internal/qycms_blog/service"
)

type EmailSendJobRepo struct {
	s   *service.BlogAdminJobsService
	ctx context.Context
	log *log.Helper
}

func (c *EmailSendJobRepo) Run() {
	log.Error("邮件推送回复消息定时任务开始")
	c.s.EmailToNotSend(c.ctx)

	log.Error("邮件推送回复消息定时任务结束")
}

func NewEmailSendJobRepo(s *service.BlogAdminJobsService, logger log.Logger) *EmailSendJobRepo {
	return &EmailSendJobRepo{
		s:   s,
		ctx: context.Background(),
		log: log.NewHelper(logger),
	}
}
