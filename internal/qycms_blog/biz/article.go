package biz

import (
	"context"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
	"gorm.io/gorm"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1.ErrorReason_DATA_NOT_FOUND.String(), "user not found")
)

// ArticleDO is a ArticleDO model.
type ArticleDO struct {
	metaV1.ObjectMeta
	Title          string
	PermaLink      string
	CanonicalLink  string
	Summary        string
	Thumbnail      string
	Password       string
	Status         int
	Atype          int
	CategoryId     uint64
	CommentAgentId uint64
	Published      bool
	ViewCount      int32
	LikeCount      int32
	HateCount      int32
	PublishedAt    time.Time
	Tags           []*TagsDO
}

type ArticleDOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*ArticleDO `json:"items"`
}

type ArticleDOListOption struct {
	metaV1.ListOptions `json:"page"`
	ArticleDO          `json:"item"`
}

// ArticleRepo is a Greater repo.
type ArticleRepo interface {
	Save(context.Context, *ArticleDO) (*po.ArticlePO, error)
	Update(context.Context, *ArticleDO) (*po.ArticlePO, error)
	Delete(context.Context, uint64) error
	DeleteList(c context.Context, uids []uint64) error
	FindByID(context.Context, uint64) (*po.ArticlePO, error)
	ListAll(context.Context, ArticleDOListOption) (*po.ArticlePOList, error)
}

// ArticleUsecase   is a ArticleDO usecase.
type ArticleUsecase struct {
	repo ArticleRepo
	log  *log.Helper
}

// NewArticleContentUsecase new a ArticleDO usecase.
func NewArticleUsecase(repo ArticleRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateArticle creates a ArticleDO, and returns the new ArticleDO.
func (uc *ArticleUsecase) CreateArticle(ctx context.Context, g *ArticleDO) (*ArticleDO, error) {
	uc.log.WithContext(ctx).Infof("CreateArticle: %v", g.Title)
	dataPO, err := uc.repo.Save(ctx, g)
	if err != nil {
		return nil, err
	}
	data := &ArticleDO{Title: dataPO.Title}
	data.ID = dataPO.ID
	return data, nil
}

// Update 更新
func (uc *ArticleUsecase) Update(ctx context.Context, g *ArticleDO) (*ArticleDO, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", g.Title)
	dataPO, err := uc.repo.Update(ctx, g)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	data := &ArticleDO{Title: dataPO.Title}
	data.ID = dataPO.ID
	return data, nil
}

// Delete 根据ID删除
func (uc *ArticleUsecase) Delete(ctx context.Context, id uint64) error {
	uc.log.WithContext(ctx).Infof("Delete: %v", id)
	return uc.repo.Delete(ctx, id)
}

// DeleteList 根据ID批量删除
func (uc *ArticleUsecase) DeleteList(ctx context.Context, ids []uint64) error {
	uc.log.WithContext(ctx).Infof("DeleteList: %v", ids)
	return uc.repo.DeleteList(ctx, ids)
}

// FindOneByID 根据ID查询信息
func (uc *ArticleUsecase) FindOneByID(ctx context.Context, id uint64) (*ArticleDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByID: %v", id)
	g, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	data := &ArticleDO{
		ObjectMeta:     g.ObjectMeta,
		Title:          g.Title,
		PermaLink:      g.PermaLink,
		CanonicalLink:  g.CanonicalLink,
		Summary:        g.Summary,
		Thumbnail:      g.Thumbnail,
		Password:       g.Password,
		Status:         g.Status,
		Atype:          g.Atype,
		CategoryId:     g.CategoryId,
		CommentAgentId: g.CommentAgentId,
		Published:      g.Published,
		ViewCount:      g.ViewCount,
		LikeCount:      g.LikeCount,
		HateCount:      g.HateCount,
		PublishedAt:    g.PublishedAt,
	}
	return data, nil
}

// ListAll 批量查询
func (uc *ArticleUsecase) ListAll(ctx context.Context, opts ArticleDOListOption) (*ArticleDOList, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	dataPOs, err := uc.repo.ListAll(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	infos := make([]*ArticleDO, 0, len(dataPOs.Items))
	for _, data := range dataPOs.Items {
		infos = append(infos, &ArticleDO{
			ObjectMeta:     data.ObjectMeta,
			Title:          data.Title,
			PermaLink:      data.PermaLink,
			CanonicalLink:  data.CanonicalLink,
			Summary:        data.Summary,
			Thumbnail:      data.Thumbnail,
			Password:       data.Password,
			Status:         data.Status,
			Atype:          data.Atype,
			CategoryId:     data.CategoryId,
			CommentAgentId: data.CommentAgentId,
			Published:      data.Published,
			ViewCount:      data.ViewCount,
			LikeCount:      data.LikeCount,
			HateCount:      data.HateCount,
			PublishedAt:    data.PublishedAt,
		})
	}
	return &ArticleDOList{ListMeta: dataPOs.ListMeta, Items: infos}, nil
}
