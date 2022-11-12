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

func bizToMenusAgentResponse(obj *biz.MenusAgentDO) v1.MenusAgentInfoResponse {
	objInfoRsp := v1.MenusAgentInfoResponse{
		Id:    obj.ID,
		Name:  obj.Name,
		Ftype: obj.Ftype,
	}
	return objInfoRsp
}
