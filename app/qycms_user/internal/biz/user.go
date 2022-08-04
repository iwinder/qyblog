package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/api/qycms_user/v1"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms-common/meta/v1"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// UserBiz is a UserBiz model.
type UserBiz struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	Username          string `json:"username,omitempty" gorm:"unique;colum:username;type:varchar(255);not null"`
	NickName          string `json:"nickname" gorm:"column:nickname" validate:"required,min=1,max=30"`
	Avatar            string `json:"avatar" gorm:"column:avatar" validate:"omitempty"`
	Password          string `json:"password,omitempty" gorm:"column:password" validate:"required"`
	Salt              string `json:"-" gorm:"-" validate:"omitempty"`
	Email             string `json:"email" gorm:"column:email" validate:"required,email,min=1,max=100"`
	Phone             string `json:"phone" gorm:"column:phone" validate:"omitempty"`
	AdminFlag         bool   `json:"adminFlag,omitempty" gorm:"column:admin_flag" validate:"omitempty"`
}

func (u *UserBiz) TableName() string {
	return "qy_sys_user"
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *UserBiz) (*UserBiz, error)
	Update(context.Context, *UserBiz) (*UserBiz, error)
	FindByID(context.Context, uint64) (*UserBiz, error)
	ListByHello(context.Context, string) ([]*UserBiz, error)
	ListAll(context.Context) ([]*UserBiz, error)
}

// UserUsecase is a UserBiz usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewGreeterUsecase new a UserBiz usecase.
func NewGreeterUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateUser creates a UserBiz, and returns the new UserBiz.
func (uc *UserUsecase) CreateUser(ctx context.Context, g *UserBiz) (*UserBiz, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", g.Username)
	return uc.repo.Save(ctx, g)
}

func (uc *UserUsecase) FindOneUserByID(ctx context.Context, id uint64) (*UserBiz, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", id)
	return uc.repo.FindByID(ctx, id)
}
