package middleware

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	qyjwt "github.com/iwinder/qingyucms/internal/pkg/qycms_common/auth/jwt"
	"github.com/iwinder/qingyucms/internal/qycms_blog/conf"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/db"
)

func NewMiddleware(authConf *conf.Auth, casbinData *db.CasbinData, logger log.Logger) http.ServerOption {
	return http.Middleware(
		recovery.Recovery(),
		tracing.Server(),
		logging.Server(logger),
		selector.Server(
			jwt.Server(
				func(token *jwtV4.Token) (interface{}, error) {
					return []byte(authConf.Jwt.JwtSecret), nil
				},
				jwt.WithSigningMethod(jwtV4.SigningMethodHS256),
			),
			//casbin.Server(
			//	casbin.WithCasbinData(casbinData),
			//),
		).Match(qyjwt.NewWhiteListMatcher()).Build(),
	)
}
