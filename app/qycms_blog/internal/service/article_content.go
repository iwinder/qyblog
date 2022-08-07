package service

import (
	"context"
	v1 "github.com/iwinder/qingyucms/api/qycms_blog/admin/v1"
	"github.com/iwinder/qingyucms/app/qycms_blog/internal/biz"
)

type ArticleContentService struct {
	uc *biz.ArticleContentUsecase
}

func NewArticleContentService(uc *biz.ArticleContentUsecase) *ArticleContentService {
	return &ArticleContentService{uc: uc}
}

// CreateArticleContent 创建
func (s *ArticleContentService) CreateArticleContent(ctx context.Context, in *v1.CreateArticleRequest) error {
	data := &biz.ArticleContentDO{
		Status:      int(in.Status),
		Atype:       int(in.Atype),
		Content:     in.Content,
		ContentHtml: in.ContentHtml,
	}
	data.ID = in.Uid
	data, err := s.uc.Create(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

// UpdateArticleContent 更新
func (s *ArticleContentService) UpdateArticleContent(ctx context.Context, in *v1.UpdateArticleRequest) error {
	data := &biz.ArticleContentDO{
		Status:      int(in.Status),
		Atype:       int(in.Atype),
		Content:     in.Content,
		ContentHtml: in.ContentHtml,
	}
	data.ID = in.Uid
	data, err := s.uc.Update(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (s *ArticleContentService) findOne(ctx context.Context, in *v1.GetArticleRequest) (*v1.GetArticleReply, error) {
	data, err := s.uc.FindOneByID(ctx, in.Uid)
	if err != nil {
		return nil, err
	}
	u := articleContentResponse(data)
	return &v1.GetArticleReply{Content: &u}, nil
}

func articleContentResponse(data *biz.ArticleContentDO) v1.ArticleInfoResponse {
	dataInfoRsp := v1.ArticleInfoResponse{
		Uid:         data.ID,
		Content:     data.Content,
		ContentHtml: data.ContentHtml,
	}

	return dataInfoRsp
}
