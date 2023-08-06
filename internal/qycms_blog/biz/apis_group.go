package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"gorm.io/gorm"
)

var (
	// ErrApiGroupNotFound is api group not found.
	ErrApiGroupNotFound = errors.NotFound("102404", "api group not found")
)

type ApiGroupDO struct {
	metaV1.ObjectMeta
	Name       string
	Identifier string
}

type ApiGroupDOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*ApiGroupDO `json:"items"`
}

type ApiGroupDOListOption struct {
	metaV1.ListOptions `json:"page"`
	ApiGroupDO         `json:"item"`
}

type ApiGroupRepo interface {
	Save(context.Context, *ApiGroupDO) (*ApiGroupDO, error)
	Update(context.Context, *ApiGroupDO) (*ApiGroupDO, error)
	//Delete(context.Context, uint64) error
	DeleteList(c context.Context, uids []uint64) error
	FindByID(context.Context, uint64) (*ApiGroupDO, error)
	ListAll(c context.Context, opts ApiGroupDOListOption) (*ApiGroupDOList, error)
}

type ApiGroupUsecase struct {
	repo ApiGroupRepo
	log  *log.Helper
}

func NewApiGroupUsecase(repo ApiGroupRepo, logger log.Logger) *ApiGroupUsecase {
	return &ApiGroupUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ApiGroupUsecase) Create(ctx context.Context, obj *ApiGroupDO) (*ApiGroupDO, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", obj.Name)
	objDO, err := uc.repo.Save(ctx, obj)
	if err != nil {
		return nil, err
	}
	return objDO, nil
}

// Update 更新用户
func (uc *ApiGroupUsecase) Update(ctx context.Context, obj *ApiGroupDO) (*ApiGroupDO, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", obj.Name)
	objDO, err := uc.repo.Update(ctx, obj)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrApiGroupNotFound
		}
		return nil, err
	}
	return objDO, nil
}

// Delete 根据ID删除用户
//func (uc *ApiGroupUsecase) Delete(ctx context.Context, id uint64) error {
//	uc.log.WithContext(ctx).Infof("Delete: %v", id)
//	return uc.repo.Delete(ctx, id)
//}

// DeleteList 根据ID批量删除用户
func (uc *ApiGroupUsecase) DeleteList(ctx context.Context, ids []uint64) error {
	uc.log.WithContext(ctx).Infof("DeleteList: %v", ids)
	return uc.repo.DeleteList(ctx, ids)
}

// FindOneByID 根据ID查询用户信息
func (uc *ApiGroupUsecase) FindOneByID(ctx context.Context, id uint64) (*ApiGroupDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByID: %v", id)
	obj, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrApiGroupNotFound
		}
		return nil, err
	}
	return obj, nil
}

// ListAll 批量查询
func (uc *ApiGroupUsecase) ListAll(ctx context.Context, opts ApiGroupDOListOption) (*ApiGroupDOList, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	objDOs, err := uc.repo.ListAll(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrApiGroupNotFound
		}
		return nil, err
	}

	return objDOs, nil
}
