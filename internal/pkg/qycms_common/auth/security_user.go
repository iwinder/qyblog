package auth

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport"
	jwtV4 "github.com/golang-jwt/jwt/v4"
)

const (
	ClaimAuthorityId = "authorityId"
)

type ASecurityUser struct {
	Path          string
	Method        string
	AuthorityName string
}

func NewSecurityUser() *ASecurityUser {
	return &ASecurityUser{}
}
func (su *ASecurityUser) ParseFromContext(ctx context.Context) error {
	if claims, ok := jwt.FromContext(ctx); ok {
		su.AuthorityName = claims.(jwtV4.MapClaims)[ClaimAuthorityId].(string)
	} else {
		return errors.New("jwt claim missing")
	}
	if header, ok := transport.FromServerContext(ctx); ok {
		su.Path = header.Operation()
		su.Method = "*"
	} else {
		return errors.New("jwt claim missing")
	}

	return nil
}
func (su *ASecurityUser) GetSubject() string {
	return su.AuthorityName
}
func (su *ASecurityUser) GetObject() string {
	return su.Path
}
func (su *ASecurityUser) GetAction() string {
	return su.Method
}
func (su *ASecurityUser) CreateAccessJwtToken(secretKey []byte) string {
	claims := jwtV4.NewWithClaims(jwtV4.SigningMethodHS256,
		jwtV4.MapClaims{
			ClaimAuthorityId: su.AuthorityName,
		})

	signedToken, err := claims.SignedString(secretKey)
	if err != nil {
		return ""
	}

	return signedToken
}
func (su *ASecurityUser) ParseAccessJwtTokenFromContext(ctx context.Context) error {
	claims, ok := jwt.FromContext(ctx)
	if !ok {
		return errors.New("no jwt token in context")
	}
	if err := su.ParseAccessJwtToken(claims); err != nil {
		return err
	}
	return nil
}
func (su *ASecurityUser) ParseAccessJwtTokenFromString(token string, secretKey []byte) error {
	parseAuth, err := jwtV4.Parse(token, func(*jwtV4.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}
	claims, ok := parseAuth.Claims.(jwtV4.MapClaims)
	if !ok {
		return errors.New("no jwt token in context")
	}

	if err := su.ParseAccessJwtToken(claims); err != nil {
		return err
	}

	return nil
}
func (su *ASecurityUser) ParseAccessJwtToken(claims jwtV4.Claims) error {
	if claims == nil {
		return errors.New("claims is nil")
	}

	mc, ok := claims.(jwtV4.MapClaims)
	if !ok {
		return errors.New("claims is not map claims")
	}

	strAuthorityId, ok := mc[ClaimAuthorityId]
	if ok {
		su.AuthorityName = strAuthorityId.(string)
	}

	return nil
}
