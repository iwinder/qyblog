package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"gorm.io/gorm"
)

type SiteConfigDO struct {
	metaV1.ObjectMeta
	ConfigKey   string
	ConfigValue string
	ConfigName  string
	ConfigTip   string
	Ftype       int
}
type SiteConfigDOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*SiteConfigDO `json:"items"`
}

type SiteConfigDOListOption struct {
	metaV1.ListOptions `json:"page"`
	SiteConfigDO       `json:"item"`
	Types              string
}

type SiteConfigRepo interface {
	Save(ctx context.Context, data *SiteConfigDO) (*SiteConfigDO, error)
	Update(context.Context, *SiteConfigDO) (*SiteConfigDO, error)
	UpdateInBatches(context.Context, []*SiteConfigDO) error
	GetValueByKey(key string) (any, bool)
	ListAll(c context.Context, opts SiteConfigDOListOption) ([]*SiteConfigDO, error)
}

type SiteConfigUsecase struct {
	repo SiteConfigRepo
	log  *log.Helper
}

// NewSiteConfigUsecase new a UserDO usecase.
func NewSiteConfigUsecase(repo SiteConfigRepo, logger log.Logger) *SiteConfigUsecase {
	return &SiteConfigUsecase{repo: repo, log: log.NewHelper(logger)}
}
func (uc *SiteConfigUsecase) Save(ctx context.Context, data *SiteConfigDO) (*SiteConfigDO, error) {
	uc.log.WithContext(ctx).Infof("Save: %v", data.ConfigName)
	dataDO, err := uc.repo.Save(ctx, data)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return dataDO, nil
}
func (uc *SiteConfigUsecase) Update(ctx context.Context, data *SiteConfigDO) (*SiteConfigDO, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", data.ConfigName)
	dataDO, err := uc.repo.Update(ctx, data)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return dataDO, nil
}
func (uc *SiteConfigUsecase) UpdateInBatches(ctx context.Context, data []*SiteConfigDO) error {
	uc.log.WithContext(ctx).Infof("UpdateInBatches ")
	err := uc.repo.UpdateInBatches(ctx, data)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotFound
		}
		return err
	}

	return nil
}

//func (uc *SiteConfigUsecase) DeleteList(ctx context.Context, ids []uint64) error {
//	uc.log.WithContext(ctx).Infof("DeleteList: %v", ids)
//	return uc.repo.DeleteList(ctx, ids)
//}
//func (uc *SiteConfigUsecase) FindByMd5(ctx context.Context, fmd5 string) (*SiteConfigDO, error) {
//	uc.log.WithContext(ctx).Infof("FindByMd5: %v", fmd5)
//	return uc.repo.FindByMd5(ctx, fmd5)
//}
//func (uc *SiteConfigUsecase) CountByMd5(ctx context.Context, fmd5 string) int64 {
//	uc.log.WithContext(ctx).Infof("CountByMd5: %v", fmd5)
//	count, _ := uc.repo.CountByMd5(ctx, fmd5)
//	return count
//}

func (uc *SiteConfigUsecase) ListAll(ctx context.Context, opts SiteConfigDOListOption) ([]*SiteConfigDO, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	dataDOs, err := uc.repo.ListAll(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return dataDOs, nil
}
