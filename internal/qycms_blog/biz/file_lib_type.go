package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"gorm.io/gorm"
)

var (
	// ErrFileLibTypeNotFound is file lib type not found.
	ErrFileLibTypeNotFound = errors.NotFound("111404", "file lib type not found")
)

type FileLibTypeDO struct {
	metaV1.ObjectMeta
	Name       string
	Identifier int
	Ftype      string
}
type FileLibTypeDOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*FileLibTypeDO `json:"items"`
}

type FileLibTypeDOListOption struct {
	metaV1.ListOptions `json:"page"`
	FileLibTypeDO      `json:"item"`
}

type FileLibTypeRepo interface {
	Save(ctx context.Context, data *FileLibTypeDO) (*FileLibTypeDO, error)
	Update(context.Context, *FileLibTypeDO) (*FileLibTypeDO, error)
	DeleteList(c context.Context, uids []uint64) error
	ListAll(c context.Context, opts FileLibTypeDOListOption) (*FileLibTypeDOList, error)
}

type FileLibTypeUsecase struct {
	repo FileLibTypeRepo
	log  *log.Helper
}

// NewFileLibTypeUsecase new a UserDO usecase.
func NewFileLibTypeUsecase(repo FileLibTypeRepo, logger log.Logger) *FileLibTypeUsecase {
	return &FileLibTypeUsecase{repo: repo, log: log.NewHelper(logger)}
}
func (uc *FileLibTypeUsecase) Save(ctx context.Context, data *FileLibTypeDO) (*FileLibTypeDO, error) {
	uc.log.WithContext(ctx).Infof("Save: %v", data.Name)
	dataDO, err := uc.repo.Save(ctx, data)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrFileLibTypeNotFound
		}
		return nil, err
	}

	return dataDO, nil
}
func (uc *FileLibTypeUsecase) Update(ctx context.Context, data *FileLibTypeDO) (*FileLibTypeDO, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", data.Name)
	dataDO, err := uc.repo.Update(ctx, data)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrFileLibTypeNotFound
		}
		return nil, err
	}

	return dataDO, nil
}

func (uc *FileLibTypeUsecase) DeleteList(ctx context.Context, ids []uint64) error {
	uc.log.WithContext(ctx).Infof("DeleteList: %v", ids)
	return uc.repo.DeleteList(ctx, ids)
}

func (uc *FileLibTypeUsecase) ListAll(ctx context.Context, opts FileLibTypeDOListOption) (*FileLibTypeDOList, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	dataDOs, err := uc.repo.ListAll(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrFileLibTypeNotFound
		}
		return nil, err
	}

	return dataDOs, nil
}
