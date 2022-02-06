package auth

import (
	"gitee.com/windcoder/qingyucms/internal/pkg/qy-common/core"
	code "gitee.com/windcoder/qingyucms/internal/pkg/qy-error-code"
	errors "gitee.com/windcoder/qingyucms/internal/pkg/qy-errors"
	middleware "gitee.com/windcoder/qingyucms/internal/pkg/qy-middleware"
	"github.com/gin-gonic/gin"
	"strings"
)

type AutoStrategy struct {
	basic BasicStrategy
	jwt   JWTStrategy
}

func NewAutoStrategy(basic BasicStrategy, jwt JWTStrategy) AutoStrategy {
	return AutoStrategy{
		basic: basic,
		jwt:   jwt,
	}
}

func (a AutoStrategy) AuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		operator := middleware.AuthOperator{}
		authHeader := strings.SplitN(c.Request.Header.Get(middleware.KeyAuthorization), " ", 2)
		if len(authHeader) != 2 {
			core.WriteResponse(
				c,
				errors.WithCode(code.ErrInvalidAuthHeader, "Authorization header format is wrong."),
				nil,
			)
			c.Abort()

			return
		}

		switch authHeader[0] {
		case "Basic":
			operator.SetStrategy(a.basic)
		case "bearer":
			operator.SetStrategy(a.jwt)
		default:
			core.WriteResponse(c, errors.WithCode(code.ErrSignatureInvalid,
				"unrecognized Authorization header."), nil)
			c.Abort()
			return
		}

		operator.AuthFunc()
		c.Next()
	}
}