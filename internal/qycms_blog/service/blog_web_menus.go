package service

import (
	"context"
	v1 "github.com/iwinder/qyblog/api/qycms_bff/web/v1"
	"github.com/iwinder/qyblog/internal/qycms_blog/biz"
)

func (b *BlogWebApiService) GetQyWebFooterMenus(ctx context.Context, in *v1.GetQyWebMenusRequest) (*v1.GetQyWebMenusReply, error) {
	opts := biz.MenusDOListOption{}
	opts.TargetId = 2
	opts.PageFlag = false
	opts.ParentId = 0
	objList, err := b.mu.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}

	objs := make([]*v1.MenusWebInfoResponse, 0, len(objList.Items))
	for _, item := range objList.Items {
		titem := bizToMenusWebResponse(item)
		objs = append(objs, titem)
	}
	return &v1.GetQyWebMenusReply{Items: objs}, nil
}

func (b *BlogWebApiService) GetQyWebHeaderMenus(ctx context.Context, request *v1.GetQyWebMenusRequest) (*v1.GetQyWebMenusReply, error) {
	opts := biz.MenusDOListOption{}
	opts.TargetId = 1
	opts.PageFlag = false
	opts.ParentId = 0
	objList, err := b.mu.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}

	objs := make([]*v1.MenusWebInfoResponse, 0, len(objList.Items))
	for _, item := range objList.Items {
		titem := bizToMenusWebResponse(item)
		objs = append(objs, titem)
	}
	return &v1.GetQyWebMenusReply{Items: objs}, nil
}

func bizToMenusWebResponse(obj *biz.MenusDO) *v1.MenusWebInfoResponse {
	objInfoRsp := &v1.MenusWebInfoResponse{
		Id:      obj.ID,
		Name:    obj.Name,
		Url:     obj.Url,
		Blanked: obj.Blanked,
	}
	if obj.Children != nil && len(obj.Children) > 0 {
		cobjList := make([]*v1.MenusWebInfoResponse, 0, len(obj.Children))
		for _, cobj := range obj.Children {
			citme := bizToMenusWebResponse(cobj)
			cobjList = append(cobjList, citme)
		}
		objInfoRsp.Children = cobjList
	}
	return objInfoRsp
}
