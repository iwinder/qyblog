package service

import (
	"context"
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/admin/v1"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *BlogAdminUserService) CreateQyAdminComment(ctx context.Context, in *v1.CreateQyAdminCommentRequest) (*v1.CreateQyAdminCommentReply, error) {
	objDO := &biz.CommentContentDO{
		AgentId:     in.AgentId,
		AtMemberIds: in.AtMemberIds,
		Email:       in.Email,
		Url:         in.Url,
		RootId:      in.RootId,
		Content:     in.Content,
		Meta:        in.Meta,
	}
	if objDO.StatusFlag == 0 {
		objDO.StatusFlag = 1
	}
	userAgent, ip := GetHeardInfo(ctx)
	user := GetUserInfo(ctx)
	if user != nil {
		objDO.MemberId = user.ID
		objDO.MemberName = user.Nickname
		auser, _ := s.uc.FindOneByID(ctx, user.ID)
		if auser != nil {
			objDO.Email = auser.Email
		}
	}
	objDO.Agent = userAgent
	objDO.Ip = ip
	obj, err := s.cc.CreateCommentContent(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.CreateQyAdminCommentReply{Id: obj.ID}, nil
}

// UpdateQyAdminComment 更新
func (s *BlogAdminUserService) UpdateQyAdminComment(ctx context.Context, in *v1.UpdateQyAdminCommentRequest) (*v1.UpdateQyAdminCommentReply, error) {
	objDO := &biz.CommentContentDO{
		ObjectMeta: metaV1.ObjectMeta{
			ID: in.Id,
		},
		Content: in.Content,
		RootId:  in.RootId,
	}
	obj, err := s.cc.Update(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateQyAdminCommentReply{Id: obj.ID}, nil
}

func (s *BlogAdminUserService) UpdateQyAdminCommentState(ctx context.Context, in *v1.UpdateQyAdminCommentStateRequest) (*v1.UpdateQyAdminCommentStateReply, error) {
	err := s.cc.UpdaeStateByIDs(ctx, in.Ids, int(in.StatusFlag))
	if err != nil {
		return nil, err
	}
	return &v1.UpdateQyAdminCommentStateReply{}, nil
}

// GetQyAdminCommentCount 获取总数
func (s *BlogAdminUserService) GetQyAdminCommentCount(ctx context.Context, in *v1.GetQyAdminCommentCountRequest) (*v1.GetQyAdminCommentCountReply, error) {
	data := s.cc.CountAll(ctx)
	return &v1.GetQyAdminCommentCountReply{
		NowTotal:     data.NowTotal,
		PendingTotal: data.PendingTotal,
		Total:        data.Total,
	}, nil
}

// DeleteQyAdminComment 根据ID批量删除
func (s *BlogAdminUserService) DeleteQyAdminComment(ctx context.Context, in *v1.DeleteQyAdminCommentRequest) (*v1.DeleteQyAdminCommentReply, error) {
	err := s.cc.DeleteList(ctx, in.Ids)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteQyAdminCommentReply{}, nil
}

// ListQyAdminComment 获取列表
func (s *BlogAdminUserService) ListQyAdminComment(ctx context.Context, in *v1.ListQyAdminCommentRequest) (*v1.ListQyAdminCommentReply, error) {
	opts := biz.CommentContentDOListOption{}
	opts.ListOptions.Pages = 0
	opts.ListOptions.Current = -1
	opts.ListOptions.PageSize = 20
	if in.Current > 0 {
		opts.ListOptions.Pages = in.Pages
		opts.ListOptions.Current = in.Current
		opts.ListOptions.PageSize = in.PageSize
	}
	opts.Content = in.SearchText
	opts.ListOptions.Init()
	objList, err := s.cc.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	pageInfo := &v1.CommentPageInfo{
		Current:   objList.Current,
		PageSize:  objList.PageSize,
		Total:     objList.TotalCount,
		Pages:     objList.Pages,
		FirstFlag: objList.FirstFlag,
		LastFlag:  objList.LastFlag,
	}
	objs := make([]*v1.CommentInfoResponse, 0, len(objList.Items))
	for _, item := range objList.Items {
		titem := bizToCommentResponse(item)
		objs = append(objs, &titem)
	}
	return &v1.ListQyAdminCommentReply{PageInfo: pageInfo, Items: objs}, nil
}

func bizToCommentResponse(obj *biz.CommentContentDO) v1.CommentInfoResponse {
	objInfoRsp := v1.CommentInfoResponse{
		Id:             obj.ID,
		StatusFlag:     int32(obj.StatusFlag),
		AgentId:        obj.AgentId,
		MemberId:       obj.MemberId,
		AtMemberIds:    obj.AtMemberIds,
		Agent:          obj.Agent,
		MemberName:     obj.MemberName,
		Ip:             obj.Ip,
		Email:          obj.Email,
		Url:            obj.Url,
		RootId:         obj.RootId,
		Content:        obj.Content,
		Meta:           obj.Meta,
		ParentUserName: obj.ParentUserName,
		ObjTitle:       obj.ObjTitle,
		ObjLink:        obj.ObjLink,
		Avatar:         obj.Avatar,
		CreatedAt:      timestamppb.New(obj.CreatedAt.Local()),
	}
	return objInfoRsp
}
