package service

import (
	"context"
	v1 "github.com/iwinder/qyblog/api/qycms_bff/admin/v1"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qyblog/internal/qycms_blog/biz"
)

// CreateQyAdminRole 创建
func (s *BlogAdminUserService) CreateQyAdminRole(ctx context.Context, in *v1.CreateQyAdminRoleRequest) (*v1.CreateQyAdminRoleReply, error) {
	objDO := &biz.RoleDO{
		Name:       in.Name,
		Identifier: in.Identifier,
	}
	obj, err := s.rc.Create(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.CreateQyAdminRoleReply{Id: obj.ID}, nil
}

// UpdateRole 更新用户
func (s *BlogAdminUserService) UpdateQyAdminRole(ctx context.Context, in *v1.UpdateQyAdminRoleRequest) (*v1.UpdateQyAdminRoleReply, error) {
	objDO := &biz.RoleDO{
		ObjectMeta: metaV1.ObjectMeta{
			ID: in.Id,
		},
		Name:       in.Name,
		Identifier: in.Identifier,
	}
	//if in.MenusAdmin != nil && len(in.MenusAdmin) > 0 {
	//	menus := make([]*biz.MenusAdminDO, 0, len(in.MenusAdmin))
	//	for _, item := range in.MenusAdmin {
	//		menus = append(menus, &biz.MenusAdminDO{
	//			ObjectMeta: metaV1.ObjectMeta{
	//				ID: item.Id,
	//			},
	//			Path: item.Path,
	//			Name: item.Name,
	//		})
	//	}
	//	objDO.MenusAdmins = menus
	//}
	//if in.Apis != nil && len(in.Apis) > 0 {
	//	apiDOS := make([]*biz.ApiDO, 0, len(in.Apis))
	//	for _, item := range in.Apis {
	//		apiDOS = append(apiDOS, &biz.ApiDO{
	//			ObjectMeta: metaV1.ObjectMeta{
	//				ID: item.Id,
	//			},
	//			ApiGroup:   item.ApiGroup,
	//			Identifier: item.Identifier,
	//			Method:     item.Method,
	//			Path:       item.Path,
	//		})
	//	}
	//	objDO.Apis = apiDOS
	//}

	obj, err := s.rc.Update(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateQyAdminRoleReply{Id: obj.ID}, nil
}

// DeleteRole 根据ID删除用户
func (s *BlogAdminUserService) DeleteQyAdminRole(ctx context.Context, in *v1.DeleteQyAdminRoleRequest) (*v1.DeleteQyAdminRoleReply, error) {
	err := s.rc.Delete(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteQyAdminRoleReply{}, nil
}

// DeleteRoles 根据ID批量删除用户
func (s *BlogAdminUserService) DeleteQyAdminRoles(ctx context.Context, in *v1.DeleteQyAdminRolesRequest) (*v1.DeleteQyAdminRolesReply, error) {
	err := s.rc.DeleteList(ctx, in.Ids)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteQyAdminRolesReply{}, nil
}

// GetRole 根据ID获取用户信息
func (s *BlogAdminUserService) GetQyAdminRole(ctx context.Context, in *v1.GetQyAdminRoleRequest) (*v1.GetQyAdminRoleReply, error) {
	obj, err := s.rc.FindOneByID(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	u := bizToRoleResponse(obj)
	return &v1.GetQyAdminRoleReply{Data: &u}, nil
}

// ListRole 获取用户列表
func (s *BlogAdminUserService) ListQyAdminRole(ctx context.Context, in *v1.ListQyAdminRoleRequest) (*v1.ListQyAdminRoleReply, error) {
	opts := biz.RoleDOListOption{}
	opts.ListOptions.Pages = 0
	opts.ListOptions.Current = -1
	opts.ListOptions.PageSize = 20
	if in.Current > 0 {
		opts.ListOptions.Pages = in.Pages
		opts.ListOptions.Current = in.Current
		opts.ListOptions.PageSize = in.PageSize
	}
	opts.Name = in.Name
	opts.ListOptions.Init()
	objList, err := s.rc.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	pageInfo := &v1.RolePageInfo{
		Current:   objList.Current,
		PageSize:  objList.PageSize,
		Total:     objList.TotalCount,
		Pages:     objList.Pages,
		FirstFlag: objList.FirstFlag,
		LastFlag:  objList.LastFlag,
	}
	objs := make([]*v1.RoleInfoResponse, 0, len(objList.Items))
	for _, item := range objList.Items {
		titem := bizToRoleResponse(item)
		objs = append(objs, &titem)
	}
	return &v1.ListQyAdminRoleReply{PageInfo: pageInfo, Items: objs}, nil
}

func (s *BlogAdminUserService) SaveQyAdminRoleMenus(ctx context.Context, in *v1.SaveRoleMenusRequest) (*v1.SaveRoleMenusReply, error) {
	param := &biz.RoleDO{
		MenusIDs: in.MenusIDs,
	}
	param.ID = in.Id
	err := s.rm.UpdateRoleForUser(ctx, param)
	if err != nil {
		return nil, err
	}
	return &v1.SaveRoleMenusReply{}, nil
}
func (s *BlogAdminUserService) SaveQyAdminRoleApis(ctx context.Context, in *v1.SaveRoleApisRequest) (*v1.SaveRoleApisReply, error) {

	apis := make([]*biz.ApiDO, 0, len(in.Apis))
	for _, obj := range in.Apis {
		apis = append(apis, &biz.ApiDO{
			ObjectMeta: metaV1.ObjectMeta{
				ID: obj.Id,
			},
			ApiGroup:    obj.ApiGroup,
			Method:      obj.Method,
			Path:        obj.Path,
			Description: obj.Description,
			Identifier:  obj.Identifier,
		})
	}
	param := &biz.RoleDO{
		ApiIds: in.ApiIDs,
		Apis:   apis,
	}
	param.ID = in.Id
	err := s.ra.UpdateApisForRole(ctx, param)
	if err != nil {
		return nil, err
	}
	return &v1.SaveRoleApisReply{}, nil
}
func bizToRoleResponse(obj *biz.RoleDO) v1.RoleInfoResponse {
	objInfoRsp := v1.RoleInfoResponse{
		Id:         obj.ID,
		Name:       obj.Name,
		Identifier: obj.Identifier,
	}
	if obj.Apis != nil && len(obj.Apis) > 0 {
		aobjRes := make([]*v1.RApiInfoResponse, 0, len(obj.Apis))
		for _, item := range obj.Apis {
			aobjRes = append(aobjRes, &v1.RApiInfoResponse{
				Id:          item.ID,
				ApiGroup:    item.ApiGroup,
				Identifier:  obj.Identifier,
				Method:      item.Method,
				Path:        item.Path,
				Description: item.Description,
			})
		}
		objInfoRsp.Apis = aobjRes
	}
	if obj.MenusIDs != nil && len(obj.MenusIDs) > 0 {
		objInfoRsp.MenusIDs = obj.MenusIDs
	}
	if obj.ApiIds != nil && len(obj.ApiIds) > 0 {
		objInfoRsp.ApiIDs = obj.ApiIds
	}
	return objInfoRsp
}
