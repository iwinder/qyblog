package jwt

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/auth"
)

type SecurityUser struct {
	auth.ASecurityUser
	ID          uint64
	NickName    string
	AuthorityId uint64
	jwtV4.RegisteredClaims
}

// CreateToken generate token
func CreateToken(c SecurityUser, key string) (string, error) {
	claims := jwtV4.NewWithClaims(jwtV4.SigningMethodHS256, c)
	signedString, err := claims.SignedString([]byte(key))
	if err != nil {
		return "", errors.New("generate token failed" + err.Error())
	}

	return signedString, nil
}

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]bool)
	whiteList["/api.qycms_blog.admin.v1.QyBlogAdminLogin/Login"] = true
	whiteList["/api.qycms_user.v1.User/CreateUser"] = true

	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
