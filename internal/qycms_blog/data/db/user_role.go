package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
)

type userRoleRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRoleRepo(data *Data, logger log.Logger) biz.UserRoleRepo {
	return &userRoleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRoleRepo) CreateInBatches(ctx context.Context, userRoles []*biz.UserRoleDO) error {
	userRolePos := make([]*po.UserRolePO, 0, len(userRoles))
	for _, obj := range userRoles {
		userRolePos = append(userRolePos, &po.UserRolePO{
			UserID: obj.UserID,
			RoleID: obj.RoleID,
		})
	}
	d := r.data.Db.Create(&userRolePos)
	return d.Error
}

func (r *userRoleRepo) UpdateInBatches(ctx context.Context, userRoles []*biz.UserRoleDO) error {
	err := r.DeleteByUserId(ctx, userRoles[0].UserID)
	if err != nil {
		return err
	}
	userRolePos := make([]*po.UserRolePO, 0, len(userRoles))
	for _, obj := range userRoles {
		userRolePos = append(userRolePos, &po.UserRolePO{
			UserID: obj.UserID,
			RoleID: obj.RoleID,
		})
	}
	d := r.data.Db.Create(&userRolePos)
	return d.Error
}

func (r *userRoleRepo) DeleteByUserId(ctx context.Context, userId uint64) error {
	err := r.data.Db.Where("user_id = ?", userId).Delete(&po.UserRolePO{}).Error
	if err != nil {
		return err
	}
	return nil
}
