package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/auth/auth_constants"
)

type UserRoleDO struct {
	UserID uint64
	RoleID uint64
}

type UserRoleRepo interface {
	CreateInBatches(ctx context.Context, userRoles []*UserRoleDO) error
	UpdateInBatches(ctx context.Context, userRoles []*UserRoleDO) error
	DeleteByUserId(ctx context.Context, userId uint64) error
}

type UserRoleUsecase struct {
	repo      UserRoleRepo
	cabinRepo CasbinRuleRepo
	log       *log.Helper
}

// NewUserRoleUsecase new a UserDO usecase.
func NewUserRoleUsecase(repo UserRoleRepo, cabinRepo CasbinRuleRepo, logger log.Logger) *UserRoleUsecase {
	return &UserRoleUsecase{repo: repo, cabinRepo: cabinRepo, log: log.NewHelper(logger)}
}

func (uc *UserRoleUsecase) SaveRoleForUser(ctx context.Context, user *UserDO) error {

	if user.Roles != nil && len(user.Roles) > 0 {
		rlen := len(user.Roles)
		roleIdfStrs := make([]string, 0, rlen)
		userRoles := make([]*UserRoleDO, 0, rlen)
		for _, obj := range user.Roles {
			roleIdfStrs = append(roleIdfStrs, auth_constants.PrefixRole+string(obj.ID))
			userRoles = append(userRoles, &UserRoleDO{
				UserID: user.ID,
				RoleID: obj.ID,
			})
		}
		// 关联关系
		if len(userRoles) > 0 {
			// 关联
			err := uc.repo.CreateInBatches(ctx, userRoles)
			if err != nil {
				return err
			}
			// 权限
			_, ucerr := uc.cabinRepo.SaveRoleForUser(ctx, auth_constants.PrefixUser+user.Username, roleIdfStrs, "*")
			if ucerr != nil {
				return ucerr
			}
		}

	}
	return nil
}

func (uc *UserRoleUsecase) UpdateRoleForUser(ctx context.Context, user *UserDO) error {

	if user.Roles != nil && len(user.Roles) > 0 {
		rlen := len(user.Roles)
		roleIdfStrs := make([]string, 0, rlen)
		userRoles := make([]*UserRoleDO, 0, rlen)
		for _, obj := range user.Roles {
			roleIdfStrs = append(roleIdfStrs, auth_constants.PrefixRole+string(obj.ID))
			userRoles = append(userRoles, &UserRoleDO{
				UserID: user.ID,
				RoleID: obj.ID,
			})
		}
		// 关联关系
		if len(userRoles) > 0 {
			// 关联
			err := uc.repo.UpdateInBatches(ctx, userRoles)
			if err != nil {
				return err
			}

			// 权限
			_, ucerr := uc.cabinRepo.SaveRoleForUser(ctx, auth_constants.PrefixUser+user.Username, roleIdfStrs, "*")
			if ucerr != nil {
				return ucerr
			}
		}

	} else { // 删除用户-角色关系
		// 关联
		err := uc.repo.DeleteByUserId(ctx, user.ID)
		if err != nil {
			return err
		}
		_, ucerr := uc.cabinRepo.DeleteRoleForUser(ctx, auth_constants.PrefixUser+user.Username, "*")
		if ucerr != nil {
			return ucerr
		}
	}
	return nil
}
