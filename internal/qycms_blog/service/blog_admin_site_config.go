package service

import (
	"context"
	v1 "github.com/iwinder/qyblog/api/qycms_bff/admin/v1"
	"github.com/iwinder/qyblog/internal/qycms_blog/biz"
)

func (s *BlogAdminUserService) CreateQyAdminSiteConfig(ctx context.Context, in *v1.CreateQyAdminSiteConfigRequest) (*v1.CreateQyAdminSiteConfigReply, error) {
	objDO := &biz.SiteConfigDO{
		ConfigKey:   in.ConfigKey,
		ConfigValue: in.ConfigValue,
		ConfigName:  in.ConfigName,
		ConfigTip:   in.ConfigTip,
		Ftype:       int(in.Ftype),
	}
	obj, err := s.site.Save(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.CreateQyAdminSiteConfigReply{Id: obj.ID}, nil
}

func (s *BlogAdminUserService) UpdateInBatchesQyAdminSiteConfig(ctx context.Context, in *v1.UpdateBatchesQyAdminSiteConfigRequest) (*v1.UpdateBatchesQyAdminSiteConfigReply, error) {

	objDOs := make([]*biz.SiteConfigDO, 0, len(in.Paramms))
	for _, obj := range in.Paramms {
		if obj.Id > 0 {
			objDO := &biz.SiteConfigDO{
				ConfigKey:   obj.ConfigKey,
				ConfigValue: obj.ConfigValue,
				ConfigName:  obj.ConfigName,
				ConfigTip:   obj.ConfigTip,
				Ftype:       int(obj.Ftype),
			}
			objDO.ID = obj.Id
			objDOs = append(objDOs, objDO)
		}

	}

	err := s.site.UpdateInBatches(ctx, objDOs)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateBatchesQyAdminSiteConfigReply{}, nil
}

func (s *BlogAdminUserService) ListQyAdminSiteConfig(ctx context.Context, in *v1.ListQyAdminSiteConfigRequest) (*v1.ListQyAdminSiteConfigReply, error) {
	opts := biz.SiteConfigDOListOption{
		Types: in.Types,
	}
	obj, err := s.site.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	objs := make([]*v1.SiteConfigResponse, 0, len(obj))
	for _, item := range obj {
		titem := bizToSiteConfigResponse(item)
		objs = append(objs, &titem)
	}

	return &v1.ListQyAdminSiteConfigReply{Items: objs}, nil
}

func bizToSiteConfigResponse(obj *biz.SiteConfigDO) v1.SiteConfigResponse {
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
