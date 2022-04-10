package qycms_system

import (
	code "gitee.com/windcoder/qingyucms/internal/pkg/qycms-error-code"
	"gitee.com/windcoder/qingyucms/internal/pkg/qycms-middleware/auth"
	"gitee.com/windcoder/qingyucms/internal/pkg/qygo-common/core"
	errors "gitee.com/windcoder/qingyucms/internal/pkg/qygo-errors"
	"gitee.com/windcoder/qingyucms/internal/qycms-system/controller/v1/user"
	"gitee.com/windcoder/qingyucms/internal/qycms-system/store/mysql"
	"github.com/gin-gonic/gin"
)

func initRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)
}

func installMiddleware(g *gin.Engine) {

}

func installController(g *gin.Engine) *gin.Engine {
	jwtStrategy, _ := newJWTAuth().(auth.JWTStrategy)
	ad := g.Group("/admin/")
	ad.POST("/login", jwtStrategy.LoginHandler)
	ad.POST("/logout", jwtStrategy.LogoutHandler)
	ad.POST("/refresh", jwtStrategy.RefreshHandler)

	auto := newAutoAuth()
	g.NoRoute(auto.AuthFunc(), func(c *gin.Context) {
		core.WriteResponse(c, errors.WithCode(code.ErrPageNotFound, "Page not found."), nil)
	})

	storeIns, _ := mysql.GetMySQLFactoryOr(nil)
	v1 := ad.Group("/v1")
	{
		userv1 := v1.Group("/users")
		{
			userController := user.NewUserController(storeIns)
			userv1.POST("", userController.Create)
			userv1.PUT(":username", userController.Update)
			userv1.PUT(":username/change-password", userController.ChangePassword)
			userv1.DELETE("", userController.DeleteCollection)
			userv1.DELETE(":username", userController.Delete)
			userv1.GET("", userController.List)
			userv1.GET(":username", userController.GetWithUnscoped)

		}
		v1.Use(auto.AuthFunc())
	}

	return g
}
