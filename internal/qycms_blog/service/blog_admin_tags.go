package service

import (
	"context"
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/admin/v1"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
)

// CreateQyAdminTags 新增
func (s *BlogAdminUserService) CreateQyAdminTags(ctx context.Context, in *v1.CreateQyAdminTagsRequest) (*v1.CreateQyAdminTagsReply, error) {
	objDO := &biz.TagsDO{
		Name:        in.Name,
		Identifier:  in.Identifier,
		Description: in.Description,
	}
	obj, err := s.tau.Create(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.CreateQyAdminTagsReply{Id: obj.ID}, nil
}

// UpdateQyAdminTags 更新
func (s *BlogAdminUserService) UpdateQyAdminTags(ctx context.Context, in *v1.UpdateQyAdminTagsRequest) (*v1.UpdateQyAdminTagsReply, error) {
	objDO := &biz.TagsDO{
		ObjectMeta: metaV1.ObjectMeta{
			ID: in.Id,
		},
		Name:        in.Name,
		Identifier:  in.Identifier,
		Description: in.Description,
	}

	obj, err := s.tau.Update(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateQyAdminTagsReply{Id: obj.ID}, nil
}

// DeleteQyAdminTags 根据ID删除用户
func (s *BlogAdminUserService) DeleteQyAdminTags(ctx context.Context, in *v1.DeleteQyAdminTagsRequest) (*v1.DeleteQyAdminTagsReply, error) {
	err := s.tau.DeleteList(ctx, in.Ids)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteQyAdminTagsReply{}, nil
}

// ListQyAdminTags 获取列表
func (s *BlogAdminUserService) ListQyAdminTags(ctx context.Context, in *v1.ListQyAdminTagsRequest) (*v1.ListQyAdminTagsReply, error) {
	opts := biz.TagsDOListOption{}
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
	objList, err := s.tau.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	pageInfo := &v1.TagsPageInfo{
		Current:   objList.Current,
		PageSize:  objList.PageSize,
		Total:     objList.TotalCount,
		Pages:     objList.Pages,
		FirstFlag: objList.FirstFlag,
		LastFlag:  objList.LastFlag,
	}
	objs := make([]*v1.TagsResponse, 0, len(objList.Items))
	for _, item := range objList.Items {
		titem := bizToTagsResponse(item)
		objs = append(objs, &titem)
	}
	return &v1.ListQyAdminTagsReply{PageInfo: pageInfo, Items: objs}, nil
}

func bizToTagsResponse(obj *biz.TagsDO) v1.TagsResponse {
	objInfoRsp := v1.TagsResponse{
		Id:          obj.ID,
		Name:        obj.Name,
		Identifier:  obj.Identifier,
		Description: obj.Description,
	}

	return objInfoRsp
}
