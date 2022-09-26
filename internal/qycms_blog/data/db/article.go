package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/gormutil"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

// NewArticleRepo .
func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Save 新增
func (r *articleRepo) Save(ctx context.Context, g *biz.ArticleDO) (*po.ArticlePO, error) {
	newData := &po.ArticlePO{
		ObjectMeta:     g.ObjectMeta,
		Title:          g.Title,
		PermaLink:      g.PermaLink,
		CanonicalLink:  g.CanonicalLink,
		Summary:        g.Summary,
		Thumbnail:      g.Thumbnail,
		Password:       g.Password,
		Status:         g.Status,
		Atype:          g.Atype,
		AuthorId:       g.AuthorId,
		CategoryId:     g.CategoryId,
		CommentAgentId: g.CommentAgentId,
		Published:      g.Published,
		ViewCount:      g.ViewCount,
		LikeCount:      g.LikeCount,
		HateCount:      g.HateCount,
		PublishedAt:    g.PublishedAt,
	}
	err := r.data.Db.Create(newData).Error
	if err != nil {
		return nil, err
	}
	return newData, nil
}

// Update 根据ID更新
func (r *articleRepo) Update(ctx context.Context, g *biz.ArticleDO) (*po.ArticlePO, error) {
	newData := &po.ArticlePO{
		Title:          g.Title,
		PermaLink:      g.PermaLink,
		CanonicalLink:  g.CanonicalLink,
		Summary:        g.Summary,
		Thumbnail:      g.Thumbnail,
		Password:       g.Password,
		Status:         g.Status,
		Atype:          g.Atype,
		AuthorId:       g.AuthorId,
		CategoryId:     g.CategoryId,
		CommentAgentId: g.CommentAgentId,
		Published:      g.Published,
		ViewCount:      g.ViewCount,
		LikeCount:      g.LikeCount,
		HateCount:      g.HateCount,
		PublishedAt:    g.PublishedAt,
	}
	tData := &po.ArticlePO{}
	tData.ID = g.ID
	err := r.data.Db.Model(&tData).Updates(&newData).Error
	if err != nil {
		return nil, err
	}
	return newData, nil
}

// Delete 根据ID删除
func (r *articleRepo) Delete(ctx context.Context, id uint64) error {
	tData := &po.ArticlePO{}
	tData.ID = id
	err := r.data.Db.Delete(&tData).Error
	return err

}

// DeleteList 批量删除
func (r *articleRepo) DeleteList(ctx context.Context, ids []uint64) error {
	tData := &po.ArticlePO{}
	err := r.data.Db.Delete(&tData, ids).Error
	return err
}

// FindByID 根据ID查询
func (r *articleRepo) FindByID(ctx context.Context, id uint64) (*po.ArticlePO, error) {
	tData := &po.ArticlePO{}
	err := r.data.Db.Where("id = ?", id).First(&tData).Error
	if err != nil {
		return nil, err
	}

	return tData, nil
}

// ListAll 批量查询
func (r *articleRepo) ListAll(ctx context.Context, opts biz.ArticleDOListOption) (*po.ArticlePOList, error) {
	ret := &po.ArticlePOList{}

	where := &po.ArticlePO{}
	var err error

	if opts.PageFlag {
		ol := gormutil.Unpointer(opts.Offset, opts.Limit)
		d := r.data.Db.Model(where).Where(where).
			Offset(ol.Offset).
			Limit(ol.Limit).
			Order("id desc").
			Find(&ret.Items).
			Offset(-1).
			Limit(-1).
			Count(&ret.TotalCount)
		err = d.Error
	} else {
		d := r.data.Db.Model(where).Where(where).
			Find(&ret.Items).
			Count(&ret.TotalCount)
		err = d.Error
	}
	opts.TotalCount = ret.TotalCount
	opts.IsLast()
	ret.FirstFlag = opts.FirstFlag
	ret.Current = opts.Current
	ret.PageSize = opts.PageSize
	ret.LastFlag = opts.LastFlag
	return ret, err
}
