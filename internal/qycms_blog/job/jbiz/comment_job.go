package jbiz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/qycms_blog/service"
)

type CountCommentJobRepo struct {
	s   *service.BlogAdminJobsService
	ctx context.Context
	log *log.Helper
}

func (c *CountCommentJobRepo) Run() {
	log.Error("评论数据更新定时任务开始")
	err := c.s.UpdateContentCountAndObjIdsJob(c.ctx)
	if err != nil {
		log.Error("评论数据更新失败: %s", err)
	}
	log.Error("评论数据更新定时任务结束")
}

func NewCountCommentJobRepo(s *service.BlogAdminJobsService, logger log.Logger) *CountCommentJobRepo {
	return &CountCommentJobRepo{
		s:   s,
		ctx: context.Background(),
		log: log.NewHelper(logger),
	}
}
