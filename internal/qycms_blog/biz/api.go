package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
	"gorm.io/gorm"
)

type ApiDO struct {
	metaV1.ObjectMeta
	ApiGroup    string
	Method      string
	Path        string
	Description string
}

type ApiDOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*ApiDO `json:"items"`
}

type ApiDOListOption struct {
	metaV1.ListOptions `json:"page"`
	ApiDO              `json:"item"`
}

type ApiRepo interface {
	Save(context.Context, *ApiDO) (*po.ApiPO, error)
	Update(context.Context, *ApiDO) (*po.ApiPO, error)
	Delete(context.Context, uint64) error
	DeleteList(c context.Context, uids []uint64) error
	FindByID(context.Context, uint64) (*po.ApiPO, error)
	ListAll(c context.Context, opts ApiDOListOption) (*po.ApiPOList, error)
}

type ApiUsecase struct {
	repo ApiRepo
	log  *log.Helper
}

func NewApiUsecase(repo ApiRepo, logger log.Logger) *ApiUsecase {
	return &ApiUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ApiUsecase) Create(ctx context.Context, obj *ApiDO) (*ApiDO, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", obj.Path)
	objPO, err := uc.repo.Save(ctx, obj)
	if err != nil {
		return nil, err
	}
	objDO := &ApiDO{Path: objPO.Path}
	objDO.ID = objPO.ID
	return objDO, nil
}

// Update 更新用户
func (uc *ApiUsecase) Update(ctx context.Context, obj *ApiDO) (*ApiDO, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", obj.Path)
	objPO, err := uc.repo.Update(ctx, obj)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	objDO := &ApiDO{Path: objPO.Path}
	objDO.ID = objPO.ID
	return objDO, nil
}

// Delete 根据ID删除用户
func (uc *ApiUsecase) Delete(ctx context.Context, id uint64) error {
	uc.log.WithContext(ctx).Infof("Delete: %v", id)
	return uc.repo.Delete(ctx, id)
}

// DeleteList 根据ID批量删除用户
func (uc *ApiUsecase) DeleteList(ctx context.Context, ids []uint64) error {
	uc.log.WithContext(ctx).Infof("DeleteList: %v", ids)
	return uc.repo.DeleteList(ctx, ids)
}

// FindOneByID 根据ID查询用户信息
func (uc *ApiUsecase) FindOneByID(ctx context.Context, id uint64) (*ApiDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByID: %v", id)
	obj, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	objDO := &ApiDO{
		ObjectMeta:  obj.ObjectMeta,
		ApiGroup:    obj.ApiGroup,
		Method:      obj.Method,
		Path:        obj.Path,
		Description: obj.Description,
	}
	return objDO, nil
}

// ListAll 批量查询
func (uc *ApiUsecase) ListAll(ctx context.Context, opts ApiDOListOption) (*ApiDOList, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	objPOs, err := uc.repo.ListAll(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	infos := make([]*ApiDO, 0, len(objPOs.Items))
	for _, obj := range objPOs.Items {
		infos = append(infos, &ApiDO{
			ObjectMeta: metaV1.ObjectMeta{
				ID:         obj.ID,
				InstanceID: obj.InstanceID,
				Extend:     obj.Extend,
				CreatedAt:  obj.CreatedAt,
				UpdatedAt:  obj.UpdatedAt,
			},
			ApiGroup:    obj.ApiGroup,
			Method:      obj.Method,
			Path:        obj.Path,
			Description: obj.Description,
		})
	}
	return &ApiDOList{ListMeta: objPOs.ListMeta, Items: infos}, nil
}
