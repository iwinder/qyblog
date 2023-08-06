package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"gorm.io/gorm"
)

var (
	// ErrShortLinkNotFound is short link not found.
	ErrShortLinkNotFound = errors.NotFound("116404", "short link not found")
)

type ShortLinkDO struct {
	metaV1.ObjectMeta
	Url         string
	Description string
	Identifier  string
}

type ShortLinkDOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*ShortLinkDO `json:"items"`
}

type ShortLinkDOListOption struct {
	metaV1.ListOptions `json:"page"`
	ShortLinkDO        `json:"item"`
}

type ShortLinkRepo interface {
	Save(context.Context, *ShortLinkDO) (*ShortLinkDO, error)
	Update(context.Context, *ShortLinkDO) (*ShortLinkDO, error)
	Delete(context.Context, uint64) error
	DeleteList(c context.Context, uids []uint64) error
	FindByID(context.Context, uint64) (*ShortLinkDO, error)
	ListAll(c context.Context, opts ShortLinkDOListOption) (*ShortLinkDOList, error)
	FindAllWitchCache(context.Context) ([]*ShortLinkDO, error)
}

type ShortLinkUsecase struct {
	repo ShortLinkRepo
	log  *log.Helper
}

func NewShortLinkUsecase(repo ShortLinkRepo, logger log.Logger) *ShortLinkUsecase {
	return &ShortLinkUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ShortLinkUsecase) Create(ctx context.Context, obj *ShortLinkDO) (*ShortLinkDO, error) {
	uc.log.WithContext(ctx).Infof("CreateData: %v", obj.Identifier)
	objDO, err := uc.repo.Save(ctx, obj)
	if err != nil {
		return nil, err
	}

	return objDO, nil
}

// Update 更新
func (uc *ShortLinkUsecase) Update(ctx context.Context, obj *ShortLinkDO) (*ShortLinkDO, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", obj.Identifier)
	objDO, err := uc.repo.Update(ctx, obj)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrShortLinkNotFound
		}
		return nil, err
	}
	return objDO, nil
}

// Delete 根据ID删除
func (uc *ShortLinkUsecase) Delete(ctx context.Context, id uint64) error {
	uc.log.WithContext(ctx).Infof("Delete: %v", id)
	return uc.repo.Delete(ctx, id)
}

// DeleteList 根据ID批量删除
func (uc *ShortLinkUsecase) DeleteList(ctx context.Context, ids []uint64) error {
	uc.log.WithContext(ctx).Infof("DeleteList: %v", ids)
	return uc.repo.DeleteList(ctx, ids)
}

// FindOneByID 根据ID查询信息
func (uc *ShortLinkUsecase) FindOneByID(ctx context.Context, id uint64) (*ShortLinkDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByID: %v", id)
	obj, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrShortLinkNotFound
		}
		return nil, err
	}

	return obj, nil
}

// ListAll 批量查询
func (uc *ShortLinkUsecase) ListAll(ctx context.Context, opts ShortLinkDOListOption) (*ShortLinkDOList, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	objDOs, err := uc.repo.ListAll(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrShortLinkNotFound
		}
		return nil, err
	}
	return objDOs, nil
}
func (uc *ShortLinkUsecase) FindAllWitchCache(ctx context.Context) ([]*ShortLinkDO, error) {
	uc.log.WithContext(ctx).Infof("FindAllWitchCache")
	objDOs, err := uc.repo.FindAllWitchCache(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrShortLinkNotFound
		}
		return nil, err
	}
	return objDOs, nil
}
