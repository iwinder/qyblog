package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/gormutil"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
	"gorm.io/gorm/clause"
)

// MenusAdmin管理
type menusAdminRepo struct {
	data *Data
	log  *log.Helper
}

// NewMenusAdminRepo .
func NewMenusAdminRepo(data *Data, logger log.Logger) biz.MenusAdminRepo {
	return &menusAdminRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Save 创建
func (r *menusAdminRepo) Save(ctx context.Context, obj *biz.MenusAdminDO) (*po.MenusAdminPO, error) {
	objPO := &po.MenusAdminPO{
		ObjectMeta: obj.ObjectMeta,
		Level:      obj.Level,
		ParentId:   obj.ParentId,
		Path:       obj.Path,
		Name:       obj.Name,
		Hidden:     obj.Hidden,
		Component:  obj.Component,
		Sort:       obj.Sort,
	}
	err := r.data.Db.Omit(clause.Associations).Create(objPO).Error
	if err != nil {
		return nil, err
	}
	return objPO, nil
}

// Update 更新
func (r *menusAdminRepo) Update(ctx context.Context, obj *biz.MenusAdminDO) (*po.MenusAdminPO, error) {
	objPO := &po.MenusAdminPO{
		Level:     obj.Level,
		ParentId:  obj.ParentId,
		Path:      obj.Path,
		Name:      obj.Name,
		Hidden:    obj.Hidden,
		Component: obj.Component,
		Sort:      obj.Sort,
	}
	tObj := &po.MenusAdminPO{}
	tObj.ID = obj.ID
	err := r.data.Db.Model(&tObj).Where("id=?", obj.ID).Updates(&objPO).Error
	if err != nil {
		return nil, err
	}
	return objPO, nil
}

// Delete 根据ID删除
func (r *menusAdminRepo) Delete(c context.Context, id uint64) error {
	objPO := &po.MenusAdminPO{}
	objPO.ID = id
	err := r.data.Db.Delete(&objPO).Error
	return err
}

// DeleteList 根据ID批量删除
func (r *menusAdminRepo) DeleteList(c context.Context, ids []uint64) error {
	objPO := &po.MenusAdminPO{}
	err := r.data.Db.Delete(&objPO, ids).Error
	return err
}

// FindByID 根据ID查询
func (r *menusAdminRepo) FindByID(ctx context.Context, id uint64) (*po.MenusAdminPO, error) {
	obj := &po.MenusAdminPO{}
	err := r.data.Db.Where("id = ?", id).First(&obj).Error
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// ListAll 批量查询
func (r *menusAdminRepo) ListAll(c context.Context, opts biz.MenusAdminDOListOption) (*po.MenusAdminPOList, error) {
	ret := &po.MenusAdminPOList{}

	where := &po.MenusAdminPO{}

	var err error

	if opts.PageFlag {
		ol := gormutil.Unpointer(opts.Offset, opts.Limit)
		d := r.data.Db.Model(where).Where(where).
			Offset(ol.Offset).
			Limit(ol.Limit).
			Order("sort ").
			Find(&ret.Items).
			Offset(-1).
			Limit(-1).
			Count(&ret.TotalCount)
		err = d.Error
	} else {
		d := r.data.Db.Model(where).Where(where).
			Order("sort").
			Find(&ret.Items).
			Count(&ret.TotalCount)
		err = d.Error
	}
	opts.TotalCount = ret.TotalCount
	opts.IsLast()
	ret.FirstFlag = opts.FirstFlag
	ret.Page = opts.Page
	ret.PageSize = opts.PageSize
	ret.LastFlag = opts.LastFlag
	return ret, err
}
