package service

import (
	"context"

	v1 "github.com/iwinder/qingyucms/api/helloworld/v1"
	"github.com/iwinder/qingyucms/app/qycms-user/internal/biz"
)

// UserService is a greeter service.
type UserService struct {
	v1.UnimplementedGreeterServer

	uc *biz.UserUsecase
}

// NewUserService new a greeter service.
func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *UserService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateUser(ctx, &biz.UserBiz{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
