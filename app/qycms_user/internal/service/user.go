package service

import (
	"context"
	"github.com/iwinder/qingyucms/api/qycms_user/v1"
	"github.com/iwinder/qingyucms/app/qycms_user/internal/biz"
)

// UserService is a greeter service.
type UserService struct {
	v1.UnimplementedUserServer

	uc *biz.UserUsecase
}

// NewUserService new a greeter service.
func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

// CreateUser implements qycms_user.CreateUser.
func (s *UserService) CreateUser(ctx context.Context, in *v1.CreateUserRequest) (*v1.CreateUserReply, error) {
	user, err := s.uc.CreateUser(ctx, &biz.UserBiz{
		Username: in.Username,
		NickName: in.NickName,
		Password: in.Password,
	})
	if err != nil {
		return nil, err
	}
	u := UserResponse(user)
	return &v1.CreateUserReply{User: &u}, nil
}

func (s *UserService) GetUser(ctx context.Context, in *v1.GetUserRequest) (*v1.GetUserReply, error) {
	user, err := s.uc.FindOneUserByID(ctx, in.Uid)
	if err != nil {
		return nil, err
	}
	u := UserResponse(user)
	return &v1.GetUserReply{User: &u}, nil
}

func UserResponse(user *biz.UserBiz) v1.UserInfoResponse {
	userInfoRsp := v1.UserInfoResponse{
		Uid:      user.ID,
		Username: user.Username,
		Password: user.Password,
		NickName: user.NickName,
	}

	return userInfoRsp
}
