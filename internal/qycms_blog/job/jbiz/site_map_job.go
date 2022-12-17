package jbiz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/qycms_blog/service"
)

type SiteMapJobRepo struct {
	s   *service.BlogAdminJobsService
	ctx context.Context
	log *log.Helper
}

func (c *SiteMapJobRepo) Run() {
	log.Error("网站地图更新定时任务开始")
	err := c.s.GeneratorMap(c.ctx)
	if err != nil {
		log.Error("网站地图更新失败: %s", err)
	}
	log.Error("网站地图更新定时任务结束")
}

func NewSiteMapJobRepo(s *service.BlogAdminJobsService, logger log.Logger) *SiteMapJobRepo {
	return &SiteMapJobRepo{
		s:   s,
		ctx: context.Background(),
		log: log.NewHelper(logger),
	}
}
