package service

import (
	"context"
	v1 "github.com/iwinder/qingyucms/api/qycms_blog/admin/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"github.com/iwinder/qingyucms/internal/qycms_blog/conf"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ArticleService is a greeter service.
type ArticleService struct {
	v1.UnimplementedArticleServer
	auc *biz.ArticleUsecase
	//ac       *blog.ArticleContentUsecase
	authConf *conf.Auth
}

// NewArticleService new a greeter service.
func NewArticleService(auc *biz.ArticleUsecase, authConf *conf.Auth) *ArticleService {
	return &ArticleService{auc: auc,
		authConf: authConf,
	}
}

// CreateArticle implements server.CreateArticle. 创建
func (s *ArticleService) CreateArticle(ctx context.Context, in *v1.CreateArticleRequest) (*v1.CreateArticleReply, error) {
	data := &biz.ArticleDO{
		Title:          in.Title,
		PermaLink:      in.PermaLink,
		CanonicalLink:  in.CanonicalLink,
		Summary:        in.Summary,
		Thumbnail:      in.Thumbnail,
		Password:       in.Password,
		Atype:          int(in.Atype),
		CategoryId:     in.CategoryId,
		CommentAgentId: in.CommentAgentId,
		Published:      in.Published,
		ViewCount:      in.ViewCount,
		LikeCount:      in.LikeCount,
		HateCount:      in.HateCount,
	}
	data, err := s.auc.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	in.Uid = data.ID
	u := articleResponse(data)
	return &v1.CreateArticleReply{Content: &u}, nil
}

// UpdateArticle 更新
func (s *ArticleService) UpdateArticle(ctx context.Context, in *v1.UpdateArticleRequest) (*v1.UpdateArticleReply, error) {
	ArticleDO := &biz.ArticleDO{
		Title:          in.Title,
		PermaLink:      in.PermaLink,
		CanonicalLink:  in.CanonicalLink,
		Summary:        in.Summary,
		Thumbnail:      in.Thumbnail,
		Password:       in.Password,
		Atype:          int(in.Atype),
		CategoryId:     in.CategoryId,
		CommentAgentId: in.CommentAgentId,
		Published:      in.Published,
	}
	data, err := s.auc.Update(ctx, ArticleDO)
	if err != nil {
		return nil, err
	}
	u := articleResponse(data)
	return &v1.UpdateArticleReply{Content: &u}, nil
}

// DeleteArticle 根据ID删除
func (s *ArticleService) DeleteArticle(ctx context.Context, in *v1.DeleteArticleRequest) (*v1.DeleteArticleReply, error) {
	err := s.auc.Delete(ctx, in.Uid)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteArticleReply{}, nil
}

// DeleteArticles 根据ID批量删除
func (s *ArticleService) DeleteArticles(ctx context.Context, in *v1.DeleteArticlesRequest) (*v1.DeleteArticlesReply, error) {
	err := s.auc.DeleteList(ctx, in.Uids)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteArticlesReply{}, nil
}

// GetArticle 通过ID获取
func (s *ArticleService) GetArticle(ctx context.Context, in *v1.GetArticleRequest) (*v1.GetArticleReply, error) {
	data, err := s.auc.FindOneByID(ctx, in.Uid)
	if err != nil {
		return nil, err
	}
	u := articleResponse(data)
	// 获取详情
	//dataContent, errCont := s.ac.FindOneByID(ctx, in.Uid)
	//if errCont == nil {
	//	u.Content = dataContent.Content
	//	u.ContentHtml = dataContent.ContentHtml
	//}
	//user, uerr := s.uc.GetUser(ctx, data.AuthorId)
	//if uerr == nil {
	//	u.NickName = user.Nickname
	//}
	return &v1.GetArticleReply{Content: &u}, nil
}

func (s *ArticleService) ListArticle(ctx context.Context, in *v1.ListArticleRequest) (*v1.ListArticleReply, error) {
	opts := biz.ArticleDOListOption{}
	opts.ListOptions.Pages = int64(in.PageInfo.Pages)
	opts.ListOptions.Current = int64(in.PageInfo.Page)
	opts.ListOptions.PageSize = int64(in.PageInfo.Size)
	opts.ListOptions.Init()
	ArticleList, err := s.auc.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	pageInfo := &v1.PageInfo{
		Page:      uint64(ArticleList.Pages),
		Size:      uint64(ArticleList.PageSize),
		Total:     uint64(ArticleList.TotalCount),
		Pages:     uint64(ArticleList.Pages),
		FirstFlag: ArticleList.FirstFlag,
		LastFlag:  ArticleList.LastFlag,
	}
	datas := make([]*v1.ArticleInfoResponse, 0, len(ArticleList.Items))
	for _, data := range ArticleList.Items {
		tdata := &v1.ArticleInfoResponse{
			Uid:            data.ID,
			Title:          data.Title,
			PermaLink:      data.PermaLink,
			CanonicalLink:  data.CanonicalLink,
			Summary:        data.Summary,
			Thumbnail:      data.Thumbnail,
			Password:       data.Password,
			Atype:          int32(data.Atype),
			CategoryId:     data.CategoryId,
			CommentAgentId: data.CommentAgentId,
			Published:      data.Published,
			ViewCount:      data.ViewCount,
			LikeCount:      data.LikeCount,
			HateCount:      data.HateCount,
			PublishedAt:    timestamppb.New(data.PublishedAt),
		}

		// 获取详情
		//dataContent, errCont := s.ac.FindOneByID(ctx, data.ID)
		//if errCont == nil {
		//	tdata.Content = dataContent.Content
		//	tdata.ContentHtml = dataContent.ContentHtml
		//}
		datas = append(datas, tdata)
	}
	return &v1.ListArticleReply{PageInfo: pageInfo, Items: datas}, nil
}

func articleResponse(data *biz.ArticleDO) v1.ArticleInfoResponse {
	dataInfoRsp := v1.ArticleInfoResponse{
		Uid:            data.ID,
		Title:          data.Title,
		PermaLink:      data.PermaLink,
		CanonicalLink:  data.CanonicalLink,
		Summary:        data.Summary,
		Thumbnail:      data.Thumbnail,
		Password:       data.Password,
		Atype:          int32(data.Atype),
		CategoryId:     data.CategoryId,
		CommentAgentId: data.CommentAgentId,
		Published:      data.Published,
		ViewCount:      data.ViewCount,
		LikeCount:      data.LikeCount,
		HateCount:      data.HateCount,
		PublishedAt:    timestamppb.New(data.PublishedAt),
	}

	return dataInfoRsp
}
