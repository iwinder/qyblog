package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/app/qycms_blog/internal/data/po"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"gorm.io/gorm"
)

type ArticleContentDO struct {
	metaV1.ObjectMeta
	Status      int
	Atype       int
	Content     string
	ContentHtml string
}

// ArticleContentRepo is a Greater repo.
type ArticleContentRepo interface {
	Save(context.Context, *ArticleContentDO) (*po.ArticleContentPO, error)
	Update(context.Context, *ArticleContentDO) (*po.ArticleContentPO, error)
	Delete(context.Context, uint64) error
	DeleteList(c context.Context, uids []uint64) error
	FindByID(context.Context, uint64) (*po.ArticleContentPO, error)
	ListAll(context.Context) ([]*po.ArticleContentPO, error)
}

// ArticleContentUsecase is a ArticleDO usecase.
type ArticleContentUsecase struct {
	repo ArticleContentRepo
	log  *log.Helper
}

// NewArticleContentUsecase new a ArticleDO usecase.
func NewArticleContentUsecase(repo ArticleContentRepo, logger log.Logger) *ArticleContentUsecase {
	return &ArticleContentUsecase{repo: repo, log: log.NewHelper(logger)}
}

// Create creates a ArticleDO, and returns the new ArticleDO.
func (uc *ArticleContentUsecase) Create(ctx context.Context, g *ArticleContentDO) (*ArticleContentDO, error) {
	uc.log.WithContext(ctx).Infof("CreateArticle: %v", g.ID)
	dataPO, err := uc.repo.Save(ctx, g)
	if err != nil {
		return nil, err
	}
	data := &ArticleContentDO{}
	data.ID = dataPO.ID
	return data, nil
}

// Update 更新
func (uc *ArticleContentUsecase) Update(ctx context.Context, g *ArticleContentDO) (*ArticleContentDO, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", g.ID)
	dataPO, err := uc.repo.Update(ctx, g)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	data := &ArticleContentDO{}
	data.ID = dataPO.ID
	return data, nil
}

// Delete 根据ID删除
func (uc *ArticleContentUsecase) Delete(ctx context.Context, id uint64) error {
	uc.log.WithContext(ctx).Infof("Delete: %v", id)
	return uc.repo.Delete(ctx, id)
}

// DeleteList 根据ID批量删除
func (uc *ArticleContentUsecase) DeleteList(ctx context.Context, ids []uint64) error {
	uc.log.WithContext(ctx).Infof("DeleteList: %v", ids)
	return uc.repo.DeleteList(ctx, ids)
}

// FindOneByID 根据ID查询信息
func (uc *ArticleContentUsecase) FindOneByID(ctx context.Context, id uint64) (*ArticleContentDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByID: %v", id)
	g, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	data := &ArticleContentDO{
		ObjectMeta:  g.ObjectMeta,
		Status:      g.Status,
		Atype:       g.Atype,
		Content:     g.Content,
		ContentHtml: g.ContentHtml,
	}
	return data, nil
}

// ListAll 批量查询
func (uc *ArticleContentUsecase) ListAll(ctx context.Context, opts ArticleDOListOption) (*ArticleDOList, error) {

	return nil, nil
}
