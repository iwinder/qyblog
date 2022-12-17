package jbiz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/qycms_blog/service"
)

type PostVCountJobRepo struct {
	s   *service.BlogAdminJobsService
	ctx context.Context
	log *log.Helper
}

func (c *PostVCountJobRepo) Run() {
	log.Error("文章数据更新定时任务开始")
	c.s.UpdateAllPostsCount(c.ctx)

	log.Error("文章数据更新定时任务结束")
}

func NewPostVCountJobRepo(s *service.BlogAdminJobsService, logger log.Logger) *PostVCountJobRepo {
	return &PostVCountJobRepo{
		s:   s,
		ctx: context.Background(),
		log: log.NewHelper(logger),
	}
}
