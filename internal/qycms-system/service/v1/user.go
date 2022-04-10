package v1

import (
	"context"
	v1 "gitee.com/windcoder/qingyucms/internal/pkg/qycms-api/qycms-system/v1"
	code "gitee.com/windcoder/qingyucms/internal/pkg/qycms-error-code"
	metav1 "gitee.com/windcoder/qingyucms/internal/pkg/qygo-common/meta/v1"
	errors "gitee.com/windcoder/qingyucms/internal/pkg/qygo-errors"
	log "gitee.com/windcoder/qingyucms/internal/pkg/qygo-log"
	"gitee.com/windcoder/qingyucms/internal/qycms-system/store"
	"sync"
)

type UserSrv interface {
	Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error
	Update(ctx context.Context, user *v1.User, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, usernames []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error)
	ListOptional(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error)
	ChangePassword(ctx context.Context, user *v1.User) error
}

type userService struct {
	store store.Factory
}

func newUsers(srv *service) *userService {
	return &userService{
		store: srv.store,
	}
}

func (u *userService) Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error {
	if err := u.store.Users().Create(ctx, user, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}

func (u *userService) Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error) {
	user, err := u.store.Users().Get(ctx, username, opts)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userService) Update(ctx context.Context, user *v1.User, opts metav1.UpdateOptions) error {
	if err := u.store.Users().Update(ctx, user, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}

func (u *userService) Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error {
	if err := u.store.Users().Delete(ctx, username, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}

func (u *userService) DeleteCollection(ctx context.Context, usernames []string, opts metav1.DeleteOptions) error {
	if err := u.store.Users().DeleteCollection(ctx, usernames, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

func (u *userService) List(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error) {
	users, err := u.store.Users().List(ctx, opts)
	if err != nil {
		log.L(ctx).Errorf("list users from storage failed: %s", err.Error())

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}
	var m sync.Map
	for _, user := range users.Items {
		m.Store(user.ID, &v1.User{
			ObjectMeta: metav1.ObjectMeta{
				ID:         user.ID,
				InstanceID: user.InstanceID,
				Extend:     user.Extend,
				CreatedAt:  user.CreatedAt,
				UpdatedAt:  user.UpdatedAt,
			},
			Nickname: user.Nickname,
			Email:    user.Email,
			Phone:    user.Phone,
			Username: user.Username,
			//TotalPolicy: policies.TotalCount,
		})
	}
	infos := make([]*v1.User, 0, len(users.Items))
	for _, user := range users.Items {
		info, _ := m.Load(user.ID)
		infos = append(infos, info.(*v1.User))
	}

	log.L(ctx).Debugf("get %d users from backend storage.", len(infos))

	return &v1.UserList{ListMeta: users.ListMeta, Items: infos}, nil
}
func (u *userService) ListOptional(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error) {
	return nil, nil
}
func (u *userService) ChangePassword(ctx context.Context, user *v1.User) error {

	if err := u.store.Users().Update(ctx, user, metav1.UpdateOptions{}); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}
