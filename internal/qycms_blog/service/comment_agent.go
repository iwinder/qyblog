package service

import (
	"context"
	v1 "github.com/iwinder/qingyucms/api/qycms_comments/admin/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
)

// CommentAgentService is a greeter service.
type CommentAgentService struct {
	v1.UnimplementedCommentAgentServer

	uc *biz.CommentAgentUsecase
}

// NewGreeterService new a greeter service.
func NewCommentAgentService(uc *biz.CommentAgentUsecase) *CommentAgentService {
	return &CommentAgentService{uc: uc}
}

// CreateCommentAgent implements helloworld.GreeterServer.
func (s *CommentAgentService) CreateCommentAgent(ctx context.Context, in *v1.CreateCommentAgentRequest) (*v1.CreateCommentAgentReply, error) {
	data, err := s.uc.CreateCommentAgent(ctx, &biz.CommentAgentDO{
		ObjId:     in.ObjId,
		ObjType:   in.ObjType,
		MemberId:  in.MemberId,
		Count:     in.Count,
		RootCount: in.RootCount,
		AllCount:  in.AllCount,
		Attrs:     in.Attrs,
	})
	if err != nil {
		return nil, err
	}
	u := dataResponse(data)
	return &v1.CreateCommentAgentReply{Content: &u}, nil
}

// UpdateArticle 更新
func (s *CommentAgentService) UpdateArticle(ctx context.Context, in *v1.UpdateCommentAgentRequest) (*v1.UpdateCommentAgentReply, error) {
	dataDO := &biz.CommentAgentDO{
		ObjId:     in.ObjId,
		ObjType:   in.ObjType,
		MemberId:  in.MemberId,
		Count:     in.Count,
		RootCount: in.RootCount,
		AllCount:  in.AllCount,
		Attrs:     in.Attrs,
	}
	data, err := s.uc.Update(ctx, dataDO)
	if err != nil {
		return nil, err
	}
	u := dataResponse(data)
	return &v1.UpdateCommentAgentReply{Content: &u}, nil
}

// GetArticle 通过ID获取
func (s *CommentAgentService) GetArticle(ctx context.Context, in *v1.GetCommentAgentRequest) (*v1.GetCommentAgentReply, error) {
	data, err := s.uc.FindByID(ctx, in.Uid)
	if err != nil {
		return nil, err
	}
	u := dataResponse(data)
	return &v1.GetCommentAgentReply{Content: &u}, nil
}

func dataResponse(data *biz.CommentAgentDO) v1.CommentAgentResponse {
	dataInfoRsp := v1.CommentAgentResponse{
		Uid:       data.ID,
		ObjId:     data.ObjId,
		ObjType:   int32(data.ObjType),
		MemberId:  data.MemberId,
		Count:     data.Count,
		RootCount: data.RootCount,
		AllCount:  data.AllCount,
		Attrs:     data.Attrs,
	}

	return dataInfoRsp
}
