package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"gorm.io/gorm"
)

type MenusAdminDO struct {
	metaV1.ObjectMeta
	Name           string          // 展示名称
	BreadcrumbName string          // 标签页名称
	Identifier     string          // 路由名称
	ParentId       uint64          // 父菜单ID
	Icon           string          // Icon图标
	MType          int             // 路由类型
	Path           string          // 路由 path
	Redirect       string          // 路由重定向
	Component      string          // 对应前端文件路径
	Sort           int             // 排序标记
	Children       []*MenusAdminDO // 子集
	HasChildren    bool            // 子集
}

type MenusAdminDOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*MenusAdminDO `json:"items"`
}

type MenusAdminDOListOption struct {
	metaV1.ListOptions `json:"page"`
	MenusAdminDO       `json:"item"`
}

type MenusAdminRepo interface {
	Save(context.Context, *MenusAdminDO) (*MenusAdminDO, error)
	Update(context.Context, *MenusAdminDO) (*MenusAdminDO, error)
	Delete(context.Context, uint64) error
	DeleteList(c context.Context, uids []uint64) error
	FindByID(context.Context, uint64) (*MenusAdminDO, error)
	//FindByKey(c context.Context, key string) (*po.MenusAdminPO, error)
	ListAll(c context.Context, opts MenusAdminDOListOption) (*MenusAdminDOList, error)
}

type MenusAdminUsecase struct {
	repo MenusAdminRepo
	log  *log.Helper
}

func NewMenusAdminUsecase(repo MenusAdminRepo, logger log.Logger) *MenusAdminUsecase {
	return &MenusAdminUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *MenusAdminUsecase) Create(ctx context.Context, obj *MenusAdminDO) (*MenusAdminDO, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", obj.Name)
	objDO, err := uc.repo.Save(ctx, obj)
	if err != nil {
		return nil, err
	}
	return objDO, nil
}

// Update 更新用户
func (uc *MenusAdminUsecase) Update(ctx context.Context, obj *MenusAdminDO) (*MenusAdminDO, error) {
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

// Delete 根据ID删除用户
func (uc *MenusAdminUsecase) Delete(ctx context.Context, id uint64) error {
	uc.log.WithContext(ctx).Infof("Delete: %v", id)
	return uc.repo.Delete(ctx, id)
}

// DeleteList 根据ID批量删除用户
func (uc *MenusAdminUsecase) DeleteList(ctx context.Context, ids []uint64) error {
	uc.log.WithContext(ctx).Infof("DeleteList: %v", ids)
	return uc.repo.DeleteList(ctx, ids)
}

// FindOneByID 根据ID查询用户信息
func (uc *MenusAdminUsecase) FindOneByID(ctx context.Context, id uint64) (*MenusAdminDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByID: %v", id)
	objDO, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return objDO, nil
}

// FindByKey 根据用户名查询用户信息
//func (uc *MenusAdminUsecase) FindByKey(ctx context.Context, objname string) (*MenusAdminDO, error) {
//	uc.log.WithContext(ctx).Infof("FindByKey: %v", objname)
//	obj, err := uc.repo.FindByKey(ctx, objname)
//	if err != nil {
//		return nil, err
//	}
//	objDO := &MenusAdminDO{
//		ObjectMeta: obj.ObjectMeta,
//		Name:       obj.Name,
//		Identifier: obj.Identifier,
//	}
//	return objDO, nil
//}

// ListAll 批量查询
func (uc *MenusAdminUsecase) ListAll(ctx context.Context, opts MenusAdminDOListOption) (*MenusAdminDOList, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	objDOs, err := uc.repo.ListAll(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	if opts.HasChildren {
		// 获取子集
		newopts := MenusAdminDOListOption{}
		if opts.MType > 0 {
			newopts.MType = opts.MType
		}
		newopts.PageFlag = false
		for _, obj := range objDOs.Items {
			newopts.ParentId = obj.ID
			cobjPOs, aerr := uc.ListAllChildren(ctx, newopts)
			obj.Children = cobjPOs.Items
			if aerr != nil {
				log.Errorf("菜单列表获取异常%v", aerr)
				obj.Children = make([]*MenusAdminDO, 0, 0)
			}

		}
	}

	return objDOs, nil
}

// ListAllParent 获取所有菜单列表
//func (uc *MenusAdminUsecase) ListAllParent(ctx context.Context, opts MenusAdminDOListOption) (*MenusAdminDOList, error) {
//	uc.log.WithContext(ctx).Infof("ListAll")
//	opts.ParentId = 0
//	objDOs, err := uc.repo.ListAll(ctx, opts)
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, ErrUserNotFound
//		}
//		return nil, err
//	}
//
//	infos := make([]*MenusAdminDO, 0, len(objDOs.Items))
//
//	var parent *MenusAdminDO
//
//	newopts := MenusAdminDOListOption{}
//	newopts.PageFlag = false
//	for _, obj := range objDOs.Items {
//		newopts.ParentId = obj.ID
//		cobjPOs, aerr := uc.ListAllChildren(ctx, newopts)
//		if aerr != nil {
//			log.Errorf("菜单列表获取异常%v", aerr)
//			cobjPOs = &MenusAdminDOList{}
//			cobjPOs.Items = make([]*MenusAdminDO, 0)
//		}
//		parent.Children = cobjPOs.Items
//		infos = append(infos, parent)
//	}
//	return &MenusAdminDOList{ListMeta: objPOs.ListMeta, Items: infos}, nil
//}

func (uc *MenusAdminUsecase) ListAllChildren(ctx context.Context, opts MenusAdminDOListOption) (*MenusAdminDOList, error) {
	uc.log.WithContext(ctx).Infof("ListAllChildren")
	objDOs, err := uc.repo.ListAll(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	newopts := MenusAdminDOListOption{}
	newopts.PageFlag = false
	if opts.MType > 0 {
		newopts.MType = opts.MType
	}
	for _, obj := range objDOs.Items {
		newopts.ParentId = obj.ID
		cobjPOs, aerr := uc.ListAllChildren(ctx, newopts)
		obj.Children = cobjPOs.Items
		if aerr != nil {
			log.Errorf("ListAllChildren 菜单列表获取异常%v", aerr)
			cobjPOs = &MenusAdminDOList{}
			obj.Children = make([]*MenusAdminDO, 0, 0)
		}
	}
	return objDOs, nil
}
