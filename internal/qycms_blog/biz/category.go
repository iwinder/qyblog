package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"

	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/utils/stringUtil"
	"gorm.io/gorm"
)

var (
	// ErrCategoryNotFound is category not found.
	ErrCategoryNotFound = errors.NotFound("105404", "category not found")
)

type CategoryDO struct {
	metaV1.ObjectMeta
	Name        string
	Identifier  string
	Description string
	ParentId    uint64
	Children    []*CategoryDO // 子集
}

type CategoryDOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*CategoryDO `json:"items"`
}

type CategoryDOListOption struct {
	metaV1.ListOptions `json:"page"`
	CategoryDO         `json:"item"`
}

type CategoryRepo interface {
	Save(context.Context, *CategoryDO) (*CategoryDO, error)
	Update(context.Context, *CategoryDO) (*CategoryDO, error)
	//Delete(context.Context, uint64) error
	DeleteList(c context.Context, uids []uint64) error
	FindByID(context.Context, uint64) (*CategoryDO, error)
	CountByIdentifier(ctx context.Context, str string) (int64, error)
	ListAll(c context.Context, opts CategoryDOListOption) (*CategoryDOList, error)
	ListAllWithChildren(c context.Context, opts CategoryDOListOption) (*CategoryDOList, error)
	FindByIdentifier(ctx context.Context, name string) (*CategoryDO, error)
}

type CategoryUsecase struct {
	repo CategoryRepo
	log  *log.Helper
}

func NewCategoryUsecase(repo CategoryRepo, logger log.Logger) *CategoryUsecase {
	return &CategoryUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CategoryUsecase) Create(ctx context.Context, obj *CategoryDO) (*CategoryDO, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", obj.Name)
	if obj.Identifier == "" || len(obj.Identifier) == 0 {
		obj.Identifier = stringUtil.PinyinConvert(obj.Name)
	}
	count, _ := uc.repo.CountByIdentifier(ctx, obj.Identifier)
	if count > 0 {
		obj.Identifier = fmt.Sprintf("%s-%d", obj.Identifier, count)
	}
	objDO, err := uc.repo.Save(ctx, obj)
	if err != nil {
		return nil, err
	}
	return objDO, nil
}

// Update 更新
func (uc *CategoryUsecase) Update(ctx context.Context, obj *CategoryDO) (*CategoryDO, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", obj.Name)
	objDO, err := uc.repo.Update(ctx, obj)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCategoryNotFound
		}
		return nil, err
	}
	return objDO, nil
}

// DeleteList 根据ID批量删除用户
func (uc *CategoryUsecase) DeleteList(ctx context.Context, ids []uint64) error {
	uc.log.WithContext(ctx).Infof("DeleteList: %v", ids)
	return uc.repo.DeleteList(ctx, ids)
}

// FindOneByID 根据ID查询用户信息
func (uc *CategoryUsecase) FindOneByID(ctx context.Context, id uint64) (*CategoryDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByID: %v", id)
	obj, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCategoryNotFound
		}
		return nil, err
	}
	return obj, nil
}
func (uc *CategoryUsecase) FindByIdentifier(ctx context.Context, name string) (*CategoryDO, error) {
	uc.log.WithContext(ctx).Infof("FindByIdentifier: %v", name)
	obj, err := uc.repo.FindByIdentifier(ctx, name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCategoryNotFound
		}
		return nil, err
	}
	return obj, nil
}

// ListAll 批量查询
func (uc *CategoryUsecase) ListAll(ctx context.Context, opts CategoryDOListOption) (*CategoryDOList, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	objDOs, err := uc.repo.ListAllWithChildren(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCategoryNotFound
		}
		return nil, err
	}
	return objDOs, nil
}
