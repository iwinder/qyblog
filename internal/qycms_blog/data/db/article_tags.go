package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
)

type articleTagsRepo struct {
	data *Data
	log  *log.Helper
}

func NewArticleTagsRepo(data *Data, logger log.Logger) biz.ArticleTagsRepo {
	return &articleTagsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *articleTagsRepo) CreateInBatches(ctx context.Context, articleTagss []*biz.ArticleTagsDO) error {
	articleTagsPos := make([]*po.ArticleTagsPO, 0, len(articleTagss))
	for _, obj := range articleTagss {
		articleTagsPos = append(articleTagsPos, &po.ArticleTagsPO{
			ArticleID: obj.ArticleID,
			TagID:     obj.TagID,
		})
	}
	d := r.data.Db.Create(&articleTagsPos)
	return d.Error
}

func (r *articleTagsRepo) UpdateInBatches(ctx context.Context, articleTagss []*biz.ArticleTagsDO) error {
	err := r.DeleteByArticleID(ctx, articleTagss[0].ArticleID)
	if err != nil {
		return err
	}
	articleTagsPos := make([]*po.ArticleTagsPO, 0, len(articleTagss))
	for _, obj := range articleTagss {
		articleTagsPos = append(articleTagsPos, &po.ArticleTagsPO{
			ArticleID: obj.ArticleID,
			TagID:     obj.TagID,
		})
	}
	d := r.data.Db.Create(&articleTagsPos)
	return d.Error
}

func (r *articleTagsRepo) DeleteByArticleID(ctx context.Context, articleId uint64) error {
	err := r.data.Db.Where("article_id = ?", articleId).Delete(&po.ArticleTagsPO{}).Error
	if err != nil {
		return err
	}
	return nil
}
