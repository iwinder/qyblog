package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/app/qycms_user/internal/biz"
	"github.com/iwinder/qingyucms/app/qycms_user/internal/data/po"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/auth"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/gormutil"
	"time"
)

var userCacheKey = func(username string) string {
	return "user_cache_key_" + username
}

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
		ObjectMeta: user.ObjectMeta,
		Username:   user.Username,
		Nickname:   user.Nickname,
		Password:   user.Password,
		Avatar:     user.Avatar,
		Salt:       user.Salt,
		Email:      user.Email,
		Phone:      user.Phone,
		AdminFlag:  user.AdminFlag,
	}
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
func (r *userRepo) FindByID(ctx context.Context, id uint64) (*po.UserPO, error) {
	cacheKey := userCacheKey(fmt.Sprintf("%d", id))
	user, err := r.getUserFromCache(ctx, cacheKey)
	if err != nil {
		user = &po.UserPO{}
		err := r.data.db.Where("id = ?", id).First(&user).Error
		if err != nil {
			return nil, biz.ErrUserNotFound
		}
		r.serUserCache(ctx, user, cacheKey)
	}

	if err != nil {
		return nil, err
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
		d := r.data.db.Model(where).Where(where).
			Offset(ol.Offset).
			Limit(ol.Limit).
			Order("id desc").
			Find(&ret.Items).
			Offset(-1).
			Limit(-1).
			Count(&ret.TotalCount)
		err = d.Error
	} else {
		d := r.data.db.Model(where).Where(where).
			Find(&ret.Items).
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

func (r *userRepo) VerifyPassword(ctx context.Context, u *biz.UserDO) (bool, error) {
	user := &po.UserPO{}
	err := r.data.db.Where("username=?", u.Username).First(&user).Error
	if err != nil {
		return false, err
	}
	aerr := auth.Compare(user.Password, u.Password+u.Salt)
	if aerr == nil {
		return true, nil
	}
	return false, aerr
}

func (r *userRepo) getUserFromCache(ctx context.Context, key string) (*po.UserPO, error) {
	result, err := r.data.redisCli.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var cacheUser = &po.UserPO{}
	err = json.Unmarshal([]byte(result), cacheUser)
	if err != nil {
		return nil, err
	}
	return cacheUser, nil
}

func (r *userRepo) serUserCache(ctx context.Context, user *po.UserPO, key string) {
	marshal, err := json.Marshal(user)
	if err != nil {
		r.log.Errorf("fail to set user cache:json.Marshal(%v) error(%v)", user, err)
	}
	err = r.data.redisCli.Set(ctx, key, string(marshal), time.Minute*30).Err()
	if err != nil {
		r.log.Errorf("fail to set user cache:redis.Set(%v) error(%v)", user, err)
	}
}
