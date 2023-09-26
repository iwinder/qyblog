package db

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
	"gorm.io/gorm/clause"
)

var (
	// ErrApiNotFound is api not found.
	ErrResourceNotFound = errors.NotFound("101404A", "测试异常效果")
)

type articleResourceRepo struct {
	data *Data
	log  *log.Helper
}

func NewArticleResourceRepo(data *Data, logger log.Logger) biz.ArticleResourceRepo {
	return &articleResourceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r articleResourceRepo) CreateInBatches(ctx context.Context, articleID uint64, articleResources []*biz.ArticleResourceDO) error {
	articleResourcePos := make([]*po.ArticleResourcePO, 0, len(articleResources))
	for i, obj := range articleResources {
		articleResourcePos = append(articleResourcePos, &po.ArticleResourcePO{
			ObjectMeta: metaV1.ObjectMeta{
				InstanceID: fmt.Sprintf("%s-%d", "def-", i),
			},
			ArticleID: articleID,
			Name:      obj.Name,
			Url:       obj.Url,
			Password:  obj.Password,
		})
	}
	d := r.data.Db.Create(&articleResourcePos)
	return d.Error
}

func (r articleResourceRepo) UpdateInBatches(ctx context.Context, articleResource []*biz.ArticleResourceDO) error {
	articleResourcePos := make([]*po.ArticleResourcePO, 0, len(articleResource))
	for _, obj := range articleResource {
		articleResourcePos = append(articleResourcePos, &po.ArticleResourcePO{
			ObjectMeta: metaV1.ObjectMeta{
				ID: obj.ID,
			},
			ArticleID: obj.ArticleID,
			Name:      obj.Name,
			Url:       obj.Url,
			Password:  obj.Password,
		})
	}
	//d := r.data.Db.Model(&articleResourcePos).Updates(&articleResourcePos)
	d := r.data.Db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "url", "password"}),
	}).Create(&articleResourcePos)
	return d.Error
}

func (r articleResourceRepo) DeleteByArticleID(ctx context.Context, articleId uint64) error {
	err := r.data.Db.Where("article_id = ?", articleId).Delete(&po.ArticleResourcePO{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r articleResourceRepo) DeleteByIDs(ctx context.Context, ids []uint64) error {
	err := r.data.Db.Delete(&po.ArticleResourcePO{}, ids).Error
	if err != nil {
		return err
	}
	return nil
}

func (r articleResourceRepo) FindAllByArticleID(ctx context.Context, articleId uint64) ([]*biz.ArticleResourceDO, error) {
	filesPO := make([]*po.ArticleResourcePO, 0, 0)
	db := r.data.Db
	e := db.Where(" article_id = ?", articleId).Find(&filesPO)
	if e.Error != nil {
		return nil, e.Error
	}
	infos := make([]*biz.ArticleResourceDO, 0, len(filesPO))
	for _, obj := range filesPO {
		infos = append(infos, &biz.ArticleResourceDO{
			ObjectMeta: metaV1.ObjectMeta{
				ID: obj.ID,
			},
			ArticleID: articleId,
			Name:      obj.Name,
			Url:       obj.Url,
			Password:  obj.Password,
		})
	}
	return infos, nil
}

func (r articleResourceRepo) FindAllByArticlePermaLink(ctx context.Context, permaLink string) ([]*biz.ArticleResourceDO, error) {
	filesPO := make([]*po.ArticleResourcePO, 0, 0)
	db := r.data.Db
	e := db.Where(" article_id = ?", db.Table(" qy_blog_article a ").Where("perma_link = ? ", permaLink)).Find(&filesPO)
	if e.Error != nil {
		return nil, e.Error
	}
	infos := make([]*biz.ArticleResourceDO, len(filesPO))
	for _, obj := range filesPO {
		infos = append(infos, &biz.ArticleResourceDO{
			Name:     obj.Name,
			Url:      obj.Url,
			Password: obj.Password,
		})
	}
	return infos, nil
}

func (r *articleResourceRepo) CountByArticleID(ctx context.Context, articleId uint64) (int64, error) {
	var obj int64
	err := r.data.Db.Model(&po.ArticlePO{}).Where("article_id = ?", articleId).Count(&obj).Error
	if err != nil {
		return 0, err
	}
	return obj, nil
}
