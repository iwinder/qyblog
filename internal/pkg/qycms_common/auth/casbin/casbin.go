package casbin

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/db"
	//"github.com/iwinder/qingyucms/internal/qycms_blog/data/db"
	//db "github.com/iwinder/qingyucms/internal/qycms_blog/data/db"
)

const (
	reason string = "FORBIDDEN"
)

var (
	ErrSecurityUserCreatorMissing = errors.Forbidden(reason, "ASecurityUser is required")
	ErrEnforcerMissing            = errors.Forbidden(reason, "Enforcer is missing")
	ErrSecurityParseFailed        = errors.Forbidden(reason, "Security Info fault")
	ErrUnauthorized               = errors.Forbidden(reason, "Unauthorized Access")
)

type Option func(*options)

type options struct {
	enableDomain bool
	//securityUser jwt2.SecurityUser
	casbinData *db.CasbinData
	//model    model.Model
	//policy   persist.Adapter
	//enforcer *stdcasbin.SyncedEnforcer

}

func WithCasbinData(data *db.CasbinData) Option {
	return func(o *options) {
		o.casbinData = data
	}
}

func Server(opts ...Option) middleware.Middleware {
	//o := &options{
	//	securityUser: nil,
	//}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if header, ok := transport.FromServerContext(ctx); ok {
				header.Operation()
				s := header.RequestHeader()
				ss := header.(*http.Transport)
				sss := ss.Request().Method
				fmt.Print(s, sss)
				//&(http.Transport)header
				//opts.casbinData.Enf
				// Do something on entering
				//defer func() {
				//	// Do something on exiting
				//}()
			}
			return handler(ctx, req)
		}
	}
}

func NewCasbinRuleMatcher() selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {

		claims, ok := jwt.FromContext(ctx)
		if !ok {
			return true
		}
		c := claims.(jwtV4.MapClaims)
		if c["AuthorityId"] == nil {
			return true
		}
		authorityId := uint64(c["AuthorityId"].(float64))
		log.Debug(authorityId)
		return true
	}
}
