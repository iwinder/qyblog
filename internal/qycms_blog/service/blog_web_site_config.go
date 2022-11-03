package service

import (
	"context"
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/web/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
)

func (b *BlogWebApiService) ListQyBaseSiteConfig(ctx context.Context, request *v1.ListQyWebSiteConfigRequest) (*v1.ListQyWebSiteConfigReply, error) {
	opts := biz.SiteConfigDOListOption{
		Types: "1,2",
	}
	obj, err := b.site.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	objs := make([]*v1.SiteConfigResponse, 0, len(obj))
	for _, item := range obj {
		titem := bizToSiteConfigWebResponse(item)
		objs = append(objs, &titem)
	}

	return &v1.ListQyWebSiteConfigReply{Items: objs}, nil
}

func bizToSiteConfigWebResponse(obj *biz.SiteConfigDO) v1.SiteConfigResponse {
	objInfoRsp := v1.SiteConfigResponse{
		Id:          obj.ID,
		ConfigKey:   obj.ConfigKey,
		ConfigValue: obj.ConfigValue,
		ConfigName:  obj.ConfigName,
		ConfigTip:   obj.ConfigTip,
		Ftype:       int32(obj.Ftype),
	}
	return objInfoRsp
}
