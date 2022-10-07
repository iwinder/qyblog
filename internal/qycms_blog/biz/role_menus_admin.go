package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type RoleMenusDO struct {
	RoleID  uint64
	MenusID uint64
}

type RoleMenusRepo interface {
	CreateInBatches(ctx context.Context, userRoles []*RoleMenusDO) error
	UpdateInBatches(ctx context.Context, userRoles []*RoleMenusDO) error
	DeleteByRoleId(ctx context.Context, roleId uint64) error
	FindMenusIdsByRoleId(ctx context.Context, roleId uint64) ([]uint64, error)
}

type RoleMenusUsecase struct {
	repo RoleMenusRepo
	log  *log.Helper
}

// NewRoleMenusUsecase new a UserDO usecase.
func NewRoleMenusUsecase(repo RoleMenusRepo, logger log.Logger) *RoleMenusUsecase {
	return &RoleMenusUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *RoleMenusUsecase) SaveMenusForRole(ctx context.Context, role *RoleDO) error {

	if role.MenusIDs != nil && len(role.MenusIDs) > 0 {
		rlen := len(role.MenusIDs)
		userRoles := make([]*RoleMenusDO, 0, rlen)
		for _, obj := range role.MenusIDs {
			userRoles = append(userRoles, &RoleMenusDO{
				RoleID:  role.ID,
				MenusID: obj,
			})
		}
		// 关联关系
		if len(userRoles) > 0 {
			// 关联
			err := uc.repo.CreateInBatches(ctx, userRoles)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func (uc *RoleMenusUsecase) UpdateRoleForUser(ctx context.Context, role *RoleDO) error {
	if role.MenusIDs != nil && len(role.MenusIDs) > 0 {
		rlen := len(role.MenusIDs)
		userRoles := make([]*RoleMenusDO, 0, rlen)
		for _, obj := range role.MenusIDs {
			userRoles = append(userRoles, &RoleMenusDO{
				RoleID:  role.ID,
				MenusID: obj,
			})
		}
		// 关联关系
		if len(userRoles) > 0 {
			// 关联
			err := uc.repo.UpdateInBatches(ctx, userRoles)
			if err != nil {
				return err
			}
		}
	} else { // 删除用户-角色关系
		// 关联
		err := uc.repo.DeleteByRoleId(ctx, role.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
func (uc *RoleMenusUsecase) FindMenusIdsByRoleId(ctx context.Context, roleId uint64) ([]uint64, error) {
	data, err := uc.repo.FindMenusIdsByRoleId(ctx, roleId)
	return data, err
}
