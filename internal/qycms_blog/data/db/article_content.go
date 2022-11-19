package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
)

type articleContentRepo struct {
	data *Data
	log  *log.Helper
}

// NewArticleContentRepo .
func NewArticleContentRepo(data *Data, logger log.Logger) biz.ArticleContentRepo {
	return &articleContentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Save 新增
func (r *articleContentRepo) Save(ctx context.Context, g *biz.ArticleContentDO) (*po.ArticleContentPO, error) {
	newData := &po.ArticleContentPO{
		ID:          g.ID,
		StatusFlag:  g.StatusFlag,
		Atype:       g.Atype,
		Content:     g.Content,
		ContentHtml: g.ContentHtml,
	}
	err := r.data.Db.Create(newData).Error
	if err != nil {
		return nil, err
	}
	return newData, nil
}

// Update 根据ID更新
func (r *articleContentRepo) Update(ctx context.Context, g *biz.ArticleContentDO) (*po.ArticleContentPO, error) {
	newData := &po.ArticleContentPO{
		ID:          g.ID,
		StatusFlag:  g.StatusFlag,
		Atype:       g.Atype,
		Content:     g.Content,
		ContentHtml: g.ContentHtml,
	}
	tData := &po.ArticleContentPO{}
	tData.ID = g.ID
	err := r.data.Db.Model(&tData).Updates(&newData).Error
	if err != nil {
		return nil, err
	}
	return newData, nil
}

// Delete 根据ID删除
func (r *articleContentRepo) Delete(ctx context.Context, id uint64) error {
	tData := &po.ArticleContentPO{}
	tData.ID = id
	err := r.data.Db.Delete(&tData).Error
	return err

}

// DeleteList 批量删除
func (r *articleContentRepo) DeleteList(ctx context.Context, ids []uint64) error {
	tData := &po.ArticleContentPO{}
	err := r.data.Db.Delete(&tData, ids).Error
	return err
}

// FindByID 根据ID查询
func (r *articleContentRepo) FindByID(ctx context.Context, id uint64) (*po.ArticleContentPO, error) {
	tData := &po.ArticleContentPO{}
	err := r.data.Db.Where("id = ?", id).First(&tData).Error
	if err != nil {
		return nil, err
	}

	return tData, nil
}

// ListAll 批量查询
func (r *articleContentRepo) ListAll(context.Context) ([]*po.ArticleContentPO, error) {
	return nil, nil
}
