package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/api/qycms_user/v1"
	"github.com/iwinder/qingyucms/app/qycms_user/internal/data/po"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"gorm.io/gorm"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// UserDO is a UserDO model.
type UserDO struct {
	metaV1.ObjectMeta
	Username  string
	Nickname  string
	Avatar    string
	Password  string
	Salt      string
	Email     string
	Phone     string
	AdminFlag bool
}
type UserDOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*UserDO `json:"items"`
}

type UserListOption struct {
	metaV1.ListOptions `json:"page"`
	UserDO             `json:"user"`
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *UserDO) (*po.UserPO, error)
	Update(context.Context, *UserDO) (*po.UserPO, error)
	Delete(context.Context, uint64) error
	DeleteList(c context.Context, uids []uint64) error
	FindByID(context.Context, uint64) (*po.UserPO, error)
	FindByUsername(c context.Context, username string) (*po.UserPO, error)
	ListAll(c context.Context, opts UserListOption) (*po.UserPOList, error)
}

// UserUsecase is a UserDO usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewGreeterUsecase new a UserDO usecase.
func NewGreeterUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateUser creates a UserDO, and returns the new UserDO.
func (uc *UserUsecase) CreateUser(ctx context.Context, user *UserDO) (*UserDO, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", user.Username)
	userPO, err := uc.repo.Save(ctx, user)
	if err != nil {
		return nil, err
	}
	userDO := &UserDO{Username: userPO.Username}
	userDO.ID = userPO.ID
	return userDO, nil
}

// Update 更新用户
func (uc *UserUsecase) Update(ctx context.Context, user *UserDO) (*UserDO, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", user.Username)
	userPO, err := uc.repo.Update(ctx, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	userDO := &UserDO{Username: userPO.Username}
	userDO.ID = userPO.ID
	return userDO, nil
}

// Delete 根据ID删除用户
func (uc *UserUsecase) Delete(ctx context.Context, id uint64) error {
	uc.log.WithContext(ctx).Infof("Delete: %v", id)
	return uc.repo.Delete(ctx, id)
}

// DeleteList 根据ID批量删除用户
func (uc *UserUsecase) DeleteList(ctx context.Context, ids []uint64) error {
	uc.log.WithContext(ctx).Infof("DeleteList: %v", ids)
	return uc.repo.DeleteList(ctx, ids)
}

// FindOneByID 根据ID查询用户信息
func (uc *UserUsecase) FindOneByID(ctx context.Context, id uint64) (*UserDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByID: %v", id)
	user, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	userDO := &UserDO{
		ObjectMeta: user.ObjectMeta,
		Username:   user.Username,
		Nickname:   user.Nickname,
		Avatar:     user.Avatar,
		//Password:   user.Password,
		//Salt:       user.Salt,
		Email:     user.Email,
		Phone:     user.Phone,
		AdminFlag: user.AdminFlag,
	}
	return userDO, nil
}

// FindOneByUsername 根据用户名查询用户信息
func (uc *UserUsecase) FindOneByUsername(ctx context.Context, username string) (*UserDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByUsername: %v", username)
	user, err := uc.repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	userDO := &UserDO{
		ObjectMeta: user.ObjectMeta,
		Username:   user.Username,
		Nickname:   user.Nickname,
		Avatar:     user.Avatar,
		Password:   user.Password,
		Salt:       user.Salt,
		Email:      user.Email,
		Phone:      user.Phone,
		AdminFlag:  user.AdminFlag,
	}
	return userDO, nil
}

// ListAll 批量查询
func (uc *UserUsecase) ListAll(ctx context.Context, opts UserListOption) (*UserDOList, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	userPOs, err := uc.repo.ListAll(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	infos := make([]*UserDO, 0, len(userPOs.Items))
	for _, user := range userPOs.Items {
		infos = append(infos, &UserDO{
			ObjectMeta: metaV1.ObjectMeta{
				ID:         user.ID,
				InstanceID: user.InstanceID,
				Extend:     user.Extend,
				CreatedAt:  user.CreatedAt,
				UpdatedAt:  user.UpdatedAt,
			},
			Username: user.Username,
			Avatar:   user.Avatar,
			Nickname: user.Nickname,
			Email:    user.Email,
			Phone:    user.Phone,
		})
	}
	return &UserDOList{ListMeta: userPOs.ListMeta, Items: infos}, nil
}
