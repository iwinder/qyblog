package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"gorm.io/gorm"
)

var (
	// ErrMenusAgentNotFound is menus agent not found.
	ErrMenusAgentNotFound = errors.NotFound("114404", "menus agent not found")
)

type MenusAgentDO struct {
	metaV1.ObjectMeta
	Name  string
	Ftype string
}

type MenusAgentDOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*MenusAgentDO `json:"items"`
}

type MenusAgentDOListOption struct {
	metaV1.ListOptions `json:"page"`
	MenusAgentDO       `json:"item"`
}

type MenusAgentRepo interface {
	Save(context.Context, *MenusAgentDO) (*MenusAgentDO, error)
	Update(context.Context, *MenusAgentDO) (*MenusAgentDO, error)
	Delete(context.Context, uint64) error
	DeleteList(c context.Context, uids []uint64) error
	FindByID(context.Context, uint64) (*MenusAgentDO, error)
	ListAll(c context.Context, opts MenusAgentDOListOption) (*MenusAgentDOList, error)
}

type MenusAgentUsecase struct {
	repo MenusAgentRepo
	log  *log.Helper
}

func NewMenusAgentUsecase(repo MenusAgentRepo, logger log.Logger) *MenusAgentUsecase {
	return &MenusAgentUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *MenusAgentUsecase) Create(ctx context.Context, obj *MenusAgentDO) (*MenusAgentDO, error) {
	uc.log.WithContext(ctx).Infof("CreateData: %v", obj.Name)
	objDO, err := uc.repo.Save(ctx, obj)
	if err != nil {
		return nil, err
	}

	return objDO, nil
}

// Update 更新
func (uc *MenusAgentUsecase) Update(ctx context.Context, obj *MenusAgentDO) (*MenusAgentDO, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", obj.Name)
	objDO, err := uc.repo.Update(ctx, obj)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMenusAgentNotFound
		}
		return nil, err
	}
	return objDO, nil
}

// Delete 根据ID删除
func (uc *MenusAgentUsecase) Delete(ctx context.Context, id uint64) error {
	uc.log.WithContext(ctx).Infof("Delete: %v", id)
	return uc.repo.Delete(ctx, id)
}

// DeleteList 根据ID批量删除
func (uc *MenusAgentUsecase) DeleteList(ctx context.Context, ids []uint64) error {
	uc.log.WithContext(ctx).Infof("DeleteList: %v", ids)
	return uc.repo.DeleteList(ctx, ids)
}

// FindOneByID 根据ID查询信息
func (uc *MenusAgentUsecase) FindOneByID(ctx context.Context, id uint64) (*MenusAgentDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByID: %v", id)
	obj, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMenusAgentNotFound
		}
		return nil, err
	}

	return obj, nil
}

// ListAll 批量查询
func (uc *MenusAgentUsecase) ListAll(ctx context.Context, opts MenusAgentDOListOption) (*MenusAgentDOList, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	objDOs, err := uc.repo.ListAll(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMenusAgentNotFound
		}
		return nil, err
	}
	return objDOs, nil
}
