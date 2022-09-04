package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/gormutil"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
	"gorm.io/gorm/clause"
)

// Api管理
type apiRepo struct {
	data *Data
	log  *log.Helper
}

// NewApiRepo .
func NewApiRepo(data *Data, logger log.Logger) biz.ApiRepo {
	return &apiRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Save 创建
func (r *apiRepo) Save(ctx context.Context, obj *biz.ApiDO) (*po.ApiPO, error) {
	objPO := &po.ApiPO{
		ObjectMeta:  obj.ObjectMeta,
		ApiGroup:    obj.ApiGroup,
		Identifier:  obj.Identifier,
		Method:      obj.Method,
		Path:        obj.Path,
		Description: obj.Description,
	}
	err := r.data.Db.Omit(clause.Associations).Create(objPO).Error
	if err != nil {
		return nil, err
	}
	return objPO, nil
}

// Update 更新
func (r *apiRepo) Update(ctx context.Context, obj *biz.ApiDO) (*po.ApiPO, error) {
	objPO := &po.ApiPO{
		ApiGroup:    obj.ApiGroup,
		Identifier:  obj.Identifier,
		Method:      obj.Method,
		Path:        obj.Path,
		Description: obj.Description,
	}
	tObj := &po.ApiPO{}
	tObj.ID = obj.ID
	err := r.data.Db.Model(&tObj).Where("id=?", obj.ID).Updates(&objPO).Error
	if err != nil {
		return nil, err
	}
	return objPO, nil
}

// Delete 根据ID删除
func (r *apiRepo) Delete(c context.Context, id uint64) error {
	objPO := &po.ApiPO{}
	objPO.ID = id
	err := r.data.Db.Delete(&objPO).Error
	return err
}

// DeleteList 根据ID批量删除
func (r *apiRepo) DeleteList(c context.Context, ids []uint64) error {
	objPO := &po.ApiPO{}
	err := r.data.Db.Delete(&objPO, ids).Error
	return err
}

// FindByID 根据ID查询
func (r *apiRepo) FindByID(ctx context.Context, id uint64) (*po.ApiPO, error) {
	obj := &po.ApiPO{}
	err := r.data.Db.Where("id = ?", id).First(&obj).Error
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// ListAll 批量查询
func (r *apiRepo) ListAll(c context.Context, opts biz.ApiDOListOption) (*po.ApiPOList, error) {
	ret := &po.ApiPOList{}

	where := &po.ApiPO{}
	var err error

	if opts.PageFlag {
		ol := gormutil.Unpointer(opts.Offset, opts.Limit)
		d := r.data.Db.Model(where).Where(where).
			Offset(ol.Offset).
			Limit(ol.Limit).
			Order("api_group,id").
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
	ret.Page = opts.Page
	ret.PageSize = opts.PageSize
	ret.LastFlag = opts.LastFlag
	return ret, err
}
