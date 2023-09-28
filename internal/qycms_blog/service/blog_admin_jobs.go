package service

import (
	"context"
	"github.com/iwinder/qyblog/internal/qycms_blog/biz"
	"github.com/iwinder/qyblog/internal/qycms_blog/conf"
)

type BlogAdminJobsService struct {
	conf *conf.Qycms
	ctu  *biz.CommentUsecase
	au   *biz.ArticleUsecase
	site *biz.SiteMapUsecase
}

func NewBlogAdminJobsService(conf *conf.Qycms, ctu *biz.CommentUsecase, au *biz.ArticleUsecase, site *biz.SiteMapUsecase) *BlogAdminJobsService {
	return &BlogAdminJobsService{conf: conf, ctu: ctu, au: au, site: site}
}

func (s *BlogAdminJobsService) UpdateContentCountAndObjIdsJob(ctx context.Context) error {
	return s.ctu.UpdateContentCountAndObjIds(ctx)
}
func (s *BlogAdminJobsService) GeneratorMap(ctx context.Context) error {
	return s.site.GeneratorMap(ctx, s.conf.SiteMapPath)
}

func (s *BlogAdminJobsService) UpdateAllPostsCount(ctx context.Context) {
	s.au.UpdateAllPostsCount(ctx)
}

func (s *BlogAdminJobsService) EmailToNotSend(ctx context.Context) {
	s.ctu.EmailToNotSend(ctx, s.conf)
}
