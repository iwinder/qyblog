package service

import (
	"context"
	v1 "github.com/iwinder/qingyucms/api/qycms_blog/admin/v1"
	"github.com/iwinder/qingyucms/app/qycms_blog/internal/biz"
)

func (s *ArticleService) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	reply, err := s.uc.CreateUser(ctx, &biz.UserDO{
		Username:  req.Username,
		Nickname:  req.Nickname,
		Avatar:    "",
		Password:  req.Password,
		Email:     "",
		Phone:     "",
		AdminFlag: false,
	})
	if err != nil {
		return nil, err
	}
	return &v1.RegisterReply{Uid: reply.ID}, nil
}

func (s *ArticleService) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	//req.Username
	token, err := s.uc.VerifyPassword(ctx, &biz.UserDO{Username: req.Username, Password: req.Password}, s.authConf)
	if err != nil {
		return nil, err
	}
	return &v1.LoginReply{Token: token}, nil
}

func (s *ArticleService) Logout(ctx context.Context, req *v1.LogoutReq) (*v1.LogoutReply, error) {
	return nil, nil
}
