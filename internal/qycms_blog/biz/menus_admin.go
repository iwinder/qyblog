package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
	"gorm.io/gorm"
)

type MenusAdminDO struct {
	metaV1.ObjectMeta
	Level     uint
	ParentId  uint64          // 父菜单ID
	Path      string          // 路由path
	Name      string          // 路由name
	Hidden    bool            // 是否在列表隐藏
	Component string          // 对应前端文件路径
	Sort      int             // 排序标记
	Children  []*MenusAdminDO // 子集
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
	Save(context.Context, *MenusAdminDO) (*po.MenusAdminPO, error)
	Update(context.Context, *MenusAdminDO) (*po.MenusAdminPO, error)
	Delete(context.Context, uint64) error
	DeleteList(c context.Context, uids []uint64) error
	FindByID(context.Context, uint64) (*po.MenusAdminPO, error)
	//FindByKey(c context.Context, key string) (*po.MenusAdminPO, error)
	ListAll(c context.Context, opts MenusAdminDOListOption) (*po.MenusAdminPOList, error)
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
	objPO, err := uc.repo.Save(ctx, obj)
	if err != nil {
		return nil, err
	}
	objDO := &MenusAdminDO{Name: objPO.Name}
	objDO.ID = objPO.ID
	return objDO, nil
}

// Update 更新用户
func (uc *MenusAdminUsecase) Update(ctx context.Context, obj *MenusAdminDO) (*MenusAdminDO, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", obj.Name)
	objPO, err := uc.repo.Update(ctx, obj)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	objDO := &MenusAdminDO{Name: objPO.Name}
	objDO.ID = objPO.ID
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
	obj, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	objDO := &MenusAdminDO{
		ObjectMeta: obj.ObjectMeta,
		Level:      obj.Level,
		ParentId:   obj.ParentId,
		Path:       obj.Path,
		Name:       obj.Name,
		Hidden:     obj.Hidden,
		Component:  obj.Component,
		Sort:       obj.Sort,
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
	objPOs, err := uc.repo.ListAll(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	infos := make([]*MenusAdminDO, 0, len(objPOs.Items))
	for _, obj := range objPOs.Items {
		infos = append(infos, &MenusAdminDO{
			ObjectMeta: metaV1.ObjectMeta{
				ID:         obj.ID,
				InstanceID: obj.InstanceID,
				Extend:     obj.Extend,
				CreatedAt:  obj.CreatedAt,
				UpdatedAt:  obj.UpdatedAt,
			},
			Level:     obj.Level,
			ParentId:  obj.ParentId,
			Path:      obj.Path,
			Name:      obj.Name,
			Hidden:    obj.Hidden,
			Component: obj.Component,
			Sort:      obj.Sort,
		})
	}
	return &MenusAdminDOList{ListMeta: objPOs.ListMeta, Items: infos}, nil
}

// ListAllParent 获取所有菜单列表
func (uc *MenusAdminUsecase) ListAllParent(ctx context.Context, opts MenusAdminDOListOption) (*MenusAdminDOList, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	opts.ParentId = 0
	objPOs, err := uc.repo.ListAll(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	infos := make([]*MenusAdminDO, 0, len(objPOs.Items))

	var parent *MenusAdminDO
	newopts := MenusAdminDOListOption{}
	newopts.PageFlag = false
	for _, obj := range objPOs.Items {
		parent = &MenusAdminDO{
			ObjectMeta: metaV1.ObjectMeta{
				ID:         obj.ID,
				InstanceID: obj.InstanceID,
				Extend:     obj.Extend,
				CreatedAt:  obj.CreatedAt,
				UpdatedAt:  obj.UpdatedAt,
			},
			Level:     obj.Level,
			ParentId:  obj.ParentId,
			Path:      obj.Path,
			Name:      obj.Name,
			Hidden:    obj.Hidden,
			Component: obj.Component,
			Sort:      obj.Sort,
		}
		newopts.ParentId = parent.ID
		cobjPOs, aerr := uc.ListAllChildren(ctx, newopts)
		if aerr != nil {
			log.Errorf("菜单列表获取异常%v", aerr)
			cobjPOs = &MenusAdminDOList{}
			cobjPOs.Items = make([]*MenusAdminDO, 0)
		}
		parent.Children = cobjPOs.Items
		infos = append(infos, parent)
	}
	return &MenusAdminDOList{ListMeta: objPOs.ListMeta, Items: infos}, nil
}

func (uc *MenusAdminUsecase) ListAllChildren(ctx context.Context, opts MenusAdminDOListOption) (*MenusAdminDOList, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	objPOs, err := uc.repo.ListAll(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	infos := make([]*MenusAdminDO, 0, len(objPOs.Items))
	newopts := MenusAdminDOListOption{}
	newopts.PageFlag = false
	var parent *MenusAdminDO
	for _, obj := range objPOs.Items {
		parent = &MenusAdminDO{
			ObjectMeta: metaV1.ObjectMeta{
				ID:         obj.ID,
				InstanceID: obj.InstanceID,
				Extend:     obj.Extend,
				CreatedAt:  obj.CreatedAt,
				UpdatedAt:  obj.UpdatedAt,
			},
			Level:     obj.Level,
			ParentId:  obj.ParentId,
			Path:      obj.Path,
			Name:      obj.Name,
			Hidden:    obj.Hidden,
			Component: obj.Component,
			Sort:      obj.Sort,
		}
		cobjPOs, aerr := uc.ListAllChildren(ctx, newopts)
		if aerr != nil {
			log.Errorf("菜单列表获取异常%v", aerr)
			cobjPOs = &MenusAdminDOList{}
			cobjPOs.Items = make([]*MenusAdminDO, 0)
		}
		parent.Children = cobjPOs.Items
		infos = append(infos, parent)
	}
	return &MenusAdminDOList{ListMeta: objPOs.ListMeta, Items: infos}, nil
}
