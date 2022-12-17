package service

import (
	"context"
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/web/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (b *BlogWebApiService) CreateQyWebComment(ctx context.Context, in *v1.CreateQyWebCommentRequest) (*v1.CreateQyWebCommentReply, error) {
	opt := &biz.CommentDO{
		AgentId:    in.AgentId,
		MemberName: in.MemberName,
		Email:      in.Email,
		Url:        in.Url,
		ParentId:   in.ParentId,
		Content:    in.Content,
	}
	agent, ip := GetHeardInfo(ctx)
	opt.Agent = agent
	opt.Ip = ip
	opt.StatusFlag = 2
	opt.EmailState = 1
	data, err := b.ctu.CreateCommentWeb(ctx, opt, b.conf)
	if err != nil {
		return nil, err
	}
	return &v1.CreateQyWebCommentReply{Id: data.ID}, err
}

func (b *BlogWebApiService) ListQyWebComment(ctx context.Context, in *v1.ListQyWebCommentRequest) (*v1.ListQyWebCommentReply, error) {
	opts := biz.CommentContentDOListOption{}
	opts.ListOptions.Pages = 0
	opts.ListOptions.Current = -1
	opts.ListOptions.PageSize = 20
	if in.Current > 0 {
		opts.ListOptions.Pages = in.Pages
		opts.ListOptions.Current = in.Current
		opts.ListOptions.PageSize = in.PageSize
	}
	opts.AgentId = in.AgentId
	opts.RootId = in.RootId
	opts.StatusFlag = 1
	opts.ListOptions.Init()
	opts.IsWeb = true
	opts.IsChild = false
	if opts.RootId > 0 {
		opts.IsChild = true
	}
	objList := b.ctu.ListAllForWeb(ctx, opts)
	pageInfo := &v1.WebCommentPageInfo{
		Current:   objList.Current,
		PageSize:  objList.PageSize,
		Total:     objList.TotalCount,
		Pages:     objList.Pages,
		FirstFlag: objList.FirstFlag,
		LastFlag:  objList.LastFlag,
	}
	objs := make([]*v1.WebCommentInfoResponse, 0, len(objList.Items))
	for _, item := range objList.Items {
		titem := bizToWebCommentResponse(item)
		objs = append(objs, titem)
	}
	return &v1.ListQyWebCommentReply{PageInfo: pageInfo, Count: objList.Agent.Count, Items: objs}, nil
}

func bizToWebCommentResponse(obj *biz.CommentContentDO) *v1.WebCommentInfoResponse {
	objInfoRsp := &v1.WebCommentInfoResponse{
		Id:             obj.ID,
		StatusFlag:     int32(obj.StatusFlag),
		AgentId:        obj.AgentId,
		MemberId:       obj.MemberId,
		AtMemberIds:    obj.AtMemberIds,
		MemberName:     obj.MemberName,
		RootId:         obj.RootId,
		Content:        obj.Content,
		ParentUserName: obj.ParentUserName,
		Avatar:         obj.Avatar,
		CreatedAt:      timestamppb.New(obj.CreatedAt.Local()),
		Count:          obj.Count,
		RootCount:      obj.RootCount,
	}
	return objInfoRsp
}
