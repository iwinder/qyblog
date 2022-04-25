package mysql

import (
	"context"
	"gitee.com/windcoder/qingyucms/internal/pkg/qycms-common/gormutil"
	metav1 "gitee.com/windcoder/qingyucms/internal/pkg/qycms-common/meta/v1"
	code "gitee.com/windcoder/qingyucms/internal/pkg/qycms-error-code"
	errors "gitee.com/windcoder/qingyucms/internal/pkg/qygo-errors"
	v1 "gitee.com/windcoder/qingyucms/internal/qycms/models/v1"
	"gorm.io/gorm"
)

type userStore struct {
	db *gorm.DB
}

func newUserStore(ds *datastore) *userStore {
	return &userStore{
		db: ds.db,
	}
}

// Create 新增用户
func (u *userStore) Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error {

	return u.db.Create(&user).Error
}

// Update 更新用户信息-根据用户名
func (u *userStore) Update(ctx context.Context, user *v1.User, opts metav1.UpdateOptions) error {
	return u.db.Save(&user).Error
}

func (u *userStore) Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		u.db = u.db.Unscoped()
	}

	err := u.db.Where("username = ?", username).Delete(&v1.User{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}

func (u *userStore) DeleteCollection(ctx context.Context, usernames []string, opts metav1.DeleteOptions) error {
	//TODO implement me
	panic("implement me")
}

// Get 获取用户详情-根据 用户名
func (u *userStore) Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error) {
	user := &v1.User{}
	if opts.Unscoped {
		u.db = u.db.Unscoped()
	}
	err := u.db.Where("username=?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrUserNotFound, err.Error())
		}
		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}
	return user, nil
}

func (u *userStore) CountByUserName(ctx context.Context, username string, opts metav1.GetOptions) (int, error) {
	var total int64 = 0
	err := u.db.Where("username =?", username).Count(&total).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, errors.WithCode(code.ErrDatabase, err.Error())
	}
	return int(total), nil
}

func (u *userStore) List(ctx context.Context, opts v1.UserListOption) (*v1.UserList, error) {
	ret := &v1.UserList{}

	ol := gormutil.Unpointer(opts.Offset, opts.Limit)

	//selector, _ := fields.ParseSelector(opts.FieldSelector)
	//username, _ := selector.RequiresExactMatch("username")
	username := ""
	if opts.Unscoped {
		u.db = u.db.Unscoped()
	}
	d := u.db.Where("username like ?", "%"+username+"%").
		Offset(ol.Offset).
		Limit(ol.Limit).
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)

	return ret, d.Error
}

func (u *userStore) ListOptional(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error) {
	ret := &v1.UserList{}
	ol := gormutil.Unpointer(opts.Offset, opts.Limit)

	where := v1.User{}
	whereNot := v1.User{
		AdminFlag: false,
	}
	//selector, _ := fields.ParseSelector(opts.FieldSelector)
	//username, found := selector.RequiresExactMatch("username")
	username, found := "", false
	if found {
		where.Username = username
	}

	d := u.db.Where(where).
		Not(whereNot).
		Offset(ol.Offset).
		Limit(ol.Limit).
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)

	return ret, d.Error
}
