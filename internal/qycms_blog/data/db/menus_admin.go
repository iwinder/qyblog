package db

import (
	"context"
	"database/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/gormutil"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
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
func (r *menusAdminRepo) Save(ctx context.Context, obj *biz.MenusAdminDO) (*biz.MenusAdminDO, error) {
	objPO := &po.MenusAdminPO{
		ObjectMeta:     obj.ObjectMeta,
		Name:           obj.Name,
		BreadcrumbName: obj.BreadcrumbName,
		Identifier:     obj.Identifier,
		ParentId: sql.NullInt64{
			Int64: int64(obj.ParentId),
			Valid: true,
		},
		Icon:      obj.Icon,
		MType:     obj.MType,
		Path:      obj.Path,
		Redirect:  obj.Redirect,
		Component: sql.NullString{String: obj.Component, Valid: true},
		Sort:      obj.Sort,
	}
	if len(objPO.BreadcrumbName) == 0 {
		objPO.BreadcrumbName = objPO.Name
	}
	err := r.data.Db.Omit(clause.Associations).Create(objPO).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.MenusAdminDO{Name: objPO.Name}
	objDO.ID = objPO.ID
	return objDO, nil
}

// Update 更新
func (r *menusAdminRepo) Update(ctx context.Context, obj *biz.MenusAdminDO) (*biz.MenusAdminDO, error) {
	objPO := &po.MenusAdminPO{
		ObjectMeta:     obj.ObjectMeta,
		Name:           obj.Name,
		BreadcrumbName: obj.BreadcrumbName,
		Identifier:     obj.Identifier,
		ParentId: sql.NullInt64{
			Int64: int64(obj.ParentId),
			Valid: true,
		},
		Icon:      obj.Icon,
		MType:     obj.MType,
		Path:      obj.Path,
		Redirect:  obj.Redirect,
		Component: sql.NullString{String: obj.Component, Valid: true},
		Sort:      obj.Sort,
	}
	if len(objPO.BreadcrumbName) == 0 {
		objPO.BreadcrumbName = objPO.Name
	}
	tObj := &po.MenusAdminPO{}
	tObj.ID = obj.ID
	err := r.data.Db.Model(&tObj).Where("id=?", obj.ID).Updates(&objPO).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.MenusAdminDO{Name: objPO.Name}
	objDO.ID = objPO.ID
	return objDO, nil
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
func (r *menusAdminRepo) FindByID(ctx context.Context, id uint64) (*biz.MenusAdminDO, error) {
	obj := &po.MenusAdminPO{}
	err := r.data.Db.Where("id = ?", id).First(&obj).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.MenusAdminDO{
		ObjectMeta:     obj.ObjectMeta,
		Name:           obj.Name,
		BreadcrumbName: obj.BreadcrumbName,
		Identifier:     obj.Identifier,
		ParentId:       uint64(obj.ParentId.Int64),
		Icon:           obj.Icon,
		MType:          obj.MType,
		Path:           obj.Path,
		Redirect:       obj.Redirect,
		Component:      obj.Component.String,
		Sort:           obj.Sort,
	}
	return objDO, nil
}

// ListAll 批量查询
func (r *menusAdminRepo) ListAll(c context.Context, opts biz.MenusAdminDOListOption) (*biz.MenusAdminDOList, error) {
	ret := &po.MenusAdminPOList{}

	where := &po.MenusAdminPO{}

	var err error
	queryDB := r.data.Db.Model(where)
	if opts.StatusFlag > 0 {
		queryDB.Scopes(withFilterKeyEquarlsValue("status_flag", opts.StatusFlag))
	}
	if opts.MType > 0 {
		queryDB.Scopes(withFilterKeyEquarlsValue("mtype", opts.MType))
	}
	if opts.ParentId >= 0 {
		queryDB.Scopes(withFilterKeyEquarlsValue("parent_id", opts.ParentId))
	}
	if len(opts.Redirect) > 0 {
		queryDB.Scopes(withFilterKeyEquarlsValue("redirect", opts.Redirect))
	}
	if opts.PageFlag {
		ol := gormutil.Unpointer(opts.Offset, opts.Limit)
		d := queryDB.Where(where).
			Offset(ol.Offset).
			Limit(ol.Limit).
			Order("sort ").
			Find(&ret.Items).
			Offset(-1).
			Limit(-1).
			Count(&ret.TotalCount)
		err = d.Error
	} else {
		d := queryDB.Where(where).
			Order("sort").
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

	infos := make([]*biz.MenusAdminDO, 0, len(ret.Items))
	for _, obj := range ret.Items {
		infos = append(infos, &biz.MenusAdminDO{
			ObjectMeta: metaV1.ObjectMeta{
				ID:         obj.ID,
				InstanceID: obj.InstanceID,
				Extend:     obj.Extend,
				CreatedAt:  obj.CreatedAt,
				UpdatedAt:  obj.UpdatedAt,
			},
			Name:           obj.Name,
			BreadcrumbName: obj.BreadcrumbName,
			Identifier:     obj.Identifier,
			ParentId:       uint64(obj.ParentId.Int64),
			Icon:           obj.Icon,
			MType:          obj.MType,
			Path:           obj.Path,
			Redirect:       obj.Redirect,
			Component:      obj.Component.String,
			Sort:           obj.Sort,
		})
	}
	return &biz.MenusAdminDOList{ListMeta: ret.ListMeta, Items: infos}, err
}
