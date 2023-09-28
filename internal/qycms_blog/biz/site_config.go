package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"gorm.io/gorm"
)

var (
	// ErrSiteConfigNotFound is site config not found.
	ErrSiteConfigNotFound = errors.NotFound("117404", "site config not found")
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
	InitLocalSiteConfigCache(ctx context.Context)
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
			return nil, ErrSiteConfigNotFound
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
			return nil, ErrSiteConfigNotFound
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
			return ErrSiteConfigNotFound
		}
		return err
	}

	return nil
}

func (uc *SiteConfigUsecase) FindValueByKey(ctx context.Context, key string) string {
	uc.log.WithContext(ctx).Infof("findByKey: %v", key)
	val, ok := uc.repo.GetValueByKey(key)
	if ok {
		return val.(string)
	}
	uc.repo.InitLocalSiteConfigCache(ctx)
	val, ok = uc.repo.GetValueByKey(key)
	if ok {
		return val.(string)
	}
	return ""
}

func (uc *SiteConfigUsecase) ListAll(ctx context.Context, opts SiteConfigDOListOption) ([]*SiteConfigDO, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	dataDOs, err := uc.repo.ListAll(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrSiteConfigNotFound
		}
		return nil, err
	}

	return dataDOs, nil
}
