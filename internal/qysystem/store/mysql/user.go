package mysql

import (
	"context"
	v1 "gitee.com/windcoder/qingyucms/internal/pkg/qy-api/qysystem/v1"
	"gitee.com/windcoder/qingyucms/internal/pkg/qy-common/auth/fields"
	metav1 "gitee.com/windcoder/qingyucms/internal/pkg/qy-common/meta/v1"
	gormutil "gitee.com/windcoder/qingyucms/internal/pkg/qy-common/utils/gormutil"
	code "gitee.com/windcoder/qingyucms/internal/pkg/qy-error-code"
	errors "gitee.com/windcoder/qingyucms/internal/pkg/qy-errors"
	"gorm.io/gorm"
)

type users struct {
	db *gorm.DB
}

func newUsers(ds *datastore) *users {
	return &users{
		db: ds.db,
	}
}

// Create 新增用户
func (u *users) Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error {

	return u.db.Create(&user).Error
}

// Update 更新用户信息-根据用户名
func (u *users) Update(ctx context.Context, user *v1.User, opts metav1.UpdateOptions) error {
	return u.db.Save(&user).Error
}

func (u *users) Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		u.db = u.db.Unscoped()
	}

	err := u.db.Where("username = ?", username).Delete(&v1.User{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}

func (u *users) DeleteCollection(ctx context.Context, usernames []string, opts metav1.DeleteOptions) error {
	//TODO implement me
	panic("implement me")
}

// Get 获取用户详情-根据 用户名
func (u *users) Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error) {
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

func (u *users) List(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error) {
	ret := &v1.UserList{}

	ol := gormutil.Unpointer(opts.Offset, opts.Limit)

	selector, _ := fields.ParseSelector(opts.FieldSelector)
	username, _ := selector.RequiresExactMatch("username")
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

func (u *users) ListOptional(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error) {
	ret := &v1.UserList{}
	ol := gormutil.Unpointer(opts.Offset, opts.Limit)

	where := v1.User{}
	whereNot := v1.User{
		AdminFlag: false,
	}
	selector, _ := fields.ParseSelector(opts.FieldSelector)
	username, found := selector.RequiresExactMatch("username")
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
