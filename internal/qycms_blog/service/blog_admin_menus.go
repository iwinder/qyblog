package service

import (
	"context"
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/admin/v1"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
)

// CreateQyAdminMenus 创建
func (s *BlogAdminUserService) CreateQyAdminMenus(ctx context.Context, in *v1.CreateQyAdminMenusRequest) (*v1.CreateQyAdminMenusReply, error) {
	objDO := &biz.MenusDO{
		Name:     in.Name,
		Url:      in.Url,
		Blanked:  in.Blanked,
		ParentId: in.ParentId,
		TargetId: in.TargetId,
	}
	obj, err := s.mu.Create(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.CreateQyAdminMenusReply{Id: obj.ID}, nil
}

// UpdateQyAdminMenus 更新
func (s *BlogAdminUserService) UpdateQyAdminMenus(ctx context.Context, in *v1.UpdateQyAdminMenusRequest) (*v1.UpdateQyAdminMenusReply, error) {
	objDO := &biz.MenusDO{
		ObjectMeta: metaV1.ObjectMeta{ID: in.Id},
		Name:       in.Name,
		Url:        in.Url,
		Blanked:    in.Blanked,
		ParentId:   in.ParentId,
		TargetId:   in.TargetId,
	}
	obj, err := s.mu.Update(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateQyAdminMenusReply{Id: obj.ID}, nil
}

// DeleteQyAdminMenus 删除
func (s *BlogAdminUserService) DeleteQyAdminMenus(ctx context.Context, in *v1.DeleteQyAdminMenusRequest) (*v1.DeleteQyAdminMenusReply, error) {
	err := s.mu.DeleteList(ctx, in.Ids, in.TargetId)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteQyAdminMenusReply{}, nil
}

// ListQyAdminMenus 列表
func (s *BlogAdminUserService) ListQyAdminMenus(ctx context.Context, in *v1.ListQyAdminMenusRequest) (*v1.ListQyAdminMenusReply, error) {
	opts := biz.MenusDOListOption{}
	opts.ListOptions.Pages = 0
	opts.ListOptions.Current = -1
	opts.ListOptions.PageSize = 20
	if in.Current > 0 {
		opts.ListOptions.Pages = in.Pages
		opts.ListOptions.Current = in.Current
		opts.ListOptions.PageSize = in.PageSize
	}

	opts.ListOptions.Init()
	opts.HasChildren = false
	opts.TargetId = in.TargetId

	if in.ParentId >= 0 {
		opts.ParentId = uint64(in.ParentId)
	}
	objList, err := s.mu.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	pageInfo := &v1.MenusPageInfo{
		Current:   objList.Current,
		PageSize:  objList.PageSize,
		Total:     objList.TotalCount,
		Pages:     objList.Pages,
		FirstFlag: objList.FirstFlag,
		LastFlag:  objList.LastFlag,
	}
	objs := make([]*v1.MenusInfoResponse, 0, len(objList.Items))
	for _, item := range objList.Items {
		titem := bizToMenusResponse(item)
		objs = append(objs, &titem)
	}
	return &v1.ListQyAdminMenusReply{PageInfo: pageInfo, Items: objs}, nil
}
func bizToMenusResponse(obj *biz.MenusDO) v1.MenusInfoResponse {
	objInfoRsp := v1.MenusInfoResponse{
		Id:       obj.ID,
		Name:     obj.Name,
		Url:      obj.Url,
		Blanked:  obj.Blanked,
		ParentId: obj.ParentId,
		TargetId: obj.TargetId,
	}
	if obj.Children != nil && len(obj.Children) > 0 {
		cobjList := make([]*v1.MenusInfoResponse, 0, len(obj.Children))
		for _, cobj := range obj.Children {
			citme := bizToMenusResponse(cobj)
			cobjList = append(cobjList, &citme)
		}
		objInfoRsp.Children = cobjList
	}
	return objInfoRsp
}
