package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/gormutil"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
)

var menusCacheKey = func(targetId string) string {
	return "menus_cache_key_" + targetId
}

type menusRepo struct {
	data *Data
	log  *log.Helper
}

func NewMenusRepo(data *Data, logger log.Logger) biz.MenusRepo {
	return &menusRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *menusRepo) Save(ctx context.Context, do *biz.MenusDO) (*biz.MenusDO, error) {
	po := &po.MenusPO{
		ObjectMeta: do.ObjectMeta,
		Name:       do.Name,
		Url:        do.Url,
		Blanked:    do.Blanked,
		ParentId: sql.NullInt64{
			Int64: int64(do.ParentId),
			Valid: true,
		},
		TargetId: sql.NullInt64{
			Int64: int64(do.TargetId),
			Valid: true,
		},
	}
	err := r.data.Db.Create(po).Error
	if err != nil {
		return nil, err
	}
	dataDO := &biz.MenusDO{Name: po.Name}
	dataDO.ID = dataDO.ID
	r.reSetMenusAllCache(ctx, do.TargetId)
	return dataDO, nil
}

func (r *menusRepo) Update(ctx context.Context, do *biz.MenusDO) (*biz.MenusDO, error) {
	po := &po.MenusPO{
		ObjectMeta: do.ObjectMeta,
		Name:       do.Name,
		Url:        do.Url,
		Blanked:    do.Blanked,
		ParentId: sql.NullInt64{
			Int64: int64(do.ParentId),
			Valid: true,
		},
		TargetId: sql.NullInt64{
			Int64: int64(do.TargetId),
			Valid: true,
		},
	}
	err := r.data.Db.Updates(po).Error
	if err != nil {
		return nil, err
	}
	dataDO := &biz.MenusDO{Name: po.Name}
	dataDO.ID = dataDO.ID
	r.reSetMenusAllCache(ctx, do.TargetId)
	return dataDO, nil
}

func (r *menusRepo) Delete(ctx context.Context, id uint64) error {
	objPO := &po.MenusPO{}
	objPO.ID = id
	err := r.data.Db.Delete(&objPO).Error
	return err
}

func (r *menusRepo) DeleteList(ctx context.Context, ids []uint64, targetId uint64) error {
	userPO := &po.MenusPO{}
	if ids == nil || len(ids) == 0 {
		return nil
	}
	err := r.data.Db.Delete(&userPO, ids).Error
	r.reSetMenusAllCache(ctx, targetId)
	return err
}

func (r *menusRepo) FindByID(ctx context.Context, id uint64) (*biz.MenusDO, error) {
	obj := &po.MenusPO{}
	err := r.data.Db.Where("id = ?", id).First(&obj).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.MenusDO{
		ObjectMeta: obj.ObjectMeta,
		Name:       obj.Name,
		Url:        obj.Url,
		Blanked:    obj.Blanked,
		ParentId:   uint64(obj.ParentId.Int64),
		TargetId:   uint64(obj.TargetId.Int64),
	}
	return objDO, nil
}
func (r *menusRepo) FindMenusAllByTargetIdWitchCache(ctx context.Context, targetId uint64) ([]*biz.MenusDO, error) {
	var ret []*biz.MenusDO

	var err error
	cacheKey := menusCacheKey(fmt.Sprintf("%d", targetId))
	ret, err = r.getMenusAllCache(ctx, cacheKey)
	if err != nil || ret == nil || len(ret) == 0 {
		ret, err = r.reSetMenusAllCache(ctx, targetId)
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}
func (r *menusRepo) ListAllWithChildren(ctx context.Context, opts biz.MenusDOListOption) (*biz.MenusDOList, error) {
	dataDOs, err := r.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	newopts := biz.MenusDOListOption{}
	newopts.TargetId = opts.TargetId
	newopts.PageFlag = false
	for _, obj := range dataDOs.Items {
		newopts.ParentId = obj.ID
		cobjPOs, aerr := r.ListAllChildren(ctx, newopts)
		obj.Children = cobjPOs
		if aerr != nil {
			log.Errorf("菜单列表获取异常%v", aerr)
			obj.Children = make([]*biz.MenusDO, 0, 0)
		}
	}
	return dataDOs, nil
}
func (r *menusRepo) ListAll(c context.Context, opts biz.MenusDOListOption) (*biz.MenusDOList, error) {
	ret := &po.MenusPOList{}

	var err error
	query := r.data.Db.Model(&po.MenusPO{})
	if len(opts.Name) > 0 {
		query.Scopes(withFilterKeyLikeValue("name", "%"+opts.Name+"%"))
	}
	if opts.StatusFlag > 0 {
		query.Scopes(withFilterKeyEquarlsValue("status_flag", opts.StatusFlag))
	}
	if opts.ParentId >= 0 {
		query.Scopes(withFilterKeyEquarlsValue("parent_id", opts.ParentId))
	}
	if opts.TargetId > 0 {
		query.Scopes(withFilterKeyEquarlsValue("target_id", opts.TargetId))
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
	infos := make([]*biz.MenusDO, 0, len(ret.Items))
	for _, obj := range ret.Items {
		infos = append(infos, &biz.MenusDO{
			ObjectMeta: obj.ObjectMeta,
			Name:       obj.Name,
			Url:        obj.Url,
			Blanked:    obj.Blanked,
			ParentId:   uint64(obj.ParentId.Int64),
			TargetId:   uint64(obj.TargetId.Int64),
		})
	}
	return &biz.MenusDOList{ListMeta: ret.ListMeta, Items: infos}, err
}

func (r *menusRepo) setMenusAllCache(ctx context.Context, data []*biz.MenusDO, key string) {

	marshal, err := json.Marshal(data)
	if err != nil {
		r.log.Errorf("fail to set data cache:json.Marshal(%v) error(%v)", data, err)
	}
	err = r.data.RedisCli.Set(ctx, key, string(marshal), -1).Err()
	if err != nil {
		r.log.Errorf("fail to set data cache:redis.Set(%v) error(%v)", data, err)
	}
}

func (r *menusRepo) reSetMenusAllCache(ctx context.Context, targetId uint64) ([]*biz.MenusDO, error) {
	opts := biz.MenusDOListOption{}
	opts.PageFlag = false
	opts.TargetId = targetId
	opts.ParentId = 0
	dataPO, err := r.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	ret := dataPO.Items

	for _, obj := range ret {
		opts.ParentId = obj.ID
		cobjPOs, aerr := r.ListAllChildren(ctx, opts)
		obj.Children = cobjPOs
		if aerr != nil {
			log.Errorf("菜单列表获取异常%v", aerr)
			obj.Children = make([]*biz.MenusDO, 0, 0)
		}
	}
	cacheKey := menusCacheKey(fmt.Sprintf("%d", targetId))
	r.setMenusAllCache(ctx, ret, cacheKey)
	return ret, nil
}

func (r *menusRepo) getMenusAllCache(ctx context.Context, key string) ([]*biz.MenusDO, error) {
	result, err := r.data.RedisCli.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var cacheUser = make([]*biz.MenusDO, 0, 0)
	err = json.Unmarshal([]byte(result), cacheUser)
	if err != nil {
		return nil, err
	}
	return cacheUser, nil
}
func (uc *menusRepo) ListAllChildren(ctx context.Context, opts biz.MenusDOListOption) ([]*biz.MenusDO, error) {
	objDOList, err := uc.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	newopts := biz.MenusDOListOption{}
	newopts.PageFlag = false
	if opts.TargetId > 0 {
		newopts.TargetId = opts.TargetId
	}
	for _, obj := range objDOList.Items {
		newopts.ParentId = obj.ID
		cobjPOs, aerr := uc.ListAllChildren(ctx, newopts)
		obj.Children = cobjPOs
		if aerr != nil {
			log.Errorf("ListAllChildren 菜单列表获取异常%v", aerr)
			obj.Children = make([]*biz.MenusDO, 0, 0)
		}
	}
	return objDOList.Items, nil
}
