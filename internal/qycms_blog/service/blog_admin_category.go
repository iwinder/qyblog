package service

import (
	"context"
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/admin/v1"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
)

func (s *BlogAdminUserService) CreateQyAdminCategory(ctx context.Context, in *v1.CreateQyAdminCategoryRequest) (*v1.CreateQyAdminCategoryReply, error) {
	objDO := &biz.CategoryDO{
		Name:       in.Name,
		Identifier: in.Identifier,
		ParentId:   in.ParentId,
	}
	obj, err := s.cu.Create(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.CreateQyAdminCategoryReply{Id: obj.ID}, nil
}

// UpdateQyAdminCategory 更新
func (s *BlogAdminUserService) UpdateQyAdminCategory(ctx context.Context, in *v1.UpdateQyAdminCategoryRequest) (*v1.UpdateQyAdminCategoryReply, error) {
	objDO := &biz.CategoryDO{
		ObjectMeta: metaV1.ObjectMeta{
			ID: in.Id,
		},
		Name:       in.Name,
		Identifier: in.Identifier,
		ParentId:   in.ParentId,
	}

	obj, err := s.cu.Update(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateQyAdminCategoryReply{Id: obj.ID}, nil
}

// DeleteQyAdminCategory 根据ID删除用户
func (s *BlogAdminUserService) DeleteQyAdminCategory(ctx context.Context, in *v1.DeleteQyAdminCategoryRequest) (*v1.DeleteQyAdminCategoryReply, error) {
	err := s.cu.DeleteList(ctx, in.Ids)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteQyAdminCategoryReply{}, nil
}

// ListQyAdminCategory 获取列表
func (s *BlogAdminUserService) ListQyAdminCategory(ctx context.Context, in *v1.ListQyAdminCategoryRequest) (*v1.ListQyAdminCategoryReply, error) {
	opts := biz.CategoryDOListOption{}
	opts.ListOptions.Pages = 0
	opts.ListOptions.Current = -1
	opts.ListOptions.PageSize = 20
	if in.Current > 0 {
		opts.ListOptions.Pages = in.Pages
		opts.ListOptions.Current = in.Current
		opts.ListOptions.PageSize = in.PageSize
	}
	opts.Name = in.Name
	opts.ParentId = in.ParentId
	opts.ListOptions.Init()
	objList, err := s.cu.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	pageInfo := &v1.CategoryPageInfo{
		Current:   objList.Current,
		PageSize:  objList.PageSize,
		Total:     objList.TotalCount,
		Pages:     objList.Pages,
		FirstFlag: objList.FirstFlag,
		LastFlag:  objList.LastFlag,
	}
	objs := make([]*v1.CategoryResponse, 0, len(objList.Items))
	for _, item := range objList.Items {
		titem := bizToCategoryResponse(item)
		objs = append(objs, &titem)
	}
	return &v1.ListQyAdminCategoryReply{PageInfo: pageInfo, Items: objs}, nil
}

func bizToCategoryResponse(obj *biz.CategoryDO) v1.CategoryResponse {
	objInfoRsp := v1.CategoryResponse{
		Id:         obj.ID,
		Name:       obj.Name,
		Identifier: obj.Identifier,
		ParentId:   obj.ParentId,
	}
	if obj.Children != nil && len(obj.Children) > 0 {
		cobjList := make([]*v1.CategoryResponse, 0, len(obj.Children))
		for _, cobj := range obj.Children {
			citme := bizToCategoryResponse(cobj)
			cobjList = append(cobjList, &citme)
		}
		objInfoRsp.Children = cobjList
	}
	return objInfoRsp
}
