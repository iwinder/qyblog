package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/admin/v1"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"strconv"
)

// CreateMenusAdmin 创建用户
func (s *BlogAdminUserService) CreateQyAdminMenusAdmin(ctx context.Context, in *v1.CreateQyAdminMenusAdminRequest) (*v1.CreateQyAdminMenusAdminReply, error) {
	objDO := &biz.MenusAdminDO{
		Name:           in.Name,
		BreadcrumbName: in.BreadcrumbName,
		Identifier:     in.Identifier,
		ParentId:       in.ParentId,
		Icon:           in.Icon,
		MType:          int(in.Type),
		Path:           in.Path,
		Redirect:       in.Redirect,
		Component:      in.Component,
		Sort:           int(in.Sort),
	}
	objDO.StatusFlag = int(in.StatusFlag)
	obj, err := s.mc.Create(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.CreateQyAdminMenusAdminReply{Id: obj.ID}, nil
}

// UpdateMenusAdmin 更新用户
func (s *BlogAdminUserService) UpdateQyAdminMenusAdmin(ctx context.Context, in *v1.UpdateQyAdminMenusAdminRequest) (*v1.UpdateQyAdminMenusAdminReply, error) {
	objDO := &biz.MenusAdminDO{
		ObjectMeta:     metaV1.ObjectMeta{ID: in.Id},
		Name:           in.Name,
		BreadcrumbName: in.BreadcrumbName,
		Identifier:     in.Identifier,
		ParentId:       in.ParentId,
		Icon:           in.Icon,
		MType:          int(in.Type),
		Path:           in.Path,
		Redirect:       in.Redirect,
		Component:      in.Component,
		Sort:           int(in.Sort),
	}
	objDO.StatusFlag = int(in.StatusFlag)
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
	return &v1.GetQyAdminMenusAdminReply{Data: &u}, nil
}

func (s *BlogAdminUserService) GetMyMenusAdminInfo(ctx context.Context, in *v1.GetMyMenusAdminInfoReq) (*v1.GetMyMenusAdminInfoReply, error) {
	var uId uint64
	var aerr error
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwtV4.MapClaims)
		if c["RoleIds"] == nil {
			return nil, ErrAuthFailed
		}
		uId, aerr = strconv.ParseUint(c["RoleIds"].(string), 10, 64)
		if aerr != nil {
			return nil, aerr
		}

	}

	objList, err := s.mc.FindAllByRoleID(ctx, uId)
	if err != nil {
		return nil, err
	}
	objs := make([]*v1.MenusAdminInfoResponse, 0, len(objList))
	for _, item := range objList {
		titem := bizToMenusAdminResponse(item)
		objs = append(objs, &titem)
	}
	return &v1.GetMyMenusAdminInfoReply{Items: objs}, nil
}

// ListMenusAdmin 获取用户列表
func (s *BlogAdminUserService) ListQyAdminMenusAdmin(ctx context.Context, in *v1.ListQyAdminMenusAdminRequest) (*v1.ListQyAdminMenusAdminReply, error) {
	opts := biz.MenusAdminDOListOption{}
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
	opts.Redirect = in.Redirect
	if in.HasChildren {
		opts.HasChildren = true
	}
	if in.Type > 0 {
		opts.MType = int(in.Type)
	}
	if in.ParentId >= 0 {
		opts.ParentId = uint64(in.ParentId)
	}
	objList, err := s.mc.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	pageInfo := &v1.MenusAdmPageInfo{
		Current:   objList.Current,
		PageSize:  objList.PageSize,
		Total:     objList.TotalCount,
		Pages:     objList.Pages,
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
		Id:             obj.ID,
		Name:           obj.Name,
		BreadcrumbName: obj.BreadcrumbName,
		Identifier:     obj.Identifier,
		ParentId:       obj.ParentId,
		Icon:           obj.Icon,
		Type:           int32(obj.MType),
		Path:           obj.Path,
		Redirect:       obj.Redirect,
		Component:      obj.Component,
		Sort:           int32(obj.Sort),
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
