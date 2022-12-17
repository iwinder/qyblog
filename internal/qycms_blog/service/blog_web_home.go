package service

import (
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/web/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"github.com/iwinder/qingyucms/internal/qycms_blog/conf"
)

type BlogWebApiService struct {
	v1.UnimplementedQyWebSiteConfigServer
	conf *conf.Qycms
	site *biz.SiteConfigUsecase
	au   *biz.ArticleUsecase
	mu   *biz.MenusUsecase
	lu   *biz.LinkUsecase
	slu  *biz.ShortLinkUsecase
	cu   *biz.CategoryUsecase
	tu   *biz.TagsUsecase
	ctu  *biz.CommentUsecase
	avu  *biz.ArticleVisitorUsecase
}

func NewBlogWebApiService(
	site *biz.SiteConfigUsecase, au *biz.ArticleUsecase,
	mu *biz.MenusUsecase, lu *biz.LinkUsecase, slu *biz.ShortLinkUsecase,
	cu *biz.CategoryUsecase, tu *biz.TagsUsecase, ctu *biz.CommentUsecase,
	avu *biz.ArticleVisitorUsecase, conf *conf.Qycms,
) *BlogWebApiService {
	return &BlogWebApiService{
		site: site, au: au, mu: mu,
		lu: lu, slu: slu, cu: cu, tu: tu, ctu: ctu,
		avu: avu, conf: conf,
	}
}
