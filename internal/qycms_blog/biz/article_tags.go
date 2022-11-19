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
	repo ArticleTagsRepo
	tu   *TagsUsecase
	log  *log.Helper
}

// NewArticleTagsUsecase new a ArticleDO usecase.
func NewArticleTagsUsecase(repo ArticleTagsRepo, tu *TagsUsecase, logger log.Logger) *ArticleTagsUsecase {
	return &ArticleTagsUsecase{repo: repo, tu: tu, log: log.NewHelper(logger)}
}

func (uc *ArticleTagsUsecase) SaveTagsForArticle(ctx context.Context, article *ArticleDO) error {
	oldTsgs, err := uc.tu.FindAllByArticleID(ctx, article.ID)
	if err != nil {
		uc.log.WithContext(ctx).Error(err)
	}
	oldTsgMap := make(map[string]*TagsDO, len(oldTsgs))
	for _, tag := range oldTsgs {
		oldTsgMap[tag.Name] = tag
	}

	if article.TagStrings != nil && len(article.TagStrings) > 0 {
		articleTagss := make([]*ArticleTagsDO, 0, 0)
		tagss := make([]*TagsDO, 0, 0)
		for _, obj := range article.TagStrings {
			tag, ok := oldTsgMap[obj]
			if ok {
				tagss = append(tagss, tag)
				articleTagss = append(articleTagss, &ArticleTagsDO{
					ArticleID: article.ID,
					TagID:     tag.ID,
				})
			} else {
				newTag, _ := uc.tu.FindOneByName(ctx, obj)
				if newTag == nil {
					newTag = &TagsDO{Name: obj}
					tagDo, terr := uc.tu.Create(ctx, newTag)
					if terr != nil {
						uc.log.WithContext(ctx).Error(terr)
					} else {
						articleTagss = append(articleTagss, &ArticleTagsDO{
							ArticleID: article.ID,
							TagID:     tagDo.ID,
						})
					}
				}
				articleTagss = append(articleTagss, &ArticleTagsDO{
					ArticleID: article.ID,
					TagID:     newTag.ID,
				})

			}
		}
		aerr := uc.repo.DeleteByArticleID(ctx, article.ID)
		if aerr != nil {
			uc.log.WithContext(ctx).Error(aerr)
		}
		// 关联关系
		if len(articleTagss) > 0 {
			// 关联
			cierr := uc.repo.CreateInBatches(ctx, articleTagss)
			if cierr != nil {
				return cierr
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
func (uc *ArticleTagsUsecase) FindAllByArticleID(ctx context.Context, aid uint64) ([]*TagsDO, error) {
	oldTsgs, err := uc.tu.FindAllByArticleID(ctx, aid)
	if err != nil {
		uc.log.WithContext(ctx).Error(err)
	}
	return oldTsgs, nil
}
