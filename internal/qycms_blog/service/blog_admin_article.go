package service

import (
	"context"
	v1 "github.com/iwinder/qyblog/api/qycms_bff/admin/v1"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qyblog/internal/qycms_blog/biz"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

// CreateQyAdminArticle 创建
func (s *BlogAdminUserService) CreateQyAdminArticle(ctx context.Context, in *v1.CreateQyAdminArticleRequest) (*v1.CreateQyAdminArticleReply, error) {
	objDO := bizToArticleDO(in)
	usr := GetUserInfo(ctx)
	if usr == nil {
		return nil, ErrAuthFailed
	}
	objDO.CreatedBy = usr.ID
	objDO.UpdatedBy = usr.ID
	objDO.Nickname = usr.Nickname
	obj, err := s.au.Create(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.CreateQyAdminArticleReply{Id: obj.ID}, nil
}

// UpdateQyAdminArticle 更新
func (s *BlogAdminUserService) UpdateQyAdminArticle(ctx context.Context, in *v1.UpdateQyAdminArticleRequest) (*v1.UpdateQyAdminArticleReply, error) {
	objDO := bizToArticleDOByUpdate(in)
	usr := GetUserInfo(ctx)
	if usr == nil {
		return nil, ErrAuthFailed
	}
	objDO.UpdatedBy = usr.ID
	objDO.Nickname = usr.Nickname

	obj, err := s.au.Update(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateQyAdminArticleReply{Id: obj.ID}, nil
}

// DeleteQyAdminArticle 根据ID批量删除用户
func (s *BlogAdminUserService) DeleteQyAdminArticle(ctx context.Context, in *v1.DeleteQyAdminArticleRequest) (*v1.DeleteQyAdminArticleReply, error) {
	err := s.au.DeleteList(ctx, in.Ids)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteQyAdminArticleReply{}, nil
}

// GetQyAdminArticle 根据ID获取信息
func (s *BlogAdminUserService) GetQyAdminArticle(ctx context.Context, in *v1.GetQyAdminArticleRequest) (*v1.GetQyAdminArticleReply, error) {
	obj, err := s.au.FindOneByID(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	u := bizToArticleResponse(obj)
	return &v1.GetQyAdminArticleReply{Data: &u}, nil
}
func (s *BlogAdminUserService) InitQyAdminArticlePermaLink(ctx context.Context, in *v1.InitQyAdminArticlePermaLinkRequest) (*v1.InitQyAdminArticlePermaLinkReply, error) {
	obj := s.au.InitArticlePermaLink(ctx, in.Title)
	return &v1.InitQyAdminArticlePermaLinkReply{PermaLink: obj}, nil
}

// ListQyAdminArticle 获取列表
func (s *BlogAdminUserService) ListQyAdminArticle(ctx context.Context, in *v1.ListQyAdminArticleRequest) (*v1.ListQyAdminArticleReply, error) {
	GetHeardInfo(ctx)
	opts := biz.ArticleDOListOption{}
	opts.ListOptions.Pages = 0
	opts.ListOptions.Current = -1
	opts.ListOptions.PageSize = 20
	if in.Current > 0 {
		opts.ListOptions.Pages = in.Pages
		opts.ListOptions.Current = in.Current
		opts.ListOptions.PageSize = in.PageSize
	}
	opts.Title = in.SearchText
	opts.Atype = int(in.Atype)

	opts.StatusFlag = int(in.StatusFlag)
	opts.ListOptions.Init()

	objList, err := s.au.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	pageInfo := &v1.ArticlePageInfo{
		Current:   objList.Current,
		PageSize:  objList.PageSize,
		Total:     objList.TotalCount,
		Pages:     objList.Pages,
		FirstFlag: objList.FirstFlag,
		LastFlag:  objList.LastFlag,
	}
	objs := make([]*v1.ArticleInfoResponse, 0, len(objList.Items))
	for _, item := range objList.Items {
		titem := bizToArticleResponse(item)
		objs = append(objs, &titem)
	}
	return &v1.ListQyAdminArticleReply{PageInfo: pageInfo, Items: objs}, nil
}

func bizToArticleDO(in *v1.CreateQyAdminArticleRequest) *biz.ArticleDO {
	objDO := &biz.ArticleDO{
		Title:          in.Title,
		PermaLink:      in.PermaLink,
		CanonicalLink:  in.CanonicalLink,
		Summary:        in.Summary,
		Thumbnail:      in.Thumbnail,
		Password:       in.Password,
		Atype:          int(in.Atype),
		CategoryId:     in.CategoryId,
		CategoryName:   in.CategoryName,
		CommentAgentId: in.CommentAgentId,
		Published:      in.Published,
		TagStrings:     in.TagStrings,
		Content:        in.Content,
		ContentHtml:    in.ContentHtml,
		CommentFlag:    in.CommentFlag,
	}

	objDO.StatusFlag = 1
	if in.Published {
		objDO.StatusFlag = 2
		objDO.PublishedAt = time.Now()
	}
	flen := len(in.Resources)
	if flen > 0 {
		files := make([]*biz.ArticleResourceDO, 0, flen)
		for _, file := range in.Resources {
			files = append(files, &biz.ArticleResourceDO{
				ObjectMeta: metaV1.ObjectMeta{
					ID: file.Id,
				},
				Name:     file.Name,
				Url:      file.Url,
				Password: file.Password,
			})
		}
		objDO.Resource = files
	}
	return objDO
}
func bizToArticleDOByUpdate(in *v1.UpdateQyAdminArticleRequest) *biz.ArticleDO {
	objDO := &biz.ArticleDO{
		Title:          in.Title,
		PermaLink:      in.PermaLink,
		CanonicalLink:  in.CanonicalLink,
		Summary:        in.Summary,
		Thumbnail:      in.Thumbnail,
		Password:       in.Password,
		Atype:          int(in.Atype),
		CategoryId:     in.CategoryId,
		CategoryName:   in.CategoryName,
		CommentAgentId: in.CommentAgentId,
		Published:      in.Published,
		Nickname:       in.NickName,
		TagStrings:     in.TagStrings,
		Content:        in.Content,
		ContentHtml:    in.ContentHtml,
		CommentFlag:    in.CommentFlag,
	}
	objDO.ID = in.Id
	objDO.StatusFlag = 1
	if in.Published {
		objDO.StatusFlag = 2
		objDO.PublishedAt = time.Now()
	}
	flen := len(in.Resources)
	if flen > 0 {
		files := make([]*biz.ArticleResourceDO, 0, flen)
		for _, file := range in.Resources {
			files = append(files, &biz.ArticleResourceDO{
				ObjectMeta: metaV1.ObjectMeta{
					ID: file.Id,
				},
				ArticleID: objDO.ID,
				Name:      file.Name,
				Url:       file.Url,
				Password:  file.Password,
			})
		}
		objDO.Resource = files
	}
	return objDO
}
func bizToArticleResponse(obj *biz.ArticleDO) v1.ArticleInfoResponse {
	objInfoRsp := v1.ArticleInfoResponse{
		Id:             obj.ID,
		Title:          obj.Title,
		PermaLink:      obj.PermaLink,
		Summary:        obj.Summary,
		Thumbnail:      obj.Thumbnail,
		Password:       obj.Password,
		StatusFlag:     int32(obj.StatusFlag),
		Atype:          int32(obj.Atype),
		AuthorId:       obj.CreatedBy,
		CategoryId:     obj.CategoryId,
		CategoryName:   obj.CategoryName,
		CommentAgentId: obj.CommentAgentId,
		Published:      obj.Published,
		PublishedAt:    timestamppb.New(obj.PublishedAt),
		NickName:       obj.Nickname,
		TagStrings:     obj.TagStrings,
		CanonicalLink:  obj.CanonicalLink,
		Content:        obj.Content,
		ContentHtml:    obj.ContentHtml,
		CommentFlag:    obj.CommentFlag,
	}

	flen := len(obj.Resource)
	if flen > 0 {
		files := make([]*v1.ArticleResourcesInfoResponse, 0, flen)
		for _, file := range obj.Resource {
			files = append(files, &v1.ArticleResourcesInfoResponse{
				Id:       file.ID,
				Name:     file.Name,
				Url:      file.Url,
				Password: file.Password,
			})
		}
		objInfoRsp.Resources = files
	}
	return objInfoRsp
}
