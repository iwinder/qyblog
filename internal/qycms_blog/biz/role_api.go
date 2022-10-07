package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/auth/auth_constants"
	"strconv"
)

type RoleApiDO struct {
	RoleID uint64
	ApiID  uint64
}

type RoleApiRepo interface {
	CreateInBatches(ctx context.Context, userRoles []*RoleApiDO) error
	UpdateInBatches(ctx context.Context, userRoles []*RoleApiDO) error
	DeleteByRoleId(ctx context.Context, roleId uint64) error
	FindApiIdsByRoleId(ctx context.Context, roleId uint64) ([]uint64, error)
}

type RoleApiUsecase struct {
	repo      RoleApiRepo
	cabinRepo CasbinRuleRepo
	log       *log.Helper
}

// NewRoleApiUsecase new a UserDO usecase.
func NewRoleApiUsecase(repo RoleApiRepo, cabinRepo CasbinRuleRepo, logger log.Logger) *RoleApiUsecase {
	return &RoleApiUsecase{repo: repo, cabinRepo: cabinRepo, log: log.NewHelper(logger)}
}

func (uc *RoleApiUsecase) SaveApisForRole(ctx context.Context, role *RoleDO) error {

	if role.Apis != nil && len(role.Apis) > 0 {
		roleKey := auth_constants.PrefixRole + strconv.FormatUint(role.ID, 10)
		rlen := len(role.Apis)
		roleApis := make([]*RoleApiDO, 0, rlen)
		rules := [][]string{}
		for _, aobj := range role.Apis {
			rules = append(rules, []string{roleKey, aobj.Identifier, aobj.Path, aobj.Method})
			roleApis = append(roleApis, &RoleApiDO{
				RoleID: role.ID,
				ApiID:  aobj.ID,
			})
		}

		// 关联关系
		if len(roleApis) > 0 {
			// 关联
			err := uc.repo.CreateInBatches(ctx, roleApis)
			if err != nil {
				return err
			}
			// 权限
			uc.cabinRepo.CleanPolicy(ctx, roleKey)
			_, cerr := uc.cabinRepo.SavePolicies(ctx, rules)
			if cerr != nil {
				return cerr
			}
		}

	}
	return nil
}

func (uc *RoleApiUsecase) UpdateApisForRole(ctx context.Context, role *RoleDO) error {
	roleKey := auth_constants.PrefixRole + strconv.FormatUint(role.ID, 10)
	if role.Apis != nil && len(role.Apis) > 0 {
		rlen := len(role.Apis)
		roleApis := make([]*RoleApiDO, 0, rlen)
		rules := [][]string{}
		for _, aobj := range role.Apis {
			rules = append(rules, []string{roleKey, aobj.Identifier, aobj.Path, aobj.Method})
			roleApis = append(roleApis, &RoleApiDO{
				RoleID: role.ID,
				ApiID:  aobj.ID,
			})
		}
		// 关联关系
		if len(roleApis) > 0 {
			// 关联
			err := uc.repo.UpdateInBatches(ctx, roleApis)
			if err != nil {
				return err
			}

			// 权限
			uc.cabinRepo.CleanPolicy(ctx, roleKey)
			_, cerr := uc.cabinRepo.SavePolicies(ctx, rules)
			if cerr != nil {
				return cerr
			}
		}

	} else { // 删除用户-角色关系
		// 关联
		err := uc.repo.DeleteByRoleId(ctx, role.ID)
		if err != nil {
			return err
		}
		_, ucerr := uc.cabinRepo.CleanPolicy(ctx, roleKey)
		if ucerr != nil {
			return ucerr
		}
	}
	return nil
}
func (uc *RoleApiUsecase) FindApiIdsByRoleId(ctx context.Context, roleId uint64) ([]uint64, error) {
	data, err := uc.repo.FindApiIdsByRoleId(ctx, roleId)
	return data, err
}
