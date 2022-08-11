package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	v1 "github.com/iwinder/qingyucms/api/qycms_blog/admin/v1"
	"github.com/iwinder/qingyucms/app/qycms_blog/internal/conf"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"time"
)

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

type UserRepo interface {
	GetUser(ctx context.Context, id uint64) (*UserDO, error)
	CreateUser(ctx context.Context, user *UserDO) (*UserDO, error)
	VerifyPassword(ctx context.Context, user *UserDO) (bool, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	log := log.NewHelper(log.With(logger, "module", "usecase/interface"))
	return &UserUseCase{
		repo: repo,
		log:  log,
	}
}

func (r *UserUseCase) CreateUser(ctx context.Context, user *UserDO) (*UserDO, error) {
	reply, err := r.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return reply, err
}

func (r *UserUseCase) GetUser(ctx context.Context, id uint64) (*UserDO, error) {
	reply, err := r.repo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return reply, err
}

func (r *UserUseCase) VerifyPassword(ctx context.Context, user *UserDO, authConf *conf.Auth) (string, error) {
	reply, err := r.repo.VerifyPassword(ctx, user)
	if !reply || err != nil {
		return "", fmt.Errorf("登录失败:%w", v1.ErrorReason_ErrPasswordWrong)
	}
	// 生成token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(authConf.ExpireDuration.AsDuration())), // 设置token的过期时间
	})
	token, err := claims.SignedString([]byte(authConf.JwtSecret))
	if err != nil {
		log.Errorf("登录失败，生成token失败：%v", err)
		return "", fmt.Errorf("登录失败")
	}
	return token, err
}
