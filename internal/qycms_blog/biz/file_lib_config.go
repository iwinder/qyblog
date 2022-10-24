package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"gorm.io/gorm"
)

type FileLibConfigDO struct {
	metaV1.ObjectMeta
	AccessKey string
	SecretKey string
	Bucket    string
	Prefix    string
	Domain    string
	Endpoint  string
	TypeId    uint64
}

type FileLibConfigRepo interface {
	Save(ctx context.Context, data *FileLibConfigDO) (*FileLibConfigDO, error)
	Update(context.Context, *FileLibConfigDO) (*FileLibConfigDO, error)
	FindByID(context.Context, uint64) (*FileLibConfigDO, error)
	FindByTypeId(context.Context, uint64) (*FileLibConfigDO, error)
}

type FileLibConfigUsecase struct {
	repo FileLibConfigRepo
	log  *log.Helper
}

func NewFileLibConfigUsecase(repo FileLibConfigRepo, logger log.Logger) *FileLibConfigUsecase {
	return &FileLibConfigUsecase{repo: repo, log: log.NewHelper(logger)}
}
func (uc *FileLibConfigUsecase) SaveOrUpdate(ctx context.Context, data *FileLibConfigDO) (*FileLibConfigDO, error) {
	uc.log.WithContext(ctx).Infof("Save: %v", data.TypeId)
	var dataDO *FileLibConfigDO
	var err error
	if data.ID > 0 {
		dataDO, err = uc.repo.Update(ctx, data)
	} else {
		dataDO, err = uc.repo.Save(ctx, data)
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return dataDO, nil
}
func (uc *FileLibConfigUsecase) FindById(ctx context.Context, id uint64) (*FileLibConfigDO, error) {
	uc.log.WithContext(ctx).Infof("FindById: %v", id)
	data, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return data, nil
}
func (uc *FileLibConfigUsecase) FindByTypeId(ctx context.Context, id uint64) (*FileLibConfigDO, error) {
	uc.log.WithContext(ctx).Infof("FindByTypeId: %v", id)
	data, err := uc.repo.FindByTypeId(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return data, nil
}
