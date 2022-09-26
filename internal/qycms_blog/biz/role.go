package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/auth/auth_constants"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
	"gorm.io/gorm"
)

type RoleDO struct {
	metaV1.ObjectMeta
	Name        string
	Identifier  string
	MenusAdmins []*MenusAdminDO
	Apis        []*ApiDO
}

type RoleDOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*RoleDO `json:"items"`
}

type RoleDOListOption struct {
	metaV1.ListOptions `json:"page"`
	RoleDO             `json:"item"`
}

type RoleRepo interface {
	Save(context.Context, *RoleDO) (*RoleDO, error)
	Update(context.Context, *RoleDO) (*RoleDO, error)
	Delete(context.Context, uint64) error
	DeleteList(c context.Context, uids []uint64) error
	FindByID(context.Context, uint64) (*po.RolePO, error)
	FindByKey(c context.Context, key string) (*po.RolePO, error)
	ListAll(c context.Context, opts RoleDOListOption) (*po.RolePOList, error)
	FindByUserId(c context.Context, userId uint64) ([]*RoleDO, error)
}

type RoleUsecase struct {
	repo      RoleRepo
	cabinRepo CasbinRuleRepo
	log       *log.Helper
}

func NewRoleUsecase(repo RoleRepo, cabinRepo CasbinRuleRepo, logger log.Logger) *RoleUsecase {
	return &RoleUsecase{repo: repo, cabinRepo: cabinRepo, log: log.NewHelper(logger)}
}

func (uc *RoleUsecase) Create(ctx context.Context, obj *RoleDO) (*RoleDO, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", obj.Name)
	objDO, err := uc.repo.Save(ctx, obj)
	if err != nil {
		return nil, err
	}

	return objDO, nil
}

// Update 更新
func (uc *RoleUsecase) Update(ctx context.Context, obj *RoleDO) (*RoleDO, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", obj.Name)
	objDO, err := uc.repo.Update(ctx, obj)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	if obj.Apis != nil && len(obj.Apis) > 0 {
		rules := [][]string{}
		for _, aobj := range obj.Apis {
			rules = append(rules, []string{auth_constants.PrefixRole + obj.Identifier, aobj.Identifier, aobj.Path, aobj.Method})
		}
		uc.cabinRepo.CleanPolicy(ctx, auth_constants.PrefixRole+obj.Identifier)
		_, cerr := uc.cabinRepo.SavePolicies(ctx, rules)
		if cerr != nil {
			return nil, cerr
		}
	}
	return objDO, nil
}

// Delete 根据ID删除
func (uc *RoleUsecase) Delete(ctx context.Context, id uint64) error {
	uc.log.WithContext(ctx).Infof("Delete: %v", id)
	return uc.repo.Delete(ctx, id)
}

// DeleteList 根据ID批量删除
func (uc *RoleUsecase) DeleteList(ctx context.Context, ids []uint64) error {
	uc.log.WithContext(ctx).Infof("DeleteList: %v", ids)
	return uc.repo.DeleteList(ctx, ids)
}

// FindOneByID 根据ID查询信息
func (uc *RoleUsecase) FindOneByID(ctx context.Context, id uint64) (*RoleDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByID: %v", id)
	obj, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	objDO := &RoleDO{
		ObjectMeta: obj.ObjectMeta,
		Name:       obj.Name,
		Identifier: obj.Identifier,
	}
	return objDO, nil
}
func (uc *RoleUsecase) FindByUserId(ctx context.Context, userId uint64) ([]*RoleDO, error) {
	uc.log.WithContext(ctx).Infof("FindByUserId: %v", userId)
	obj, err := uc.repo.FindByUserId(ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return obj, nil
}

// FindByKey 根据用户名查询信息
func (uc *RoleUsecase) FindByKey(ctx context.Context, objname string) (*RoleDO, error) {
	uc.log.WithContext(ctx).Infof("FindByKey: %v", objname)
	obj, err := uc.repo.FindByKey(ctx, objname)
	if err != nil {
		return nil, err
	}
	objDO := &RoleDO{
		ObjectMeta: obj.ObjectMeta,
		Name:       obj.Name,
		Identifier: obj.Identifier,
	}
	return objDO, nil
}

// ListAll 批量查询
func (uc *RoleUsecase) ListAll(ctx context.Context, opts RoleDOListOption) (*RoleDOList, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	objPOs, err := uc.repo.ListAll(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	infos := make([]*RoleDO, 0, len(objPOs.Items))
	for _, obj := range objPOs.Items {
		infos = append(infos, &RoleDO{
			ObjectMeta: metaV1.ObjectMeta{
				ID:         obj.ID,
				InstanceID: obj.InstanceID,
				Extend:     obj.Extend,
				CreatedAt:  obj.CreatedAt,
				UpdatedAt:  obj.UpdatedAt,
			},
			Name:       obj.Name,
			Identifier: obj.Identifier,
		})
	}
	return &RoleDOList{ListMeta: objPOs.ListMeta, Items: infos}, nil
}
