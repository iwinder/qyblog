package db

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/gormutil"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
)

const indexLinkCacheKey = "indexLink_cache_key"
const allLinkCacheKey = "allLink_cache_key"

type linkRepo struct {
	data *Data
	log  *log.Helper
}

func NewLinkRepo(data *Data, logger log.Logger) biz.LinkRepo {
	return &linkRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *linkRepo) Save(ctx context.Context, do *biz.LinkDO) (*biz.LinkDO, error) {
	po := &po.LinkPO{
		ObjectMeta:  do.ObjectMeta,
		Name:        do.Name,
		Url:         do.Url,
		Description: do.Description,
		Ftype:       do.Ftype,
	}
	err := r.data.Db.Create(po).Error
	if err != nil {
		return nil, err
	}
	dataDO := &biz.LinkDO{Name: po.Name}
	dataDO.ID = dataDO.ID
	if do.Ftype == 1 {
		r.FindIndexLinkAllWitchCache(ctx)
	} else {
		r.FindAllWitchCache(ctx)
	}
	return dataDO, nil
}

func (r *linkRepo) Update(ctx context.Context, do *biz.LinkDO) (*biz.LinkDO, error) {
	po := &po.LinkPO{
		ObjectMeta:  do.ObjectMeta,
		Name:        do.Name,
		Url:         do.Url,
		Description: do.Description,
		Ftype:       do.Ftype,
	}
	err := r.data.Db.Updates(po).Error
	if err != nil {
		return nil, err
	}
	dataDO := &biz.LinkDO{Name: po.Name}
	dataDO.ID = dataDO.ID
	if do.Ftype == 1 {
		r.FindIndexLinkAllWitchCache(ctx)
	} else {
		r.FindAllWitchCache(ctx)
	}
	return dataDO, nil
}

func (r *linkRepo) Delete(ctx context.Context, id uint64) error {
	objPO := &po.LinkPO{}
	objPO.ID = id
	err := r.data.Db.Delete(&objPO).Error
	return err
}

func (r *linkRepo) DeleteList(c context.Context, ids []uint64) error {
	userPO := &po.LinkPO{}
	if ids == nil || len(ids) == 0 {
		return nil
	}
	err := r.data.Db.Delete(&userPO, ids).Error
	return err
}

func (r *linkRepo) FindByID(ctx context.Context, id uint64) (*biz.LinkDO, error) {
	obj := &po.LinkPO{}
	err := r.data.Db.Where("id = ?", id).First(&obj).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.LinkDO{
		ObjectMeta:  obj.ObjectMeta,
		Name:        obj.Name,
		Url:         obj.Url,
		Description: obj.Description,
		Ftype:       obj.Ftype,
	}
	return objDO, nil
}
func (r *linkRepo) FindIndexLinkAllWitchCache(ctx context.Context) ([]*biz.LinkDO, error) {
	var ret []*biz.LinkDO

	var err error
	ret, err = r.getLinkAllCache(ctx, indexLinkCacheKey)
	if err != nil || ret == nil || len(ret) == 0 {
		dataPO := &biz.LinkDOList{}
		opts := biz.LinkDOListOption{}
		opts.PageFlag = false
		opts.Ftype = 1
		dataPO, err = r.ListAll(ctx, opts)
		if err != nil {
			return nil, err
		}
		ret = dataPO.Items
		r.setLinkAllCache(ctx, ret, indexLinkCacheKey)
	}
	return ret, nil
}
func (r *linkRepo) FindAllWitchCache(ctx context.Context) ([]*biz.LinkDO, error) {
	var ret []*biz.LinkDO

	var err error
	ret, err = r.getLinkAllCache(ctx, allLinkCacheKey)
	if err != nil || ret == nil || len(ret) == 0 {
		dataPO := &biz.LinkDOList{}
		opts := biz.LinkDOListOption{}
		opts.PageFlag = false
		opts.Ftype = 1
		dataPO, err = r.ListAll(ctx, opts)
		if err != nil {
			return nil, err
		}
		ret = dataPO.Items
		r.setLinkAllCache(ctx, ret, allLinkCacheKey)
	}
	return ret, nil
}
func (r *linkRepo) ListAll(c context.Context, opts biz.LinkDOListOption) (*biz.LinkDOList, error) {
	ret := &po.LinkPOList{}

	var err error
	query := r.data.Db.Model(&po.LinkPO{})
	if len(opts.Name) > 0 {
		query.Scopes(withFilterKeyLikeValue("name", "%"+opts.Name+"%"))
	}
	if opts.StatusFlag > 0 {
		query.Scopes(withFilterKeyEquarlsValue("status_flag", opts.StatusFlag))
	}
	if opts.Ftype > 0 {
		query.Scopes(withFilterKeyEquarlsValue("ftype", opts.Ftype))
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
	infos := make([]*biz.LinkDO, 0, len(ret.Items))
	for _, obj := range ret.Items {
		infos = append(infos, &biz.LinkDO{
			ObjectMeta:  obj.ObjectMeta,
			Name:        obj.Name,
			Url:         obj.Url,
			Description: obj.Description,
			Ftype:       obj.Ftype,
		})
	}
	return &biz.LinkDOList{ListMeta: ret.ListMeta, Items: infos}, err
}

func (r *linkRepo) setLinkAllCache(ctx context.Context, data []*biz.LinkDO, key string) {

	marshal, err := json.Marshal(data)
	if err != nil {
		r.log.Errorf("fail to set data cache:json.Marshal(%v) error(%v)", data, err)
	}
	err = r.data.RedisCli.Set(ctx, key, string(marshal), -1).Err()
	if err != nil {
		r.log.Errorf("fail to set data cache:redis.Set(%v) error(%v)", data, err)
	}
}

func (r *linkRepo) getLinkAllCache(ctx context.Context, key string) ([]*biz.LinkDO, error) {
	result, err := r.data.RedisCli.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var cacheUser = make([]*biz.LinkDO, 0, 0)
	err = json.Unmarshal([]byte(result), cacheUser)
	if err != nil {
		return nil, err
	}
	return cacheUser, nil
}
