package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	userv1 "github.com/iwinder/qingyucms/api/qycms_user/v1"
	"github.com/iwinder/qingyucms/app/qycms_blog/internal/biz"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, user *biz.UserDO) (*biz.UserDO, error) {
	r.log.WithContext(ctx).Infof("CreateUser: %v", user.Username)
	userPO, err := r.data.uc.CreateUser(ctx, &userv1.CreateUserRequest{
		Username: user.Username,
		Password: user.Password,
		NickName: user.Nickname,
		Avatar:   user.Avatar,
		Email:    user.Email,
		Phone:    user.Phone,
	})
	if err != nil {
		return nil, err
	}
	userDO := &biz.UserDO{Username: userPO.User.Username, Nickname: userPO.User.NickName}
	userDO.ID = userPO.User.Uid
	return userDO, nil
}
func (r *userRepo) GetUser(ctx context.Context, id uint64) (*biz.UserDO, error) {
	reply, err := r.data.uc.GetUser(ctx, &userv1.GetUserRequest{
		Uid: id,
	})
	if err != nil {
		return nil, err
	}
	user := &biz.UserDO{
		Username: reply.User.Username,
		Nickname: reply.User.NickName,
	}
	user.ID = reply.User.Uid
	return user, err
}

func (r *userRepo) VerifyPassword(ctx context.Context, user *biz.UserDO) (bool, error) {
	rv, err := r.data.uc.VerifyPassword(ctx, &userv1.VerifyPasswordReq{Username: user.Username, Password: user.Password})

	return rv.Ok, err
}
