package service

import (
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/web/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
)

type BlogWebApiService struct {
	v1.UnimplementedQyWebSiteConfigServer
	site *biz.SiteConfigUsecase
	au   *biz.ArticleUsecase
	mu   *biz.MenusUsecase
	lu   *biz.LinkUsecase
	slu  *biz.ShortLinkUsecase
	cu   *biz.CategoryUsecase
	tu   *biz.TagsUsecase
	ctu  *biz.CommentUsecase
}

func NewBlogWebApiService(
	site *biz.SiteConfigUsecase, au *biz.ArticleUsecase,
	mu *biz.MenusUsecase, lu *biz.LinkUsecase, slu *biz.ShortLinkUsecase,
	cu *biz.CategoryUsecase, tu *biz.TagsUsecase, ctu *biz.CommentUsecase,
) *BlogWebApiService {
	return &BlogWebApiService{
		site: site, au: au, mu: mu,
		lu: lu, slu: slu, cu: cu, tu: tu, ctu: ctu,
	}
}
