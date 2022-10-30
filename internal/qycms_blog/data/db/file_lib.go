package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/gormutil"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
	"gorm.io/gorm/clause"
)

type fileLibRepo struct {
	data *Data
	log  *log.Helper
}

// NewFileLibRepo .
func NewFileLibRepo(data *Data, logger log.Logger) biz.FileLibRepo {
	return &fileLibRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *fileLibRepo) Save(ctx context.Context, data *biz.FileLibDO) (*biz.FileLibDO, error) {
	objPO := &po.FileLibPO{
		ObjectMeta:     data.ObjectMeta,
		OriginFileName: data.OriginFileName,
		Fname:          data.Fname,
		Fsize:          data.Fsize,
		Extention:      data.Extention,
		MimeType:       data.MimeType,
		Fmd5:           data.Fmd5,
		RelativePath:   data.RelativePath,
	}
	err := r.data.Db.Omit(clause.Associations).Create(objPO).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.FileLibDO{OriginFileName: objPO.OriginFileName}
	objDO.ID = objPO.ID
	return objDO, nil
}

func (r *fileLibRepo) Update(ctx context.Context, data *biz.FileLibDO) (*biz.FileLibDO, error) {
	objPO := &po.FileLibPO{
		ObjectMeta:     data.ObjectMeta,
		OriginFileName: data.OriginFileName,
		Fname:          data.Fname,
		Fsize:          data.Fsize,
		Extention:      data.Extention,
		MimeType:       data.MimeType,
		Fmd5:           data.Fmd5,
		RelativePath:   data.RelativePath,
	}
	err := r.data.Db.Model(&po.FileLibPO{}).Where("id=?", data.ID).Updates(&objPO).Error
	if err != nil {
		return nil, err
	}

	dataDO := &biz.FileLibDO{OriginFileName: objPO.OriginFileName}
	dataDO.ID = objPO.ID
	return dataDO, nil
}
func (r *fileLibRepo) FindByMd5(ctx context.Context, fmd5 string) (*biz.FileLibDO, error) {
	obj := &po.FileLibPO{}
	err := r.data.Db.Where("fmd5 = ?", fmd5).First(&obj).Error
	if err != nil {
		return nil, err
	}
	objDo := &biz.FileLibDO{
		ObjectMeta:     obj.ObjectMeta,
		OriginFileName: obj.OriginFileName,
		Fname:          obj.Fname,
		Fsize:          obj.Fsize,
		Extention:      obj.Extention,
		MimeType:       obj.MimeType,
		Fmd5:           obj.Fmd5,
		RelativePath:   obj.RelativePath,
	}
	return objDo, nil
}
func (r *fileLibRepo) CountByMd5(ctx context.Context, fmd5 string) (int64, error) {
	var obj int64
	err := r.data.Db.Model(&po.FileLibPO{}).Where("fmd5 = ?", fmd5).Count(&obj).Error
	if err != nil {
		return 0, err
	}
	return obj, nil
}
func (r *fileLibRepo) DeleteList(c context.Context, ids []uint64) error {
	objPO := &po.FileLibPO{}
	err := r.data.Db.Delete(&objPO, ids).Error
	return err
}

func (r *fileLibRepo) ListAll(ctx context.Context, opts biz.FileLibDOListOption) (*biz.FileLibDOList, error) {
	ret := &po.FileLibPOList{}

	var err error
	query := r.data.Db.Model(&po.FileLibPO{})
	if len(opts.OriginFileName) > 0 {
		query.Scopes(withFilterKeyLikeValue("origin_fileName", "%"+opts.OriginFileName+"%"))
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
	infos := make([]*biz.FileLibDO, 0, len(ret.Items))
	for _, obj := range ret.Items {
		infos = append(infos, &biz.FileLibDO{
			ObjectMeta:     obj.ObjectMeta,
			OriginFileName: obj.OriginFileName,
			Fname:          obj.Fname,
			Fsize:          obj.Fsize,
			Extention:      obj.Extention,
			MimeType:       obj.MimeType,
			Fmd5:           obj.Fmd5,
			Fhash:          obj.Fmd5,
			RelativePath:   obj.RelativePath,
		})
	}
	return &biz.FileLibDOList{ListMeta: ret.ListMeta, Items: infos}, err
}
