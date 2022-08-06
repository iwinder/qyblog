package db

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/iwinder/qingyucms/app/qycms_user/internal/data/po"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/gormutil"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/app/qycms_user/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Save 创建用户
func (r *userRepo) Save(ctx context.Context, user *biz.UserDO) (*po.UserPO, error) {
	userPO := &po.UserPO{
		Username:  user.Username,
		Nickname:  user.Nickname,
		Password:  user.Password,
		Avatar:    user.Avatar,
		Salt:      user.Salt,
		Email:     user.Email,
		Phone:     user.Phone,
		AdminFlag: user.AdminFlag,
	}
	userPO.InstanceID = user.InstanceID
	err := r.data.db.Create(userPO).Error
	if err != nil {
		return nil, err
	}
	return userPO, nil
}

// Update 更新用户
func (r *userRepo) Update(ctx context.Context, user *biz.UserDO) (*po.UserPO, error) {
	userPO := &po.UserPO{
		Username:  user.Username,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Email:     user.Email,
		Phone:     user.Phone,
		AdminFlag: user.AdminFlag,
	}
	tUser := &po.UserPO{}
	tUser.ID = user.ID
	err := r.data.db.Model(&tUser).Updates(&userPO).Error
	if err != nil {
		return nil, err
	}
	return userPO, nil
}

// Delete 根据ID删除用户
func (r *userRepo) Delete(c context.Context, id uint64) error {
	userPO := &po.UserPO{}
	userPO.ID = id
	err := r.data.db.Delete(&userPO).Error
	return err
}

// DeleteList 根据ID批量删除用户
func (r *userRepo) DeleteList(c context.Context, ids []uint64) error {
	userPO := &po.UserPO{}
	err := r.data.db.Delete(&userPO, ids).Error
	return err
}

// FindByID 根据ID查询用户信息
func (r *userRepo) FindByID(c context.Context, id uint64) (*po.UserPO, error) {
	user := &po.UserPO{}
	err := r.data.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%d,data not found,error is [%s]:", 111, err)
		}

		return nil, fmt.Errorf("%d,data query error,error is [%s]:", 222, err)
	}

	return user, nil
}

// FindByUsername 根据用户名查询用户
func (r *userRepo) FindByUsername(c context.Context, username string) (*po.UserPO, error) {
	user := &po.UserPO{}
	err := r.data.db.Where("id = ?", username).First(&user).Error
	return user, err
}

// ListAll 批量查询
func (r *userRepo) ListAll(c context.Context, opts biz.UserListOption) (*po.UserPOList, error) {
	ret := &po.UserPOList{}
	where := &po.UserPO{}
	var err error

	if opts.PageFlag {
		ol := gormutil.Unpointer(opts.Offset, opts.Limit)
		d := r.data.db.Where(where).
			Offset(ol.Offset).
			Limit(ol.Limit).
			Order("id desc").
			Find(&ret.Items).
			Offset(-1).
			Limit(-1).
			Count(&ret.TotalCount)
		err = d.Error
	} else {
		d := r.data.db.Where(where).
			Count(&ret.TotalCount)
		err = d.Error
	}
	opts.TotalCount = ret.TotalCount
	opts.IsLast()
	ret.FirstFlag = opts.FirstFlag
	ret.Page = opts.Page
	ret.PageSize = opts.PageSize
	ret.LastFlag = opts.LastFlag
	return ret, err
}
