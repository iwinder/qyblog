package db

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/gormutil"
	"github.com/iwinder/qyblog/internal/qycms_blog/biz"
	"github.com/iwinder/qyblog/internal/qycms_blog/data/po"
)

const shortCacheKey = "shortLink_cache_key"

type shortLinkRepo struct {
	data *Data
	log  *log.Helper
}

func NewShortLinkRepo(data *Data, logger log.Logger) biz.ShortLinkRepo {
	return &shortLinkRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *shortLinkRepo) Save(ctx context.Context, do *biz.ShortLinkDO) (*biz.ShortLinkDO, error) {
	po := &po.ShortLinkPO{
		ObjectMeta:  do.ObjectMeta,
		Url:         do.Url,
		Description: do.Description,
		Identifier:  do.Identifier,
	}
	err := r.data.Db.Create(po).Error
	if err != nil {
		return nil, err
	}
	dataDO := &biz.ShortLinkDO{Identifier: po.Identifier}
	dataDO.ID = dataDO.ID
	r.setShortLinkAllCache(ctx, nil)
	return dataDO, nil
}

func (r *shortLinkRepo) Update(ctx context.Context, do *biz.ShortLinkDO) (*biz.ShortLinkDO, error) {
	po := &po.ShortLinkPO{
		ObjectMeta:  do.ObjectMeta,
		Url:         do.Url,
		Description: do.Description,
		Identifier:  do.Identifier,
	}
	err := r.data.Db.Updates(po).Error
	if err != nil {
		return nil, err
	}
	dataDO := &biz.ShortLinkDO{Identifier: po.Identifier}
	dataDO.ID = dataDO.ID
	r.setShortLinkAllCache(ctx, nil)
	return dataDO, nil
}

func (r *shortLinkRepo) Delete(ctx context.Context, id uint64) error {
	objPO := &po.ShortLinkPO{}
	objPO.ID = id
	err := r.data.Db.Delete(&objPO).Error
	return err
}

func (r *shortLinkRepo) DeleteList(c context.Context, ids []uint64) error {
	userPO := &po.ShortLinkPO{}
	if ids == nil || len(ids) == 0 {
		return nil
	}
	err := r.data.Db.Delete(&userPO, ids).Error
	return err
}

func (r *shortLinkRepo) FindByID(ctx context.Context, id uint64) (*biz.ShortLinkDO, error) {
	obj := &po.ShortLinkPO{}
	err := r.data.Db.Where("id = ?", id).First(&obj).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.ShortLinkDO{
		ObjectMeta:  obj.ObjectMeta,
		Url:         obj.Url,
		Description: obj.Description,
		Identifier:  obj.Identifier,
	}
	return objDO, nil
}
func (r *shortLinkRepo) FindAllWitchCache(ctx context.Context) ([]*biz.ShortLinkDO, error) {
	var ret []*biz.ShortLinkDO

	var err error
	ret, err = r.getShortLinkAllCache(ctx)
	if err != nil || ret == nil || len(ret) == 0 {
		dataPO := &biz.ShortLinkDOList{}
		opts := biz.ShortLinkDOListOption{}
		opts.PageFlag = false
		dataPO, err = r.ListAll(ctx, opts)
		if err != nil {
			return nil, err
		}
		ret = dataPO.Items
		r.setShortLinkAllCache(ctx, ret)
	}
	return ret, nil
}
func (r *shortLinkRepo) ListAll(c context.Context, opts biz.ShortLinkDOListOption) (*biz.ShortLinkDOList, error) {
	ret := &po.ShortLinkPOList{}

	var err error
	query := r.data.Db.Model(&po.ShortLinkPO{})
	if len(opts.Identifier) > 0 {
		query.Scopes(withFilterKeyLikeValue("name", "%"+opts.Identifier+"%"))
	}
	if opts.StatusFlag > 0 {
		query.Scopes(withFilterKeyEquarlsValue("status_flag", opts.StatusFlag))
	}
	if opts.PageFlag {
		ol := gormutil.Unpointer(opts.Offset, opts.Limit)
		d := query.
			Offset(ol.Offset).
			Limit(ol.Limit).
			Order("id ").
			Find(&ret.Items).
			Offset(-1).
			Limit(-1).
			Count(&ret.TotalCount)
		err = d.Error
	} else {
		d := query.
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
	infos := make([]*biz.ShortLinkDO, 0, len(ret.Items))
	for _, obj := range ret.Items {
		infos = append(infos, &biz.ShortLinkDO{
			ObjectMeta:  obj.ObjectMeta,
			Url:         obj.Url,
			Description: obj.Description,
			Identifier:  obj.Identifier,
		})
	}
	return &biz.ShortLinkDOList{ListMeta: ret.ListMeta, Items: infos}, err
}

func (r *shortLinkRepo) setShortLinkAllCache(ctx context.Context, data []*biz.ShortLinkDO) {

	marshal, err := json.Marshal(data)
	if err != nil {
		r.log.Errorf("fail to set data cache:json.Marshal(%v) error(%v)", data, err)
	}
	err = r.data.RedisCli.Set(ctx, shortCacheKey, string(marshal), -1).Err()
	if err != nil {
		r.log.Errorf("fail to set data cache:redis.Set(%v) error(%v)", data, err)
	}
}

func (r *shortLinkRepo) getShortLinkAllCache(ctx context.Context) ([]*biz.ShortLinkDO, error) {
	result, err := r.data.RedisCli.Get(ctx, shortCacheKey).Result()
	if err != nil {
		return nil, err
	}
	var cacheUser = make([]*biz.ShortLinkDO, 0, 0)
	err = json.Unmarshal([]byte(result), cacheUser)
	if err != nil {
		return nil, err
	}
	return cacheUser, nil
}
