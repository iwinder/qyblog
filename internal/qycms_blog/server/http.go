package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/iwinder/qyblog/api/qycms_bff/admin/v1"
	w1 "github.com/iwinder/qyblog/api/qycms_bff/web/v1"
	mid "github.com/iwinder/qyblog/internal/pkg/qycms_common/auth/middleware"
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/filter"
	"github.com/iwinder/qyblog/internal/qycms_blog/conf"
	"github.com/iwinder/qyblog/internal/qycms_blog/data/db"
	"github.com/iwinder/qyblog/internal/qycms_blog/service"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, authConf *conf.Auth, casbinData *db.CasbinData,
	userService *service.BlogAdminUserService,
	webService *service.BlogWebApiService,
	logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		mid.NewMiddleware(authConf, casbinData, logger),
		filter.NewFilter(),
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
	r := srv.Route("/api/admin/v1")
	r.POST("/file/upload/byType/{typeId}", func(ctx http.Context) error {
		return userService.UploadQyAdminFile(ctx)
	})
	r.POST("/file/upload", func(ctx http.Context) error {
		return userService.UploadQyAdminFileDef(ctx)
	})
	r.GET("/metrics", func(ctx http.Context) error {
		promhttp.Handler().ServeHTTP(ctx.Response(), ctx.Request())
		return nil
	})

	//srv.Route("", )

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
	v1.RegisterQyAdminMenusAgentHTTPServer(srv, userService)
	v1.RegisterQyAdminMenusHTTPServer(srv, userService)
	v1.RegisterQyAdminTagsHTTPServer(srv, userService)
	v1.RegisterQyAdminCategoryHTTPServer(srv, userService)
	v1.RegisterQyAdminArticleHTTPServer(srv, userService)
	v1.RegisterQyAdminCommentHTTPServer(srv, userService)
	v1.RegisterQyAdminHomeHTTPServer(srv, userService)

	w1.RegisterQyWebSiteConfigHTTPServer(srv, webService)
	w1.RegisterQyWebArticleHTTPServer(srv, webService)
	w1.RegisterQyWebMenusHTTPServer(srv, webService)
	w1.RegisterQyWebLinksHTTPServer(srv, webService)
	w1.RegisterQyWebCommentHTTPServer(srv, webService)

	return srv
}
