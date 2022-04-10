package auth

import (
	"encoding/base64"
	code "gitee.com/windcoder/qingyucms/internal/pkg/qycms-error-code"
	middleware "gitee.com/windcoder/qingyucms/internal/pkg/qycms-middleware"
	"gitee.com/windcoder/qingyucms/internal/pkg/qygo-common/core"
	errors "gitee.com/windcoder/qingyucms/internal/pkg/qygo-errors"
	"github.com/gin-gonic/gin"
	"strings"
)

type BasicStrategy struct {
	compare func(username, password string) bool
}

func NewBasicStrategy(compare func(username string, password string) bool) BasicStrategy {
	return BasicStrategy{compare: compare}
}

func (b BasicStrategy) AuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := strings.SplitN(c.Request.Header.Get(middleware.KeyAuthorization), "", 2)
		if len(auth) != 2 || auth[0] != middleware.AuthBasic {
			core.WriteResponse(c, errors.WithCode(code.ErrSignatureInvalid, "Authorization header format is wrong."), nil)
			c.Abort()
			return
		}
		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !b.compare(pair[0], pair[1]) {
			core.WriteResponse(c,
				errors.WithCode(code.ErrSignatureInvalid,
					"Authorization header format is wrong."), nil)
			c.Abort()
			return
		}

		c.Set(middleware.KeyUsername, pair[0])
		c.Next()

	}

}
