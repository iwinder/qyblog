package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/gormutil"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qyblog/internal/qycms_blog/biz"
	"github.com/iwinder/qyblog/internal/qycms_blog/data/po"
	"gorm.io/gorm/clause"
)

// 角色管理
type roleRepo struct {
	data *Data
	log  *log.Helper
}

// NewRoleRepo .
func NewRoleRepo(data *Data, logger log.Logger) biz.RoleRepo {
	return &roleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Save 创建用户
func (r *roleRepo) Save(ctx context.Context, obj *biz.RoleDO) (*biz.RoleDO, error) {
	objPO := &po.RolePO{
		ObjectMeta: obj.ObjectMeta,
		Name:       obj.Name,
		Identifier: obj.Identifier,
	}
	err := r.data.Db.Omit(clause.Associations).Create(objPO).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.RoleDO{Name: objPO.Name}
	objDO.ID = objPO.ID
	return objDO, nil
}

// Update 更新
func (r *roleRepo) Update(ctx context.Context, obj *biz.RoleDO) (*biz.RoleDO, error) {
	objPO := &po.RolePO{
		Name:       obj.Name,
		Identifier: obj.Identifier,
	}
	//if obj.MenusAdmins != nil && len(obj.MenusAdmins) > 0 {
	//	objPos := make([]*po.MenusAdminPO, len(obj.MenusAdmins))
	//	for _, aobj := range obj.MenusAdmins {
	//		objPos = append(objPos, &po.MenusAdminPO{ObjectMeta: metaV1.ObjectMeta{
	//			ID: aobj.ID,
	//		}})
	//	}
	//	objPO.MenusAdmins = objPos
	//}
	if obj.Apis != nil && len(obj.Apis) > 0 {
		objPos := make([]*po.ApiPO, len(obj.Apis))
		for _, aobj := range obj.Apis {
			objPos = append(objPos, &po.ApiPO{ObjectMeta: metaV1.ObjectMeta{
				ID: aobj.ID,
			}})
		}
		objPO.Apis = objPos
	}
	tObj := &po.RolePO{}
	//tObj.ID = obj.ID
	err := r.data.Db.Model(&tObj).Where("id=?", obj.ID).Updates(&objPO).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.RoleDO{Name: objPO.Name}
	objDO.ID = objPO.ID
	return objDO, nil
}

// Delete 根据ID删除
func (r *roleRepo) Delete(c context.Context, id uint64) error {
	objPO := &po.RolePO{}
	objPO.ID = id
	err := r.data.Db.Delete(&objPO).Error
	return err
}

// DeleteList 根据ID批量删除
func (r *roleRepo) DeleteList(c context.Context, ids []uint64) error {
	objPO := &po.RolePO{}
	err := r.data.Db.Delete(&objPO, ids).Error
	return err
}

// FindByID 根据ID查询
func (r *roleRepo) FindByID(ctx context.Context, id uint64) (*po.RolePO, error) {
	obj := &po.RolePO{}
	err := r.data.Db.Where("id = ?", id).First(&obj).Error
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// FindByKey 根据角色Key名查询
func (r *roleRepo) FindByKey(c context.Context, identifier string) (*po.RolePO, error) {
	obj := &po.RolePO{}
	err := r.data.Db.Where("identifier = ?", identifier).First(&obj).Error
	return obj, err
}

// FindByUserId 获取用户拥有的角色
func (r *roleRepo) FindByUserId(c context.Context, userId uint64) ([]*biz.RoleDO, error) {
	db := r.data.Db
	rolePos := make([]*po.RolePO, 0)
	e := db.Where("ID in (?)", db.Table("qy_sys_user_role").Select("role_id ").Where(" user_id = ?", userId)).Find(&rolePos)
	if e.Error != nil {
		return nil, e.Error
	}
	infos := make([]*biz.RoleDO, 0, len(rolePos))
	for _, obj := range rolePos {
		infos = append(infos, &biz.RoleDO{
			ObjectMeta: metaV1.ObjectMeta{
				ID: obj.ID,
			},
			Name:       obj.Name,
			Identifier: obj.Identifier,
		})
	}
	return infos, nil
}

// ListAll 批量查询
func (r *roleRepo) ListAll(c context.Context, opts biz.RoleDOListOption) (*po.RolePOList, error) {
	ret := &po.RolePOList{}

	where := &po.RolePO{}
	var err error
	query := r.data.Db.Model(where)
	if len(opts.Name) > 0 {
		query.Where(" name like ? ", "%"+opts.Name+"%")
	}
	if opts.PageFlag {
		ol := gormutil.Unpointer(opts.Offset, opts.Limit)
		d := query.
			Offset(ol.Offset).
			Limit(ol.Limit).
			Order("id desc").
			Find(&ret.Items).
			Offset(-1).
			Limit(-1).
			Count(&ret.TotalCount)
		err = d.Error
	} else {
		d := query.
			Find(&ret.Items).
			Count(&ret.TotalCount)
		err = d.Error
	}
	opts.TotalCount = ret.TotalCount
	opts.IsLast()
	ret.FirstFlag = opts.FirstFlag
	ret.Current = opts.Current
	ret.PageSize = opts.PageSize
	ret.LastFlag = opts.LastFlag
	return ret, err
}
