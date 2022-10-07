package service

import (
	"context"
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/admin/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"github.com/iwinder/qingyucms/internal/qycms_blog/conf"
)

type BlogAdminUserService struct {
	v1.UnimplementedQyAdminLoginServer
	v1.UnimplementedQyAdminUserServer
	v1.UnimplementedQyAdminRoleServer
	v1.UnimplementedQyAdminMenusAdminServer
	v1.UnimplementedQyAdminApiServer
	v1.UnimplementedQyAdminApiGroupServer
	conf     *conf.Qycms
	authConf *conf.Auth
	uc       *biz.UserUsecase
	rc       *biz.RoleUsecase
	ac       *biz.ApiUsecase
	acg      *biz.ApiGroupUsecase
	mc       *biz.MenusAdminUsecase
	rm       *biz.RoleMenusUsecase
	ra       *biz.RoleApiUsecase
}

func NewBlogAdminUserService(uc *biz.UserUsecase, rc *biz.RoleUsecase, ac *biz.ApiUsecase, rm *biz.RoleMenusUsecase,
	acg *biz.ApiGroupUsecase, mc *biz.MenusAdminUsecase, ra *biz.RoleApiUsecase,
	conf *conf.Qycms, authConf *conf.Auth) *BlogAdminUserService {
	return &BlogAdminUserService{uc: uc, rc: rc, ac: ac, rm: rm,
		acg: acg, mc: mc, ra: ra,
		conf: conf, authConf: authConf}
}

func (s *BlogAdminUserService) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	//req.Username
	user, err := s.uc.VerifyPassword(ctx, &biz.UserDO{Username: req.Username, Password: req.Password, Salt: s.conf.Token}, s.authConf)
	if err != nil {
		return nil, err
	}
	return &v1.LoginReply{
		Token:     user.Token,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		RoleNames: user.RoleNames,
	}, nil
}

func (s *BlogAdminUserService) Logout(ctx context.Context, req *v1.LogoutReq) (*v1.LogoutReply, error) {
	return nil, nil
}
