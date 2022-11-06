package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/admin/v1"
	w1 "github.com/iwinder/qingyucms/api/qycms_bff/web/v1"
	mid "github.com/iwinder/qingyucms/internal/pkg/qycms_common/auth/middleware"
	"github.com/iwinder/qingyucms/internal/qycms_blog/conf"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/db"
	"github.com/iwinder/qingyucms/internal/qycms_blog/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, authConf *conf.Auth, casbinData *db.CasbinData,
	userService *service.BlogAdminUserService,
	webService *service.BlogWebApiService,
	logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		mid.NewMiddleware(authConf, casbinData, logger),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	r := srv.Route("/api/admin/v1/file")
	r.POST("/upload/byType/{typeId}", func(ctx http.Context) error {
		return userService.UploadQyAdminFile(ctx)
	})
	v1.RegisterQyAdminLoginHTTPServer(srv, userService)
	v1.RegisterQyAdminRoleHTTPServer(srv, userService)
	v1.RegisterQyAdminUserHTTPServer(srv, userService)
	v1.RegisterQyAdminApiHTTPServer(srv, userService)
	v1.RegisterQyAdminMenusAdminHTTPServer(srv, userService)
	v1.RegisterQyAdminApiGroupHTTPServer(srv, userService)
	v1.RegisterQyAdminFileHTTPServer(srv, userService)
	v1.RegisterQyAdminSiteConfigHTTPServer(srv, userService)
	v1.RegisterQyAdminLinkHTTPServer(srv, userService)
	v1.RegisterQyAdminShortLinkHTTPServer(srv, userService)

	w1.RegisterQyWebSiteConfigHTTPServer(srv, webService)

	return srv
}
