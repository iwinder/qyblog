package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ArticleTagsDO struct {
	ArticleID uint64
	TagID     uint64
}

type ArticleTagsRepo interface {
	CreateInBatches(ctx context.Context, articleTagss []*ArticleTagsDO) error
	UpdateInBatches(ctx context.Context, articleTagss []*ArticleTagsDO) error
	DeleteByArticleID(ctx context.Context, articleId uint64) error
}

type ArticleTagsUsecase struct {
	repo      ArticleTagsRepo
	cabinRepo CasbinRuleRepo
	log       *log.Helper
}

// NewArticleTagsUsecase new a ArticleDO usecase.
func NewArticleTagsUsecase(repo ArticleTagsRepo, cabinRepo CasbinRuleRepo, logger log.Logger) *ArticleTagsUsecase {
	return &ArticleTagsUsecase{repo: repo, cabinRepo: cabinRepo, log: log.NewHelper(logger)}
}

func (uc *ArticleTagsUsecase) SaveRoleForUser(ctx context.Context, article *ArticleDO) error {

	if article.Tags != nil && len(article.Tags) > 0 {
		rlen := len(article.Tags)
		articleTagss := make([]*ArticleTagsDO, 0, rlen)
		for _, obj := range article.Tags {
			articleTagss = append(articleTagss, &ArticleTagsDO{
				ArticleID: article.ID,
				TagID:     obj.ID,
			})
		}
		// 关联关系
		if len(articleTagss) > 0 {
			// 关联
			err := uc.repo.CreateInBatches(ctx, articleTagss)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func (uc *ArticleTagsUsecase) UpdateRoleForUser(ctx context.Context, article *ArticleDO) error {

	if article.Tags != nil && len(article.Tags) > 0 {
		rlen := len(article.Tags)
		articleTagss := make([]*ArticleTagsDO, 0, rlen)
		for _, obj := range article.Tags {
			articleTagss = append(articleTagss, &ArticleTagsDO{
				ArticleID: article.ID,
				TagID:     obj.ID,
			})
		}
		// 关联关系
		if len(articleTagss) > 0 {
			// 关联
			err := uc.repo.UpdateInBatches(ctx, articleTagss)
			if err != nil {
				return err
			}

		}

	} else { // 删除用户-角色关系
		// 关联
		err := uc.repo.DeleteByArticleID(ctx, article.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
