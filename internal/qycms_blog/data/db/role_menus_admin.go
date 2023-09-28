package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qyblog/internal/qycms_blog/biz"
	"github.com/iwinder/qyblog/internal/qycms_blog/data/po"
)

type roleMenusRepo struct {
	data *Data
	log  *log.Helper
}

func NewRoleMenusRepo(data *Data, logger log.Logger) biz.RoleMenusRepo {
	return &roleMenusRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *roleMenusRepo) CreateInBatches(ctx context.Context, roleMenuss []*biz.RoleMenusDO) error {
	roleMenusPos := make([]*po.RoleMenusPO, 0, len(roleMenuss))
	for _, obj := range roleMenuss {
		roleMenusPos = append(roleMenusPos, &po.RoleMenusPO{
			RoleID:  obj.RoleID,
			MenusID: obj.MenusID,
		})
	}
	d := r.data.Db.Create(&roleMenusPos)
	return d.Error
}

func (r *roleMenusRepo) UpdateInBatches(ctx context.Context, roleMenuss []*biz.RoleMenusDO) error {
	err := r.DeleteByRoleId(ctx, roleMenuss[0].RoleID)
	if err != nil {
		return err
	}
	roleMenusPos := make([]*po.RoleMenusPO, 0, len(roleMenuss))
	for _, obj := range roleMenuss {
		roleMenusPos = append(roleMenusPos, &po.RoleMenusPO{
			RoleID:  obj.RoleID,
			MenusID: obj.MenusID,
		})
	}
	d := r.data.Db.Create(&roleMenusPos)
	return d.Error
}

func (r *roleMenusRepo) FindMenusIdsByRoleId(ctx context.Context, roleId uint64) ([]uint64, error) {
	menusIds := make([]uint64, 0)
	err := r.data.Db.Model(&po.RoleMenusPO{}).Select("menus_id").Where("role_id = ?", roleId).Scan(&menusIds).Error
	return menusIds, err
}

func (r *roleMenusRepo) DeleteByRoleId(ctx context.Context, roleId uint64) error {
	err := r.data.Db.Where("role_id = ?", roleId).Delete(&po.RoleMenusPO{}).Error
	if err != nil {
		return err
	}
	return nil
}
