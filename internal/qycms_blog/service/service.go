package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService,
	NewArticleService, NewArticleContentService,
	NewCommentAgentService, NewBlogAdminUserService,
	NewBlogWebApiService,
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
