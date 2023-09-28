package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"
	"github.com/iwinder/qyblog/internal/qycms_blog/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewBlogAdminUserService,
	NewBlogWebApiService, NewBlogAdminJobsService,
)

func GetUserInfo(ctx context.Context) *biz.UserInfoDO {
	userIndo := &biz.UserInfoDO{}
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwtV4.MapClaims)
		if c["ID"] == nil {
			return nil
		}
		userIndo.ID = uint64(c["ID"].(float64))
		userIndo.Nickname = c["NickName"].(string)
	}
	return userIndo
}

func GetHeardInfo(ctx context.Context) (string, string) {
	userAgent := ""
	userIp := ""
	if header, ok := transport.FromServerContext(ctx); ok {
		header.Operation()
		userIp = header.RequestHeader().Get("X-Remoteaddr")
		userAgent = header.RequestHeader().Get("User-Agent")

	}
	return userAgent, userIp
}
