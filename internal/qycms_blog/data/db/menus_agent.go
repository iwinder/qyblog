package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/gormutil"
	"github.com/iwinder/qyblog/internal/qycms_blog/biz"
	"github.com/iwinder/qyblog/internal/qycms_blog/data/po"
)

type menusAgentRepo struct {
	data *Data
	log  *log.Helper
}

func NewMenusAgentRepo(data *Data, logger log.Logger) biz.MenusAgentRepo {
	return &menusAgentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *menusAgentRepo) Save(ctx context.Context, do *biz.MenusAgentDO) (*biz.MenusAgentDO, error) {
	po := &po.MenusAgentPO{
		ObjectMeta: do.ObjectMeta,
		Name:       do.Name,
		Ftype:      do.Ftype,
	}
	err := r.data.Db.Create(po).Error
	if err != nil {
		return nil, err
	}
	dataDO := &biz.MenusAgentDO{Name: po.Name}
	dataDO.ID = dataDO.ID

	return dataDO, nil
}

func (r *menusAgentRepo) Update(ctx context.Context, do *biz.MenusAgentDO) (*biz.MenusAgentDO, error) {
	po := &po.MenusAgentPO{
		ObjectMeta: do.ObjectMeta,
		Name:       do.Name,
		Ftype:      do.Ftype,
	}
	err := r.data.Db.Updates(po).Error
	if err != nil {
		return nil, err
	}
	dataDO := &biz.MenusAgentDO{Name: po.Name}
	dataDO.ID = dataDO.ID
	return dataDO, nil
}

func (r *menusAgentRepo) Delete(ctx context.Context, id uint64) error {
	objPO := &po.MenusAgentPO{}
	objPO.ID = id
	err := r.data.Db.Delete(&objPO).Error
	return err
}

func (r *menusAgentRepo) DeleteList(c context.Context, ids []uint64) error {
	userPO := &po.MenusAgentPO{}
	if ids == nil || len(ids) == 0 {
		return nil
	}
	err := r.data.Db.Delete(&userPO, ids).Error
	return err
}

func (r *menusAgentRepo) FindByID(ctx context.Context, id uint64) (*biz.MenusAgentDO, error) {
	obj := &po.MenusAgentPO{}
	err := r.data.Db.Where("id = ?", id).First(&obj).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.MenusAgentDO{
		ObjectMeta: obj.ObjectMeta,
		Name:       obj.Name,
		Ftype:      obj.Ftype,
	}
	return objDO, nil
}

func (r *menusAgentRepo) ListAll(c context.Context, opts biz.MenusAgentDOListOption) (*biz.MenusAgentDOList, error) {
	ret := &po.MenusAgentPOList{}

	var err error
	query := r.data.Db.Model(&po.MenusAgentPO{})
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
	infos := make([]*biz.MenusAgentDO, 0, len(ret.Items))
	for _, obj := range ret.Items {
		infos = append(infos, &biz.MenusAgentDO{
			ObjectMeta: obj.ObjectMeta,
			Name:       obj.Name,
			Ftype:      obj.Ftype,
		})
	}
	return &biz.MenusAgentDOList{ListMeta: ret.ListMeta, Items: infos}, err
}
