package service

import (
	"context"
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/admin/v1"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
)

// CreateMenusAdmin 创建用户
func (s *BlogAdminUserService) CreateQyAdminMenusAdmin(ctx context.Context, in *v1.CreateQyAdminMenusAdminRequest) (*v1.CreateQyAdminMenusAdminReply, error) {
	objDO := &biz.MenusAdminDO{
		Level:     int(in.Level),
		ParentId:  in.ParentId,
		Path:      in.Path,
		Name:      in.Name,
		Hidden:    in.Hidden,
		Component: in.Component,
		Sort:      int(in.Sort),
	}
	obj, err := s.mc.Create(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.CreateQyAdminMenusAdminReply{Id: obj.ID}, nil
}

// UpdateMenusAdmin 更新用户
func (s *BlogAdminUserService) UpdateQyAdminMenusAdmin(ctx context.Context, in *v1.UpdateQyAdminMenusAdminRequest) (*v1.UpdateQyAdminMenusAdminReply, error) {
	objDO := &biz.MenusAdminDO{
		ObjectMeta: metaV1.ObjectMeta{ID: in.Id},
		Level:      int(in.Level),
		ParentId:   in.ParentId,
		Path:       in.Path,
		Name:       in.Name,
		Hidden:     in.Hidden,
		Component:  in.Component,
		Sort:       int(in.Sort),
	}
	obj, err := s.mc.Update(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateQyAdminMenusAdminReply{Id: obj.ID}, nil
}

// DeleteMenusAdmin 根据ID删除用户
func (s *BlogAdminUserService) DeleteQyAdminMenusAdmin(ctx context.Context, in *v1.DeleteQyAdminMenusAdminRequest) (*v1.DeleteQyAdminMenusAdminReply, error) {
	err := s.mc.Delete(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteQyAdminMenusAdminReply{}, nil
}

// DeleteMenusAdmins 根据ID批量删除用户
func (s *BlogAdminUserService) DeleteQyAdminMenusAdmins(ctx context.Context, in *v1.DeleteQyAdminMenusAdminsRequest) (*v1.DeleteQyAdminMenusAdminsReply, error) {
	err := s.mc.DeleteList(ctx, in.Ids)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteQyAdminMenusAdminsReply{}, nil
}

// GetMenusAdmin 根据ID获取用户信息
func (s *BlogAdminUserService) GetQyAdminMenusAdmin(ctx context.Context, in *v1.GetQyAdminMenusAdminRequest) (*v1.GetQyAdminMenusAdminReply, error) {
	obj, err := s.mc.FindOneByID(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	u := bizToMenusAdminResponse(obj)
	return &v1.GetQyAdminMenusAdminReply{Info: &u}, nil
}

// ListMenusAdmin 获取用户列表
func (s *BlogAdminUserService) ListQyAdminMenusAdmin(ctx context.Context, in *v1.ListQyAdminMenusAdminRequest) (*v1.ListQyAdminMenusAdminReply, error) {
	opts := biz.MenusAdminDOListOption{}
	opts.ListOptions.Pages = 0
	opts.ListOptions.Page = -1
	opts.ListOptions.PageSize = 20
	if in.PageInfo != nil {
		opts.ListOptions.Pages = int64(in.PageInfo.Pages)
		opts.ListOptions.Page = int64(in.PageInfo.Page)
		opts.ListOptions.PageSize = int64(in.PageInfo.Size)
	}

	opts.ListOptions.Init()
	objList, err := s.mc.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	pageInfo := &v1.MenusAdmPageInfo{
		Page:      uint64(objList.Pages),
		Size:      uint64(objList.PageSize),
		Total:     uint64(objList.TotalCount),
		Pages:     uint64(objList.Pages),
		FirstFlag: objList.FirstFlag,
		LastFlag:  objList.LastFlag,
	}
	objs := make([]*v1.MenusAdminInfoResponse, 0, len(objList.Items))
	for _, item := range objList.Items {
		titem := bizToMenusAdminResponse(item)
		objs = append(objs, &titem)
	}
	return &v1.ListQyAdminMenusAdminReply{PageInfo: pageInfo, Items: objs}, nil
}
func bizToMenusAdminResponse(obj *biz.MenusAdminDO) v1.MenusAdminInfoResponse {
	objInfoRsp := v1.MenusAdminInfoResponse{
		Id:        obj.ID,
		Name:      obj.Name,
		ParentId:  obj.ParentId,
		Path:      obj.Path,
		Hidden:    obj.Hidden,
		Component: obj.Component,
		Sort:      int32(obj.Sort),
		Level:     uint32(obj.Level),
	}
	if obj.Children != nil && len(obj.Children) > 0 {
		cobjList := make([]*v1.MenusAdminInfoResponse, 0, len(obj.Children))
		for _, cobj := range obj.Children {
			citme := bizToMenusAdminResponse(cobj)
			cobjList = append(cobjList, &citme)
		}
		objInfoRsp.Children = cobjList
	}
	return objInfoRsp
}
