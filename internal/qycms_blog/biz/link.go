package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"gorm.io/gorm"
)

var (
	// ErrLinkNotFound is link not found.
	ErrLinkNotFound = errors.NotFound("112404", "link not found")
)

type LinkDO struct {
	metaV1.ObjectMeta
	Name        string
	Url         string
	Description string
	Ftype       int32
}

type LinkDOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*LinkDO `json:"items"`
}

type LinkDOListOption struct {
	metaV1.ListOptions `json:"page"`
	LinkDO             `json:"item"`
}

type LinkRepo interface {
	Save(context.Context, *LinkDO) (*LinkDO, error)
	Update(context.Context, *LinkDO) (*LinkDO, error)
	Delete(context.Context, uint64) error
	DeleteList(c context.Context, uids []uint64) error
	FindByID(context.Context, uint64) (*LinkDO, error)
	ListAll(c context.Context, opts LinkDOListOption) (*LinkDOList, error)
	FindIndexLinkAllWitchCache(context.Context) ([]*LinkDO, error)
	FindAllWitchCache(context.Context) ([]*LinkDO, error)
}

type LinkUsecase struct {
	repo LinkRepo
	log  *log.Helper
}

func NewLinkUsecase(repo LinkRepo, logger log.Logger) *LinkUsecase {
	return &LinkUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *LinkUsecase) Create(ctx context.Context, obj *LinkDO) (*LinkDO, error) {
	uc.log.WithContext(ctx).Infof("CreateData: %v", obj.Name)
	objDO, err := uc.repo.Save(ctx, obj)
	if err != nil {
		return nil, err
	}

	return objDO, nil
}

// Update 更新
func (uc *LinkUsecase) Update(ctx context.Context, obj *LinkDO) (*LinkDO, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", obj.Name)
	objDO, err := uc.repo.Update(ctx, obj)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrLinkNotFound
		}
		return nil, err
	}
	return objDO, nil
}

// Delete 根据ID删除
func (uc *LinkUsecase) Delete(ctx context.Context, id uint64) error {
	uc.log.WithContext(ctx).Infof("Delete: %v", id)
	return uc.repo.Delete(ctx, id)
}

// DeleteList 根据ID批量删除
func (uc *LinkUsecase) DeleteList(ctx context.Context, ids []uint64) error {
	uc.log.WithContext(ctx).Infof("DeleteList: %v", ids)
	return uc.repo.DeleteList(ctx, ids)
}

// FindOneByID 根据ID查询信息
func (uc *LinkUsecase) FindOneByID(ctx context.Context, id uint64) (*LinkDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByID: %v", id)
	obj, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrLinkNotFound
		}
		return nil, err
	}

	return obj, nil
}

func (uc *LinkUsecase) FindAllWitchCache(ctx context.Context) ([]*LinkDO, error) {
	data, err := uc.repo.FindAllWitchCache(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (uc *LinkUsecase) FindIndexLinkAllWitchCache(ctx context.Context) ([]*LinkDO, error) {
	data, err := uc.repo.FindIndexLinkAllWitchCache(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ListAll 批量查询
func (uc *LinkUsecase) ListAll(ctx context.Context, opts LinkDOListOption) (*LinkDOList, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	objDOs, err := uc.repo.ListAll(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrLinkNotFound
		}
		return nil, err
	}
	return objDOs, nil
}
