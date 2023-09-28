package service

import (
	"context"
	v1 "github.com/iwinder/qyblog/api/qycms_bff/admin/v1"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qyblog/internal/qycms_blog/biz"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *BlogAdminUserService) CreateQyAdminComment(ctx context.Context, in *v1.CreateQyAdminCommentRequest) (*v1.CreateQyAdminCommentReply, error) {
	opt := &biz.CommentDO{
		AgentId:  in.AgentId,
		ParentId: in.ParentId,
		Content:  in.Content,
	}
	opt.StatusFlag = 1
	opt.EmailState = 2
	agent, ip := GetHeardInfo(ctx)
	opt.Agent = agent
	opt.Ip = ip
	user := GetUserInfo(ctx)
	if user != nil {
		opt.MemberId = user.ID
		opt.MemberName = user.Nickname
		auser, _ := s.uc.FindOneByID(ctx, user.ID)
		if auser != nil {
			opt.Email = auser.Email
		}
	}
	data, err := s.ctu.CreateComment(ctx, opt)
	if err != nil {
		return nil, err
	}
	return &v1.CreateQyAdminCommentReply{Id: data.ID}, nil
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
	err := s.ctu.UpdateCommentState(ctx, in.Ids, int(in.StatusFlag))
	if err != nil {
		return nil, err
	}
	return &v1.UpdateQyAdminCommentStateReply{}, nil
}
func (s *BlogAdminUserService) UpdateQyAdminCommentContent(ctx context.Context, in *v1.UpdateQyAdminCommentRequest) (*v1.UpdateQyAdminCommentReply, error) {
	data := &biz.CommentDO{}
	data.ID = in.Id
	data.Content = in.Content
	err := s.ctu.UpdateCommentComent(ctx, data)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateQyAdminCommentReply{}, nil
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
	err := s.ctu.DeleteList(ctx, in.Ids)
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
	opts.StatusFlag = int(in.StatusFlag)
	opts.ListOptions.Init()
	opts.Order = "status_flag desc ,created_at desc,id desc"
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
		Attrs:          obj.Attrs,
		ParentUserName: obj.ParentUserName,
		ObjTitle:       obj.ObjTitle,
		ObjLink:        obj.ObjLink,
		Avatar:         obj.Avatar,
		CreatedAt:      timestamppb.New(obj.CreatedAt.Local()),
	}
	return objInfoRsp
}
