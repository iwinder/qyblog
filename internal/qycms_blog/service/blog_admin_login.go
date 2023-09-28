package service

import (
	"context"
	v1 "github.com/iwinder/qyblog/api/qycms_bff/admin/v1"
	"github.com/iwinder/qyblog/internal/qycms_blog/biz"
	"github.com/iwinder/qyblog/internal/qycms_blog/conf"
)

type BlogAdminUserService struct {
	v1.UnimplementedQyAdminLoginServer
	v1.UnimplementedQyAdminUserServer
	v1.UnimplementedQyAdminRoleServer
	v1.UnimplementedQyAdminMenusAdminServer
	v1.UnimplementedQyAdminApiServer
	v1.UnimplementedQyAdminApiGroupServer
	v1.UnimplementedQyAdminFileServer
	v1.UnsafeQyAdminSiteConfigServer
	v1.UnimplementedQyAdminLinkServer
	v1.UnimplementedQyAdminShortLinkServer
	v1.UnimplementedQyAdminTagsServer
	conf     *conf.Qycms
	authConf *conf.Auth
	uc       *biz.UserUsecase
	rc       *biz.RoleUsecase
	ac       *biz.ApiUsecase
	acg      *biz.ApiGroupUsecase
	mc       *biz.MenusAdminUsecase
	rm       *biz.RoleMenusUsecase
	ra       *biz.RoleApiUsecase
	fit      *biz.FileLibTypeUsecase
	fic      *biz.FileLibConfigUsecase
	fi       *biz.FileLibUsecase
	site     *biz.SiteConfigUsecase
	lk       *biz.LinkUsecase
	slk      *biz.ShortLinkUsecase
	mau      *biz.MenusAgentUsecase
	mu       *biz.MenusUsecase
	tau      *biz.TagsUsecase
	cu       *biz.CategoryUsecase
	au       *biz.ArticleUsecase
	cc       *biz.CommentContentUsecase
	ctu      *biz.CommentUsecase
	siteMap  *biz.SiteMapUsecase
}

func NewBlogAdminUserService(uc *biz.UserUsecase, rc *biz.RoleUsecase, ac *biz.ApiUsecase, rm *biz.RoleMenusUsecase,
	acg *biz.ApiGroupUsecase, mc *biz.MenusAdminUsecase, ra *biz.RoleApiUsecase,
	fit *biz.FileLibTypeUsecase, fic *biz.FileLibConfigUsecase, fi *biz.FileLibUsecase,
	site *biz.SiteConfigUsecase, lk *biz.LinkUsecase, slk *biz.ShortLinkUsecase,
	mau *biz.MenusAgentUsecase, mu *biz.MenusUsecase,
	tau *biz.TagsUsecase, cu *biz.CategoryUsecase,
	au *biz.ArticleUsecase, cc *biz.CommentContentUsecase, ctu *biz.CommentUsecase,
	siteMap *biz.SiteMapUsecase,
	conf *conf.Qycms, authConf *conf.Auth) *BlogAdminUserService {
	return &BlogAdminUserService{uc: uc, rc: rc, ac: ac, rm: rm,
		acg: acg, mc: mc, ra: ra, fit: fit, fic: fic, fi: fi,
		site: site, lk: lk, slk: slk, mau: mau, mu: mu,
		tau: tau, cu: cu, au: au, cc: cc, ctu: ctu,
		siteMap: siteMap,
		conf:    conf, authConf: authConf}
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
