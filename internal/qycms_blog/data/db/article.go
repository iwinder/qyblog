package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/gormutil"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
	"time"
)

var articleCacheKey = func(link string) string {
	return "article_cache_key_" + link
}

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
func (r *articleRepo) Save(ctx context.Context, g *biz.ArticleDO) (*biz.ArticleDO, error) {
	newData := &po.ArticlePO{
		ObjectMeta:     g.ObjectMeta,
		Title:          g.Title,
		PermaLink:      g.PermaLink,
		CanonicalLink:  g.CanonicalLink,
		Summary:        g.Summary,
		Thumbnail:      g.Thumbnail,
		Password:       g.Password,
		Atype:          g.Atype,
		CategoryId:     g.CategoryId,
		CategoryName:   g.CategoryName,
		CommentAgentId: g.CommentAgentId,
		Nickname:       g.Nickname,
		Published: sql.NullBool{
			Bool:  g.Published,
			Valid: true,
		},
		ViewCount:   g.ViewCount,
		LikeCount:   g.LikeCount,
		HateCount:   g.HateCount,
		PublishedAt: g.PublishedAt,
		CommentFlag: sql.NullBool{
			Bool:  g.CommentFlag,
			Valid: true,
		},
	}
	newData.StatusFlag = g.StatusFlag
	err := r.data.Db.Create(newData).Error
	if err != nil {
		return nil, err
	}
	data := &biz.ArticleDO{Title: newData.Title}
	data.ID = newData.ID
	return data, nil
}

// Update 根据ID更新
func (r *articleRepo) Update(ctx context.Context, g *biz.ArticleDO) (*biz.ArticleDO, error) {
	newData := &po.ArticlePO{
		Title:          g.Title,
		PermaLink:      g.PermaLink,
		CanonicalLink:  g.CanonicalLink,
		Summary:        g.Summary,
		Thumbnail:      g.Thumbnail,
		Password:       g.Password,
		Atype:          g.Atype,
		CategoryId:     g.CategoryId,
		CategoryName:   g.CategoryName,
		CommentAgentId: g.CommentAgentId,
		Nickname:       g.Nickname,
		Published: sql.NullBool{
			Bool:  g.Published,
			Valid: true,
		},
		ViewCount:   g.ViewCount,
		LikeCount:   g.LikeCount,
		HateCount:   g.HateCount,
		PublishedAt: g.PublishedAt,
		CommentFlag: sql.NullBool{
			Bool:  g.CommentFlag,
			Valid: true,
		},
	}
	newData.StatusFlag = g.StatusFlag
	tData := &po.ArticlePO{}
	tData.ID = g.ID
	err := r.data.Db.Model(&tData).Updates(&newData).Error
	if err != nil {
		return nil, err
	}
	data := &biz.ArticleDO{Title: newData.Title}
	data.ID = newData.ID
	r.SetArticleCache(ctx, nil, g.PermaLink)
	return data, nil
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

func (r *articleRepo) CountByPermaLink(ctx context.Context, str string) (int64, error) {
	var obj int64
	err := r.data.Db.Model(&po.ArticlePO{}).Where("perma_link like ?", str+"%").Count(&obj).Error
	if err != nil {
		return 0, err
	}
	return obj, nil
}

// FindByID 根据ID查询
func (r *articleRepo) FindByID(ctx context.Context, id uint64) (*biz.ArticleDO, error) {
	g := &po.ArticlePO{}
	err := r.data.Db.Where("id = ?", id).First(&g).Error
	if err != nil {
		return nil, err
	}
	data := bizToArticleDO(g)
	return data, nil
}
func (r *articleRepo) FindByLink(ctx context.Context, link string) (*biz.ArticleDO, error) {
	g := &po.ArticlePO{}
	err := r.data.Db.Where("perma_link = ?", link).First(&g).Error
	if err != nil {
		return nil, err
	}
	data := bizToArticleDO(g)
	return data, nil
}

// FindByAgentID 根据ID查询
func (r *articleRepo) FindByAgentID(ctx context.Context, id uint64) (*biz.ArticleDO, error) {
	g := &po.ArticlePO{}
	err := r.data.Db.Where("comment_agent_id = ?", id).First(&g).Error
	if err != nil {
		return nil, err
	}
	data := bizToArticleDO(g)
	return data, nil
}

// ListAll 批量查询
func (r *articleRepo) ListAll(ctx context.Context, opts biz.ArticleDOListOption) (*biz.ArticleDOList, error) {
	ret := &po.ArticlePOList{}

	where := &po.ArticlePO{}
	var err error
	db := r.data.Db
	query := db.Model(where)

	if len(opts.Title) > 0 {
		query.Where(" title like ? ", "%"+opts.Title+"%")
	}
	if opts.Atype > 0 {
		query.Scopes(withFilterKeyEquarlsValue("atype", opts.Atype))
	}
	if opts.StatusFlag > 0 {
		query.Scopes(withFilterKeyEquarlsValue("status_flag", opts.StatusFlag))
	}
	if len(opts.TagName) > 0 {
		query.Where("ID in (?)", db.Table("qy_blog_article_tags").Select("article_id").Where("tag_id = (?)", db.Table("qy_blog_tags").Select("ID").Where("identifier = ?", opts.TagName)))
	}
	if len(opts.CategoryName) > 0 {
		query.Where("category_id = (?)", db.Table("qy_blog_category").Select("ID").Where("identifier = ?", opts.CategoryName))
	}
	query.Order("created_at desc,id desc")
	if len(opts.Order) > 0 {
		query.Order(opts.Order)
	}
	if opts.PageFlag {
		ol := gormutil.Unpointer(opts.Offset, opts.Limit)
		d := query.Where(where).
			Offset(ol.Offset).
			Limit(ol.Limit).
			Find(&ret.Items).
			Offset(-1).
			Limit(-1).
			Count(&ret.TotalCount)
		err = d.Error
	} else {
		d := query.Where(where).
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

	infos := make([]*biz.ArticleDO, 0, len(ret.Items))
	for _, data := range ret.Items {
		infos = append(infos, bizToArticleDO(data))
	}
	return &biz.ArticleDOList{ListMeta: ret.ListMeta, Items: infos}, err
}
func bizToArticleDO(g *po.ArticlePO) *biz.ArticleDO {
	data := &biz.ArticleDO{
		ObjectMeta:     g.ObjectMeta,
		Title:          g.Title,
		PermaLink:      g.PermaLink,
		CanonicalLink:  g.CanonicalLink,
		Summary:        g.Summary,
		Thumbnail:      g.Thumbnail,
		Password:       g.Password,
		Atype:          g.Atype,
		CategoryId:     g.CategoryId,
		CategoryName:   g.CategoryName,
		CommentAgentId: g.CommentAgentId,
		Published:      g.Published.Bool,
		ViewCount:      g.ViewCount,
		CommentCount:   g.CommentCount,
		LikeCount:      g.LikeCount,
		HateCount:      g.HateCount,
		PublishedAt:    g.PublishedAt,
		Nickname:       g.Nickname,
		CommentFlag:    g.CommentFlag.Bool,
	}
	return data
}

func (r *articleRepo) GetUserFromCache(ctx context.Context, key string) (*biz.ArticleDO, error) {
	skey := articleCacheKey(key)
	result, err := r.data.RedisCli.Get(ctx, skey).Result()
	if err != nil {
		return nil, err
	}
	if result == "null" {
		return nil, errors.New("数据超时")
	}
	var cacheUser = &biz.ArticleDO{}
	err = json.Unmarshal([]byte(result), cacheUser)
	if err != nil {
		return nil, err
	}
	return cacheUser, nil
}

func (r *articleRepo) SetArticleCache(ctx context.Context, user *biz.ArticleDO, key string) {
	skey := articleCacheKey(key)
	marshal, err := json.Marshal(user)
	log.Info("dd marshal", string(marshal))
	if err != nil {
		r.log.Errorf("fail to set ArticleDO cache:json.Marshal(%v) error(%v)", user, err)
	}
	err = r.data.RedisCli.Set(ctx, skey, string(marshal), time.Minute*30).Err()
	if err != nil {
		r.log.Errorf("fail to set ArticleDO cache:redis.Set(%v) error(%v)", user, err)
	}
}

//func (r *articleRepo) SetArticleCountCache(ctx context.Context, user *biz.ArticleDO, key string) {
//	skey := articleCacheKey(key)
//	marshal, err := json.Marshal(user)
//	if err != nil {
//		r.log.Errorf("fail to set user cache:json.Marshal(%v) error(%v)", user, err)
//	}
//
//	if err != nil {
//		r.log.Errorf("fail to set user cache:redis.Set(%v) error(%v)", user, err)
//	}
//}
