package service

import (
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/web/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
)

type BlogWebApiService struct {
	v1.UnimplementedQyWebSiteConfigServer
	site *biz.SiteConfigUsecase
}

func NewBlogWebApiService(
	site *biz.SiteConfigUsecase) *BlogWebApiService {
	return &BlogWebApiService{
		site: site,
	}
}
