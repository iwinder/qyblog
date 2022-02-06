package qysystem

import (
	"context"
	"encoding/base64"
	v1 "gitee.com/windcoder/qingyucms/internal/pkg/qy-api/qysystem/v1"
	metav1 "gitee.com/windcoder/qingyucms/internal/pkg/qy-common/meta/v1"
	log "gitee.com/windcoder/qingyucms/internal/pkg/qy-log"
	middleware "gitee.com/windcoder/qingyucms/internal/pkg/qy-middleware"
	"gitee.com/windcoder/qingyucms/internal/pkg/qy-middleware/auth"
	"gitee.com/windcoder/qingyucms/internal/qysystem/config"
	"gitee.com/windcoder/qingyucms/internal/qysystem/store"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strings"
	"time"
)

const (
	// APIServerAudience defines the value of jwt audience field.
	APIServerAudience = "api.windcoder.com"

	// APIServerIssuer defines the value of jwt issuer field.
	APIServerIssuer = "qycms-system"
)

type loginInfo struct {
	Username string `form:"username" json:"username" validate:"required,username"`
	Password string `form:"password" json:"password" validate:"required,password"`
}

func newBasicAuth() middleware.AuthStrategy {
	return auth.NewBasicStrategy(func(username string, password string) bool {
		user, err := store.Client().Users().Get(context.TODO(), username, metav1.GetOptions{})
		if err != nil {
			return false
		}
		user.Salt = config.GetQYConfig().GetToken()
		if err := user.Compare(password); err != nil {
			return false
		}
		return true
	})
}

func newJWTAuth() middleware.AuthStrategy {
	gubJwt, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            viper.GetString("jwt.Realm"),
		SigningAlgorithm: "HS256",
		Key:              []byte(viper.GetString("jwt.key")),
		Timeout:          viper.GetDuration("jwt.timeout"),
		MaxRefresh:       viper.GetDuration("jwt.max-refresh"),
		Authenticator:    authenticator(), //登录验证
		Authorizator:     nil,             //权限验证
		PayloadFunc:      payloadFunc(),
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"message": message,
			})
		},
		LoginResponse: loginResponse(),
		LogoutResponse: func(c *gin.Context, code int) {
			c.JSON(http.StatusOK, nil)
		},
		RefreshResponse: refershResponse(),
		IdentityHandler: func(g *gin.Context) interface{} {
			claims := jwt.ExtractClaims(g)
			return claims[jwt.IdentityKey]
		},
		IdentityKey:   middleware.KeyUsername,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		//HTTPStatusMessageFunc: nil,
		//PrivKeyFile:           "",
		//PrivKeyBytes:          nil,
		//PubKeyFile:            "",
		//PrivateKeyPassphrase: "",
		//PubKeyBytes:          nil,
		SendCookie: true,
		//CookieMaxAge:      0,
		//SecureCookie:      false,
		//CookieHTTPOnly:    false,
		//CookieDomain:      "",
		//SendAuthorization: false,
		//DisabledAbort:     false,
		//CookieName:        "",
		//CookieSameSite:    0,
	})
	return auth.NewJWTStrategy(*gubJwt)
}

// 身份验证-登录校验
func authenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var login loginInfo
		var err error

		if c.Request.Header.Get(middleware.KeyAuthorization) != "" {
			login, err = parseWithHeader(c)
		} else {
			login, err = parseWithBody(c)
		}

		if err != nil {
			return "", jwt.ErrFailedAuthentication
		}

		user, err := store.Client().Users().Get(c, login.Username, metav1.GetOptions{})
		if err != nil {
			log.Errorf("get user information failed: %s", err.Error())

			return "", jwt.ErrFailedAuthentication
		}
		user.Salt = config.GetQYConfig().GetToken()
		if err := user.Compare(login.Password); err != nil {
			return "", jwt.ErrFailedAuthentication
		}
		return user, nil
	}
}

func parseWithHeader(c *gin.Context) (loginInfo, error) {
	auth := strings.SplitN(c.Request.Header.Get(middleware.KeyAuthorization), "", 2)
	if len(auth) != 2 || auth[0] != middleware.AuthBasic {
		log.Errorf("Get basic string from Authorization header failed")
		return loginInfo{}, jwt.ErrEmptyAuthHeader
	}

	payload, err := base64.StdEncoding.DecodeString(auth[1])
	if err != nil {
		log.Errorf("decode basic string: %s", err.Error())

		return loginInfo{}, jwt.ErrFailedAuthentication
	}

	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 {
		log.Errorf("parse payload failed")

		return loginInfo{}, jwt.ErrFailedAuthentication
	}

	return loginInfo{
		Username: pair[0],
		Password: pair[1],
	}, nil
}
func parseWithBody(c *gin.Context) (loginInfo, error) {
	var login loginInfo
	if err := c.ShouldBindJSON(&login); err != nil {
		log.Errorf("parse login parameters: %s", err.Error())
		return loginInfo{}, jwt.ErrFailedAuthentication
	}
	return login, nil
}
func refershResponse() func(c *gin.Context, code int, token string, expire time.Time) {
	return func(c *gin.Context, code int, token string, expire time.Time) {
		c.JSON(http.StatusOK, gin.H{
			"token":  token,
			"expire": expire.Format(time.RFC3339),
		})
	}
}

// 生成payload
func payloadFunc() func(data interface{}) jwt.MapClaims {
	return func(data interface{}) jwt.MapClaims {
		claims := jwt.MapClaims{
			"iss": APIServerIssuer,
			"aud": APIServerAudience,
		}
		if u, ok := data.(*v1.User); ok {
			claims[jwt.IdentityKey] = u.Username
			claims["sub"] = u.Username
		}
		return claims
	}
}

func loginResponse() func(c *gin.Context, code int, token string, expire time.Time) {
	return func(c *gin.Context, code int, token string, expire time.Time) {
		c.JSON(http.StatusOK, gin.H{
			"token":  token,
			"expire": expire.Format(time.RFC3339),
		})
	}
}

func newAutoAuth() middleware.AuthStrategy {
	return auth.NewAutoStrategy(newBasicAuth().(auth.BasicStrategy), newJWTAuth().(auth.JWTStrategy))
}
