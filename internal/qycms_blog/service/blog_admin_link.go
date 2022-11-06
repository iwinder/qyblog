package service

import (
	"context"
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/admin/v1"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
)

// CreateQyAdminLink 创建
func (s *BlogAdminUserService) CreateQyAdminLink(ctx context.Context, in *v1.CreateQyAdminLinkRequest) (*v1.CreateQyAdminLinkReply, error) {
	objDO := &biz.LinkDO{
		Name:        in.Name,
		Url:         in.Url,
		Description: in.Description,
		Ftype:       in.Ftype,
	}
	obj, err := s.lk.Create(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.CreateQyAdminLinkReply{Id: obj.ID}, nil
}

// UpdateQyAdminLink 更新
func (s *BlogAdminUserService) UpdateQyAdminLink(ctx context.Context, in *v1.UpdateQyAdminLinkRequest) (*v1.UpdateQyAdminLinkReply, error) {
	objDO := &biz.LinkDO{
		ObjectMeta: metaV1.ObjectMeta{
			ID: in.Id,
		},
		Name:        in.Name,
		Url:         in.Url,
		Description: in.Description,
		Ftype:       in.Ftype,
	}
	obj, err := s.lk.Update(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateQyAdminLinkReply{Id: obj.ID}, nil
}

// DeleteQyAdminLinks 根据ID批量删除
func (s *BlogAdminUserService) DeleteQyAdminLinks(ctx context.Context, in *v1.DeleteQyAdminLinksRequest) (*v1.DeleteQyAdminLinkReply, error) {
	err := s.lk.DeleteList(ctx, in.Ids)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteQyAdminLinkReply{}, nil
}

// ListQyAdminLink 获取列表
func (s *BlogAdminUserService) ListQyAdminLink(ctx context.Context, in *v1.ListQyAdminLinkRequest) (*v1.ListQyAdminLinkReply, error) {
	opts := biz.LinkDOListOption{}
	opts.ListOptions.Pages = 0
	opts.ListOptions.Current = -1
	opts.ListOptions.PageSize = 20
	if in.Current > 0 {
		opts.ListOptions.Pages = in.Pages
		opts.ListOptions.Current = in.Current
		opts.ListOptions.PageSize = in.PageSize
	}
	opts.Name = in.SearchText
	opts.ListOptions.Init()
	objList, err := s.lk.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	pageInfo := &v1.LinkPageInfo{
		Current:   objList.Current,
		PageSize:  objList.PageSize,
		Total:     objList.TotalCount,
		Pages:     objList.Pages,
		FirstFlag: objList.FirstFlag,
		LastFlag:  objList.LastFlag,
	}
	objs := make([]*v1.LinkInfo, 0, len(objList.Items))
	for _, item := range objList.Items {
		titem := bizToLinkResponse(item)
		objs = append(objs, &titem)
	}
	return &v1.ListQyAdminLinkReply{PageInfo: pageInfo, Items: objs}, nil
}

// CreateQyAdminShortLink 创建
func (s *BlogAdminUserService) CreateQyAdminShortLink(ctx context.Context, in *v1.CreateQyAdminShortLinkRequest) (*v1.CreateQyAdminShortLinkReply, error) {
	objDO := &biz.ShortLinkDO{
		Url:         in.Url,
		Description: in.Description,
		Identifier:  in.Identifier,
	}
	obj, err := s.slk.Create(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.CreateQyAdminShortLinkReply{Id: obj.ID}, nil
}

// UpdateQyAdminShortLink 更新
func (s *BlogAdminUserService) UpdateQyAdminShortLink(ctx context.Context, in *v1.UpdateQyAdminShortLinkRequest) (*v1.UpdateQyAdminShortLinkReply, error) {
	objDO := &biz.ShortLinkDO{
		ObjectMeta: metaV1.ObjectMeta{
			ID: in.Id,
		},
		Url:         in.Url,
		Description: in.Description,
		Identifier:  in.Identifier,
	}
	obj, err := s.slk.Update(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateQyAdminShortLinkReply{Id: obj.ID}, nil
}

// DeleteQyAdminShortLinks 根据ID批量删除
func (s *BlogAdminUserService) DeleteQyAdminShortLinks(ctx context.Context, in *v1.DeleteQyAdminShortLinksRequest) (*v1.DeleteQyAdminShortLinkReply, error) {
	err := s.slk.DeleteList(ctx, in.Ids)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteQyAdminShortLinkReply{}, nil
}

// ListQyAdminShortLink
func (s *BlogAdminUserService) ListQyAdminShortLink(ctx context.Context, in *v1.ListQyAdminShortLinkRequest) (*v1.ListQyAdminShortLinkReply, error) {
	opts := biz.ShortLinkDOListOption{}
	opts.ListOptions.Pages = 0
	opts.ListOptions.Current = -1
	opts.ListOptions.PageSize = 20
	if in.Current > 0 {
		opts.ListOptions.Pages = in.Pages
		opts.ListOptions.Current = in.Current
		opts.ListOptions.PageSize = in.PageSize
	}
	opts.Identifier = in.SearchText
	opts.ListOptions.Init()
	objList, err := s.slk.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	pageInfo := &v1.ShortLinkPageInfo{
		Current:   objList.Current,
		PageSize:  objList.PageSize,
		Total:     objList.TotalCount,
		Pages:     objList.Pages,
		FirstFlag: objList.FirstFlag,
		LastFlag:  objList.LastFlag,
	}
	objs := make([]*v1.ShortLinkInfo, 0, len(objList.Items))
	for _, item := range objList.Items {
		titem := bizToShortLinkResponse(item)
		objs = append(objs, &titem)
	}
	return &v1.ListQyAdminShortLinkReply{PageInfo: pageInfo, Items: objs}, nil
}

func bizToLinkResponse(obj *biz.LinkDO) v1.LinkInfo {
	objInfoRsp := v1.LinkInfo{
		Id:          obj.ID,
		Name:        obj.Name,
		Url:         obj.Url,
		Description: obj.Description,
		Ftype:       obj.Ftype,
	}
	return objInfoRsp
}

func bizToShortLinkResponse(obj *biz.ShortLinkDO) v1.ShortLinkInfo {
	objInfoRsp := v1.ShortLinkInfo{
		Id:          obj.ID,
		Url:         obj.Url,
		Description: obj.Description,
		Identifier:  obj.Identifier,
	}
	return objInfoRsp
}
