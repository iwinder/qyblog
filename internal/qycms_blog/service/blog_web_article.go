package service

import (
	"context"
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/web/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (b *BlogWebApiService) GetQyWebArticle(ctx context.Context, in *v1.GetQyWebArticleRequest) (*v1.GetQyWebArticleReply, error) {
	obj, err := b.au.FindOneByLink(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	// 增加浏览记录
	u := bizToWebArticleResponse(obj)
	return &v1.GetQyWebArticleReply{Data: u}, nil
}

func (b *BlogWebApiService) ListQyWebArticle(ctx context.Context, in *v1.ListQyWebArticleRequest) (*v1.ListQyWebArticleReply, error) {
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
	opts.CategoryName = in.CategoryName
	opts.TagName = in.TagName
	opts.StatusFlag = 2
	opts.Atype = 1 // TODO: 暂且指定死，后期有其它需要前端展示列表的再放开
	opts.ListOptions.Init()
	objList, err := b.au.ListAllForWeb(ctx, opts)
	if err != nil {
		return nil, err
	}
	pageInfo := &v1.WebArticlePageInfo{
		Current:   objList.Current,
		PageSize:  objList.PageSize,
		Total:     objList.TotalCount,
		Pages:     objList.Pages,
		FirstFlag: objList.FirstFlag,
		LastFlag:  objList.LastFlag,
	}
	objs := make([]*v1.WebArticleInfoResponse, 0, len(objList.Items))
	for _, item := range objList.Items {
		titem := bizToWebArticleResponse(item)
		objs = append(objs, titem)
	}
	return &v1.ListQyWebArticleReply{PageInfo: pageInfo, Items: objs}, nil
}
func (b *BlogWebApiService) GetQyWebCategory(ctx context.Context, in *v1.GetQyWebCategoryRequest) (*v1.GetQyWebCategoryReply, error) {
	data, err := b.cu.FindByIdentifier(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	dataDto := &v1.WebCategoryResponse{
		Name:       data.Name,
		Identifier: data.Identifier,
	}
	return &v1.GetQyWebCategoryReply{Data: dataDto}, nil
}

func (b *BlogWebApiService) GetQyWebTag(ctx context.Context, in *v1.GetQyWebTagRequest) (*v1.GetQyWebTagReply, error) {
	data, err := b.tu.FindOneByIdentifier(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	dataDto := &v1.WebTagsResponse{
		Name:       data.Name,
		Identifier: data.Identifier,
	}
	return &v1.GetQyWebTagReply{Data: dataDto}, nil
}
func bizToWebArticleResponse(obj *biz.ArticleDO) *v1.WebArticleInfoResponse {
	objInfoRsp := &v1.WebArticleInfoResponse{
		Title:          obj.Title,
		PermaLink:      obj.PermaLink,
		Summary:        obj.Summary,
		Thumbnail:      obj.Thumbnail,
		StatusFlag:     int32(obj.StatusFlag),
		Atype:          int32(obj.Atype),
		AuthorId:       obj.CreatedBy,
		CommentAgentId: obj.CommentAgentId,
		Published:      obj.Published,
		PublishedAt:    timestamppb.New(obj.PublishedAt),
		NickName:       obj.Nickname,
		CanonicalLink:  obj.CanonicalLink,
		ContentHtml:    obj.ContentHtml,
		CommentFlag:    obj.CommentFlag,
		CommentCount:   obj.CommentCount,
	}
	if obj.Category != nil {
		objInfoRsp.Category = &v1.WebCategoryResponse{
			Name:       obj.Category.Name,
			Identifier: obj.Category.Identifier,
		}
	} else {
		objInfoRsp.Category = &v1.WebCategoryResponse{}
	}

	if len(obj.Tags) > 0 {
		tags := make([]*v1.WebTagsResponse, 0, len(obj.Tags))
		for _, tag := range obj.Tags {
			tags = append(tags, &v1.WebTagsResponse{
				Name:       tag.Name,
				Identifier: tag.Identifier,
			})
		}
		objInfoRsp.Tags = tags
	}
	return objInfoRsp
}
