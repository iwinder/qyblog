package db

import (
	"context"
	"database/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/gormutil"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qyblog/internal/qycms_blog/biz"
	"github.com/iwinder/qyblog/internal/qycms_blog/data/po"
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
func (r *apiRepo) Save(ctx context.Context, obj *biz.ApiDO) (*biz.ApiDO, error) {
	objPO := &po.ApiPO{
		ObjectMeta:  obj.ObjectMeta,
		ApiGroup:    obj.ApiGroup,
		Identifier:  obj.Identifier,
		Method:      obj.Method,
		Path:        obj.Path,
		Description: obj.Description,
		GroupId: sql.NullInt64{
			Int64: int64(obj.GroupId),
			Valid: true,
		},
	}
	err := r.data.Db.Omit(clause.Associations).Create(objPO).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.ApiDO{Path: objPO.Path}
	objDO.ID = objPO.ID
	return objDO, nil
}

// Update 更新
func (r *apiRepo) Update(ctx context.Context, obj *biz.ApiDO) (*biz.ApiDO, error) {
	objPO := &po.ApiPO{
		ApiGroup:    obj.ApiGroup,
		Identifier:  obj.Identifier,
		Method:      obj.Method,
		Path:        obj.Path,
		Description: obj.Description,
		GroupId: sql.NullInt64{
			Int64: int64(obj.GroupId),
			Valid: true,
		},
	}
	tObj := &po.ApiPO{}
	tObj.ID = obj.ID
	err := r.data.Db.Model(&tObj).Where("id=?", obj.ID).Updates(&objPO).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.ApiDO{Path: objPO.Path}
	objDO.ID = objPO.ID
	return objDO, nil
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
func (r *apiRepo) FindByID(ctx context.Context, id uint64) (*biz.ApiDO, error) {
	obj := &po.ApiPO{}
	err := r.data.Db.Where("id = ?", id).First(&obj).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.ApiDO{
		ObjectMeta:  obj.ObjectMeta,
		ApiGroup:    obj.ApiGroup,
		Identifier:  obj.Identifier,
		Method:      obj.Method,
		Path:        obj.Path,
		Description: obj.Description,
		GroupId:     uint64(obj.GroupId.Int64),
	}
	return objDO, nil
}

// ListAll 批量查询
func (r *apiRepo) ListAll(c context.Context, opts biz.ApiDOListOption) (*biz.ApiDOList, error) {
	ret := &po.ApiPOList{}

	where := &po.ApiPO{}
	var err error
	queryDB := r.data.Db.Model(where)
	if len(opts.ApiGroup) > 0 {
		queryDB.Scopes(withFilterKeyLikeValue("api_group", "%"+opts.ApiGroup+"%"))
	}
	if len(opts.Path) > 0 {
		queryDB.Scopes(withFilterKeyLikeValue("path", "%"+opts.Path+"%"))
	}
	if len(opts.Method) > 0 {
		queryDB.Scopes(withFilterKeyEquarlsValue("method", opts.Method))
	}
	if len(opts.Description) > 0 {
		queryDB.Scopes(withFilterKeyLikeValue("description", "%"+opts.Description+"%"))
	}
	if len(opts.Identifier) > 0 {
		queryDB.Scopes(withFilterKeyLikeValue("identifier", "%"+opts.Identifier+"%"))
	}
	if opts.GroupId > 0 {
		queryDB.Scopes(withFilterKeyEquarlsValue("group_id", opts.GroupId))
	}
	if opts.PageFlag {
		ol := gormutil.Unpointer(opts.Offset, opts.Limit)
		d := queryDB.Where(where).
			Offset(ol.Offset).
			Limit(ol.Limit).
			Order("api_group,id").
			Find(&ret.Items).
			Offset(-1).
			Limit(-1).
			Count(&ret.TotalCount)
		err = d.Error
	} else {
		d := queryDB.Where(where).
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

	infos := make([]*biz.ApiDO, 0, len(ret.Items))
	for _, obj := range ret.Items {
		infos = append(infos, &biz.ApiDO{
			ObjectMeta: metaV1.ObjectMeta{
				ID:         obj.ID,
				InstanceID: obj.InstanceID,
				Extend:     obj.Extend,
				CreatedAt:  obj.CreatedAt,
				UpdatedAt:  obj.UpdatedAt,
			},
			ApiGroup:    obj.ApiGroup,
			Identifier:  obj.Identifier,
			Method:      obj.Method,
			Path:        obj.Path,
			Description: obj.Description,
			GroupId:     uint64(obj.GroupId.Int64),
		})
	}
	return &biz.ApiDOList{ListMeta: ret.ListMeta, Items: infos}, err
}
