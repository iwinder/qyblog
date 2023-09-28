package casbin

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/auth/auth_constants"
	"github.com/iwinder/qyblog/internal/qycms_blog/data/db"
	"strings"
	//"github.com/iwinder/qyblog/internal/qycms_blog/data/db"
	//db "github.com/iwinder/qyblog/internal/qycms_blog/data/db"
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
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if header, ok := transport.FromServerContext(ctx); ok {
				header.Operation()
				httpReq := header.(*http.Transport)
				url := httpReq.PathTemplate()
				if strings.Contains(url, "/api/web/v1/") {
					return handler(ctx, req)
				}
				method := httpReq.Request().Method

				claims, cok := jwt.FromContext(ctx)
				if !cok {
					return nil, ErrSecurityParseFailed
				}
				cla := claims.(jwtV4.MapClaims)

				if cla["AuthorityName"] == nil {
					return nil, ErrSecurityUserCreatorMissing
				}
				name := auth_constants.PrefixUser + cla["AuthorityName"].(string)
				roles, rerr := o.casbinData.Enf.GetRolesForUser(name, "*")
				if rerr != nil {
					return nil, rerr
				}
				var flag bool
				var eerr error
				for _, role := range roles {
					flag, eerr = o.casbinData.Enf.Enforce(role, "*", url, method)
					if eerr != nil {
						return nil, eerr
					}
					if flag {
						break
					}

				}
				if !flag {
					return nil, ErrUnauthorized
				}
				return handler(ctx, req)
			}
			return nil, ErrUnauthorized
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
