package service

import (
	"context"
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/admin/v1"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
)

// CreateApi 创建用户
func (s *BlogAdminUserService) CreateQyAdminApiGroup(ctx context.Context, in *v1.CreateQyAdminApiGroupRequest) (*v1.CreateQyAdminApiGroupReply, error) {
	objDO := &biz.ApiGroupDO{
		Name:       in.Name,
		Identifier: in.Identifier,
	}
	obj, err := s.acg.Create(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.CreateQyAdminApiGroupReply{Id: obj.ID}, nil
}

// UpdateApi 更新用户
func (s *BlogAdminUserService) UpdateQyAdminApiGroup(ctx context.Context, in *v1.UpdateQyAdminApiGroupRequest) (*v1.UpdateQyAdminApiGroupReply, error) {
	objDO := &biz.ApiGroupDO{
		ObjectMeta: metaV1.ObjectMeta{ID: in.Id},
		Name:       in.Name,
		Identifier: in.Identifier,
	}
	obj, err := s.acg.Update(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateQyAdminApiGroupReply{Id: obj.ID}, nil
}

// DeleteApis 根据ID批量删除用户
func (s *BlogAdminUserService) DeleteQyAdminApiGroups(ctx context.Context, in *v1.DeleteQyAdminApiGroupRequest) (*v1.DeleteQyAdminApiGroupReply, error) {
	err := s.acg.DeleteList(ctx, in.Ids)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteQyAdminApiGroupReply{}, nil
}

// GetApi 根据ID获取用户信息
func (s *BlogAdminUserService) GetQyAdminApiGroup(ctx context.Context, in *v1.GetQyAdminApiGroupRequest) (*v1.GetQyAdminApiGroupReply, error) {
	obj, err := s.acg.FindOneByID(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	u := bizToApiGroupResponse(obj)
	return &v1.GetQyAdminApiGroupReply{Data: &u}, nil
}

// ListApi 获取用户列表
func (s *BlogAdminUserService) ListQyAdminApiGroup(ctx context.Context, in *v1.ListQyAdminApiGroupRequest) (*v1.ListQyAdminApiGroupReply, error) {
	opts := biz.ApiGroupDOListOption{}
	opts.ListOptions.Pages = 0
	opts.ListOptions.Current = -1
	opts.ListOptions.PageSize = 20
	if in.Current > 0 {
		opts.ListOptions.Pages = in.Pages
		opts.ListOptions.Current = in.Current
		opts.ListOptions.PageSize = in.PageSize
	}

	opts.ListOptions.Init()
	opts.Name = in.Name
	opts.Identifier = in.Identifier
	objList, err := s.acg.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	pageInfo := &v1.APIGroupPageInfo{
		Current:   objList.Current,
		PageSize:  objList.PageSize,
		Total:     objList.TotalCount,
		Pages:     objList.Pages,
		FirstFlag: objList.FirstFlag,
		LastFlag:  objList.LastFlag,
	}
	objs := make([]*v1.ApiGroupInfoResponse, 0, len(objList.Items))
	for _, item := range objList.Items {
		titem := bizToApiGroupResponse(item)
		objs = append(objs, &titem)
	}
	return &v1.ListQyAdminApiGroupReply{PageInfo: pageInfo, Items: objs}, nil
}
func bizToApiGroupResponse(obj *biz.ApiGroupDO) v1.ApiGroupInfoResponse {
	objInfoRsp := v1.ApiGroupInfoResponse{
		Id:         obj.ID,
		Name:       obj.Name,
		Identifier: obj.Identifier,
	}
	return objInfoRsp
}
