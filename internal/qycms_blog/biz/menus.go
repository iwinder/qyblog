package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"gorm.io/gorm"
)

type MenusDO struct {
	metaV1.ObjectMeta
	Name        string
	Url         string
	Blanked     int32
	ParentId    uint64
	TargetId    uint64
	Children    []*MenusDO
	HasChildren bool
}

type MenusDOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*MenusDO `json:"items"`
}

type MenusDOListOption struct {
	metaV1.ListOptions `json:"page"`
	MenusDO            `json:"item"`
}

type MenusRepo interface {
	Save(context.Context, *MenusDO) (*MenusDO, error)
	Update(context.Context, *MenusDO) (*MenusDO, error)
	Delete(context.Context, uint64) error
	DeleteList(context.Context, []uint64, uint64) error
	FindByID(context.Context, uint64) (*MenusDO, error)
	ListAll(c context.Context, opts MenusDOListOption) (*MenusDOList, error)
	ListAllWithChildren(c context.Context, opts MenusDOListOption) (*MenusDOList, error)
}

type MenusUsecase struct {
	repo MenusRepo
	log  *log.Helper
}

func NewMenusUsecase(repo MenusRepo, logger log.Logger) *MenusUsecase {
	return &MenusUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *MenusUsecase) Create(ctx context.Context, obj *MenusDO) (*MenusDO, error) {
	uc.log.WithContext(ctx).Infof("CreateData: %v", obj.Name)
	objDO, err := uc.repo.Save(ctx, obj)
	if err != nil {
		return nil, err
	}

	return objDO, nil
}

// Update 更新
func (uc *MenusUsecase) Update(ctx context.Context, obj *MenusDO) (*MenusDO, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", obj.Name)
	objDO, err := uc.repo.Update(ctx, obj)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return objDO, nil
}

// Delete 根据ID删除
func (uc *MenusUsecase) Delete(ctx context.Context, id uint64) error {
	uc.log.WithContext(ctx).Infof("Delete: %v", id)
	return uc.repo.Delete(ctx, id)
}

// DeleteList 根据ID批量删除
func (uc *MenusUsecase) DeleteList(ctx context.Context, ids []uint64, targetId uint64) error {
	uc.log.WithContext(ctx).Infof("DeleteList: %v", ids)
	return uc.repo.DeleteList(ctx, ids, targetId)
}

// FindOneByID 根据ID查询信息
func (uc *MenusUsecase) FindOneByID(ctx context.Context, id uint64) (*MenusDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByID: %v", id)
	obj, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return obj, nil
}

// ListAll 批量查询
func (uc *MenusUsecase) ListAll(ctx context.Context, opts MenusDOListOption) (*MenusDOList, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	objDOs, err := uc.repo.ListAllWithChildren(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return objDOs, nil
}
