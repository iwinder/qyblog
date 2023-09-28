package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/gormutil"
	"github.com/iwinder/qyblog/internal/qycms_blog/biz"
	"github.com/iwinder/qyblog/internal/qycms_blog/data/po"
	"gorm.io/gorm/clause"
)

type fileLibTypeRepo struct {
	data *Data
	log  *log.Helper
}

// NewFileLibTypeRepo .
func NewFileLibTypeRepo(data *Data, logger log.Logger) biz.FileLibTypeRepo {
	return &fileLibTypeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *fileLibTypeRepo) Save(ctx context.Context, data *biz.FileLibTypeDO) (*biz.FileLibTypeDO, error) {
	objPO := &po.FileLibTypePO{
		ObjectMeta: data.ObjectMeta,
		Name:       data.Name,
		Identifier: data.Identifier,
		Ftype:      data.Ftype,
	}
	err := r.data.Db.Omit(clause.Associations).Create(objPO).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.FileLibTypeDO{Name: objPO.Name}
	objDO.ID = objPO.ID
	return objDO, nil
}

func (r *fileLibTypeRepo) Update(ctx context.Context, data *biz.FileLibTypeDO) (*biz.FileLibTypeDO, error) {
	objPO := &po.FileLibTypePO{
		ObjectMeta: data.ObjectMeta,
		Name:       data.Name,
		Identifier: data.Identifier,
		Ftype:      data.Ftype,
	}
	err := r.data.Db.Model(&po.FileLibTypePO{}).Where("id=?", data.ID).Updates(&objPO).Error
	if err != nil {
		return nil, err
	}

	dataDO := &biz.FileLibTypeDO{Name: objPO.Name}
	dataDO.ID = objPO.ID
	return dataDO, nil
}

func (r *fileLibTypeRepo) DeleteList(c context.Context, ids []uint64) error {
	objPO := &po.FileLibTypePO{}
	err := r.data.Db.Delete(&objPO, ids).Error
	return err
}

func (r *fileLibTypeRepo) ListAll(ctx context.Context, opts biz.FileLibTypeDOListOption) (*biz.FileLibTypeDOList, error) {
	ret := &po.FileLibTypePOList{}

	var err error
	query := r.data.Db.Model(&po.FileLibTypePO{})
	if len(opts.Name) > 0 {
		query.Scopes(withFilterKeyLikeValue("name", "%"+opts.Name+"%"))
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
	infos := make([]*biz.FileLibTypeDO, 0, len(ret.Items))
	for _, obj := range ret.Items {
		infos = append(infos, &biz.FileLibTypeDO{
			ObjectMeta: obj.ObjectMeta,
			Name:       obj.Name,
			Identifier: obj.Identifier,
			Ftype:      obj.Ftype,
		})
	}
	return &biz.FileLibTypeDOList{ListMeta: ret.ListMeta, Items: infos}, err
}
