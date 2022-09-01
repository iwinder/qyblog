package service

import (
	"context"
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/admin/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"github.com/iwinder/qingyucms/internal/qycms_blog/conf"
)

type BlogAdminService struct {
	v1.UnimplementedQyAdminLoginServer
	v1.UnimplementedQyAdminUserServer
	v1.UnimplementedQyAdminRoleServer
	v1.UnimplementedQyAdminMenusAdminServer
	v1.UnimplementedQyAdminApiServer
	conf     *conf.Qycms
	authConf *conf.Auth
	uc       *biz.UserUsecase
}

func NewBlogAdminService(uc *biz.UserUsecase, conf *conf.Qycms, authConf *conf.Auth) *BlogAdminService {
	return &BlogAdminService{uc: uc, conf: conf, authConf: authConf}
}

func (s *BlogAdminService) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	//req.Username
	token, err := s.uc.VerifyPassword(ctx, &biz.UserDO{Username: req.Username, Password: req.Password, Salt: s.conf.Token}, s.authConf)
	if err != nil {
		return nil, err
	}
	return &v1.LoginReply{Token: token}, nil
}

func (s *BlogAdminService) Logout(ctx context.Context, req *v1.LogoutReq) (*v1.LogoutReply, error) {
	return nil, nil
}
