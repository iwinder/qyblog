package service

import (
	w1 "github.com/iwinder/qingyucms/api/qycms_bff/web/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
)

type BlogWebApiService struct {
	site *biz.SiteConfigUsecase
}

func (b BlogWebApiService) ListQyBaseSiteConfig(ctx context.Context, request *w1.ListQyWebSiteConfigRequest) (*w1.ListQyWebSiteConfigReply, error) {
	//TODO implement me
	panic("implement me")
}

func NewBlogWebApiService(
	site *biz.SiteConfigUsecase) *BlogWebApiService {
	return &BlogWebApiService{
		site: site,
	}
}
