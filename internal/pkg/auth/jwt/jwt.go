package jwt

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	"github.com/iwinder/qyblog/internal/pkg/auth"
	"strings"
)

type SecurityUser struct {
	auth.ASecurityUser
	ID            uint64
	NickName      string
	AuthorityName string
	RoleIds       string
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
	whiteList["/api.qycms_bff.admin.v1.QyAdminLogin/Login"] = true
	whiteList["/api.qycms_bff.admin.v1.QyAdminUser/CreateUser"] = true

	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		if strings.Contains(operation, "/api.qycms_bff.web.") {
			return false
		}
		return true
	}
}
