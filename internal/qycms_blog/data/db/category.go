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

type categoryRepo struct {
	data *Data
	log  *log.Helper
}

// NewCategoryRepo .
func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	return &categoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Save 创建
func (r *categoryRepo) Save(ctx context.Context, obj *biz.CategoryDO) (*biz.CategoryDO, error) {
	objPO := &po.CategoryPO{
		ObjectMeta:  obj.ObjectMeta,
		Name:        obj.Name,
		Identifier:  obj.Identifier,
		Description: obj.Description,
		ParentId: sql.NullInt64{
			Int64: int64(obj.ParentId),
			Valid: true,
		},
	}
	err := r.data.Db.Omit(clause.Associations).Create(objPO).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.CategoryDO{Name: objPO.Name}
	objDO.ID = objPO.ID
	return objDO, nil
}

// Update 更新
func (r *categoryRepo) Update(ctx context.Context, obj *biz.CategoryDO) (*biz.CategoryDO, error) {
	objPO := &po.CategoryPO{
		Name:        obj.Name,
		Identifier:  obj.Identifier,
		Description: obj.Description,
		ParentId: sql.NullInt64{
			Int64: int64(obj.ParentId),
			Valid: true,
		},
	}
	err := r.data.Db.Model(&objPO).Where("id=?", obj.ID).Updates(&objPO).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.CategoryDO{Name: objPO.Name}
	objDO.ID = objPO.ID
	return objDO, nil
}

// Delete 根据ID删除
func (r *categoryRepo) Delete(c context.Context, id uint64) error {
	objPO := &po.CategoryPO{}
	objPO.ID = id
	err := r.data.Db.Delete(&objPO).Error
	return err
}

// DeleteList 根据ID批量删除
func (r *categoryRepo) DeleteList(c context.Context, ids []uint64) error {
	objPO := &po.CategoryPO{}
	err := r.data.Db.Delete(&objPO, ids).Error
	return err
}

// FindByID 根据ID查询
func (r *categoryRepo) FindByID(ctx context.Context, id uint64) (*biz.CategoryDO, error) {
	obj := &po.CategoryPO{}
	err := r.data.Db.Where("id = ?", id).First(&obj).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.CategoryDO{
		ObjectMeta: obj.ObjectMeta,
		Name:       obj.Name,
		Identifier: obj.Identifier,
	}
	return objDO, nil
}
func (r *categoryRepo) FindByIdentifier(ctx context.Context, name string) (*biz.CategoryDO, error) {
	obj := &po.CategoryPO{}
	err := r.data.Db.Where("identifier = ?", name).First(&obj).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.CategoryDO{
		ObjectMeta: obj.ObjectMeta,
		Name:       obj.Name,
		Identifier: obj.Identifier,
	}
	return objDO, nil
}

func (r *categoryRepo) CountByIdentifier(ctx context.Context, str string) (int64, error) {
	var obj int64
	err := r.data.Db.Model(&po.CategoryPO{}).Where("identifier like ?", str+"%").Count(&obj).Error
	if err != nil {
		return 0, err
	}
	return obj, nil
}

// ListAll 批量查询
func (r *categoryRepo) ListAll(c context.Context, opts biz.CategoryDOListOption) (*biz.CategoryDOList, error) {
	ret := &po.CategoryPOList{}

	where := &po.CategoryPO{}
	var err error
	queryDB := r.data.Db.Model(where)
	if len(opts.Name) > 0 {
		queryDB.Scopes(withFilterKeyLikeValue("name", "%"+opts.Name+"%"))
	}
	if opts.ParentId >= 0 {
		queryDB.Scopes(withFilterKeyEquarlsValue("parent_id", opts.ParentId))
	}
	if len(opts.Identifier) > 0 {
		queryDB.Scopes(withFilterKeyLikeValue("identifier", "%"+opts.Identifier+"%"))
	}
	if opts.PageFlag {
		ol := gormutil.Unpointer(opts.Offset, opts.Limit)
		d := queryDB.Where(where).
			Offset(ol.Offset).
			Limit(ol.Limit).
			Order("id").
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

	infos := make([]*biz.CategoryDO, 0, len(ret.Items))
	for _, obj := range ret.Items {
		infos = append(infos, &biz.CategoryDO{
			ObjectMeta: metaV1.ObjectMeta{
				ID:         obj.ID,
				InstanceID: obj.InstanceID,
				Extend:     obj.Extend,
				CreatedAt:  obj.CreatedAt,
				UpdatedAt:  obj.UpdatedAt,
			},
			Name:        obj.Name,
			Identifier:  obj.Identifier,
			Description: obj.Description,
			ParentId:    uint64(obj.ParentId.Int64),
		})
	}
	return &biz.CategoryDOList{ListMeta: ret.ListMeta, Items: infos}, err
}

func (r *categoryRepo) ListAllWithChildren(ctx context.Context, opts biz.CategoryDOListOption) (*biz.CategoryDOList, error) {
	dataDOs, err := r.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	newopts := biz.CategoryDOListOption{}
	newopts.PageFlag = false
	newopts.Name = opts.Name
	for _, obj := range dataDOs.Items {
		newopts.ParentId = obj.ID
		cobjPOs, aerr := r.ListAllChildren(ctx, newopts)
		obj.Children = cobjPOs
		if aerr != nil {
			log.Errorf("菜单列表获取异常%v", aerr)
			obj.Children = make([]*biz.CategoryDO, 0, 0)
		}
	}
	return dataDOs, nil
}

func (uc *categoryRepo) ListAllChildren(ctx context.Context, opts biz.CategoryDOListOption) ([]*biz.CategoryDO, error) {
	objDOList, err := uc.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	newopts := biz.CategoryDOListOption{}
	newopts.PageFlag = false
	newopts.Name = opts.Name
	for _, obj := range objDOList.Items {
		newopts.ParentId = obj.ID
		cobjPOs, aerr := uc.ListAllChildren(ctx, newopts)
		obj.Children = cobjPOs
		if aerr != nil {
			log.Errorf("ListAllChildren 分类列表获取异常%v", aerr)
			obj.Children = make([]*biz.CategoryDO, 0, 0)
		}
	}
	return objDOList.Items, nil
}
