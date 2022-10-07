package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
)

type roleApiRepo struct {
	data *Data
	log  *log.Helper
}

func NewRoleApiRepo(data *Data, logger log.Logger) biz.RoleApiRepo {
	return &roleApiRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *roleApiRepo) CreateInBatches(ctx context.Context, roleApis []*biz.RoleApiDO) error {
	roleApiPos := make([]*po.RoleApiPO, 0, len(roleApis))
	for _, obj := range roleApis {
		roleApiPos = append(roleApiPos, &po.RoleApiPO{
			RoleID: obj.RoleID,
			ApiID:  obj.ApiID,
		})
	}
	d := r.data.Db.Create(&roleApiPos)
	return d.Error
}

func (r *roleApiRepo) UpdateInBatches(ctx context.Context, roleApis []*biz.RoleApiDO) error {
	err := r.DeleteByRoleId(ctx, roleApis[0].RoleID)
	if err != nil {
		return err
	}
	roleApiPos := make([]*po.RoleApiPO, 0, len(roleApis))
	for _, obj := range roleApis {
		roleApiPos = append(roleApiPos, &po.RoleApiPO{
			RoleID: obj.RoleID,
			ApiID:  obj.ApiID,
		})
	}
	d := r.data.Db.Create(&roleApiPos)
	return d.Error
}

func (r *roleApiRepo) FindApiIdsByRoleId(ctx context.Context, roleId uint64) ([]uint64, error) {
	menusIds := make([]uint64, 0)
	err := r.data.Db.Model(&po.RoleApiPO{}).Select("api_id").Where("role_id = ?", roleId).Scan(&menusIds).Error
	return menusIds, err
}

func (r *roleApiRepo) DeleteByRoleId(ctx context.Context, roleId uint64) error {
	err := r.data.Db.Where("role_id = ?", roleId).Delete(&po.RoleApiPO{}).Error
	if err != nil {
		return err
	}
	return nil
}
