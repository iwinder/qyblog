package service

import (
	"context"
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/admin/v1"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
)

// CreateQyAdminMenusAgent  创建
func (s *BlogAdminUserService) CreateQyAdminMenusAgent(ctx context.Context, in *v1.CreateQyAdminMenusAgentRequest) (*v1.CreateQyAdminMenusAgentReply, error) {
	objDO := &biz.MenusAgentDO{
		Name: in.Name,
	}
	objDO.Ftype = "USER"
	obj, err := s.mau.Create(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.CreateQyAdminMenusAgentReply{Id: obj.ID}, nil
}

// UpdateQyAdminMenusAgent 更新
func (s *BlogAdminUserService) UpdateQyAdminMenusAgent(ctx context.Context, in *v1.UpdateQyAdminMenusAgentRequest) (*v1.UpdateQyAdminMenusAgentReply, error) {
	objDO := &biz.MenusAgentDO{
		ObjectMeta: metaV1.ObjectMeta{ID: in.Id},
		Name:       in.Name,
	}
	obj, err := s.mau.Update(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateQyAdminMenusAgentReply{Id: obj.ID}, nil
}

// DeleteQyAdminMenusAgents 删除
func (s *BlogAdminUserService) DeleteQyAdminMenusAgents(ctx context.Context, in *v1.DeleteQyAdminMenusAgentRequest) (*v1.DeleteQyAdminMenusAgentReply, error) {
	err := s.mau.DeleteList(ctx, in.Ids)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteQyAdminMenusAgentReply{}, nil
}

// ListQyAdminMenusAgent 列表
func (s *BlogAdminUserService) ListQyAdminMenusAgent(ctx context.Context, in *v1.ListQyAdminMenusAgentRequest) (*v1.ListQyAdminMenusAgentReply, error) {
	opts := biz.MenusAgentDOListOption{}
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
	objList, err := s.mau.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	pageInfo := &v1.MenusAgentPageInfo{
		Current:   objList.Current,
		PageSize:  objList.PageSize,
		Total:     objList.TotalCount,
		Pages:     objList.Pages,
		FirstFlag: objList.FirstFlag,
		LastFlag:  objList.LastFlag,
	}
	objs := make([]*v1.MenusAgentInfoResponse, 0, len(objList.Items))
	for _, item := range objList.Items {
		titem := bizToMenusAgentResponse(item)
		objs = append(objs, &titem)
	}
	return &v1.ListQyAdminMenusAgentReply{PageInfo: pageInfo, Items: objs}, nil
}

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
		Id:       obj.ParentId,
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

func bizToMenusAgentResponse(obj *biz.MenusAgentDO) v1.MenusAgentInfoResponse {
	objInfoRsp := v1.MenusAgentInfoResponse{
		Id:    obj.ID,
		Name:  obj.Name,
		Ftype: obj.Ftype,
	}
	return objInfoRsp
}
