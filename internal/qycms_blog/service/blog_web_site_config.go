package service

import (
	"context"
	v1 "github.com/iwinder/qyblog/api/qycms_bff/web/v1"
	"github.com/iwinder/qyblog/internal/qycms_blog/biz"
)

func (b *BlogWebApiService) ListQyBaseSiteConfig(ctx context.Context, request *v1.ListQyWebSiteConfigRequest) (*v1.ListQyWebSiteConfigReply, error) {

	opts := biz.SiteConfigDOListOption{
		Types: "1,2",
	}
	if len(request.Ftypes) > 0 {
		opts.Types = request.Ftypes
	}
	obj, err := b.site.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	objs := make([]*v1.SiteConfigResponse, 0, len(obj))
	for _, item := range obj {
		titem := bizToSiteConfigWebResponse(item)
		objs = append(objs, titem)
	}

	return &v1.ListQyWebSiteConfigReply{Items: objs}, nil
}

func (b *BlogWebApiService) ListQyOtherSiteConfig(ctx context.Context, request *v1.ListQyWebSiteConfigRequest) (*v1.ListQyWebSiteConfigReply, error) {
	opts := biz.SiteConfigDOListOption{
		Types: "3,4",
	}
	if len(request.Ftypes) > 0 {
		opts.Types = request.Ftypes
	}
	obj, err := b.site.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	objs := make([]*v1.SiteConfigResponse, 0, len(obj))
	for _, item := range obj {
		titem := bizToSiteConfigWebResponse(item)
		objs = append(objs, titem)
	}

	return &v1.ListQyWebSiteConfigReply{Items: objs}, nil
}

func bizToSiteConfigWebResponse(obj *biz.SiteConfigDO) *v1.SiteConfigResponse {
	objInfoRsp := &v1.SiteConfigResponse{
		ConfigKey:   obj.ConfigKey,
		ConfigValue: obj.ConfigValue,
	}
	return objInfoRsp
}
