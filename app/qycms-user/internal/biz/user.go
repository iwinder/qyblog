package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/iwinder/qingyucms/api/helloworld/v1"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// UserBiz is a UserBiz model.
type UserBiz struct {
	Hello string
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *UserBiz) (*UserBiz, error)
	Update(context.Context, *UserBiz) (*UserBiz, error)
	FindByID(context.Context, int64) (*UserBiz, error)
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
	uc.log.WithContext(ctx).Infof("CreateUser: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}

func (uc *UserUsecase) FindOneUserByID(ctx context.Context, id int64) (*UserBiz, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", id)
	return uc.repo.FindByID(ctx, id)
}
