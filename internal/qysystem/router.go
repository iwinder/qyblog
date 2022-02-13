package qysystem

import (
	"gitee.com/windcoder/qingyucms/internal/pkg/qy-common/core"
	code "gitee.com/windcoder/qingyucms/internal/pkg/qy-error-code"
	errors "gitee.com/windcoder/qingyucms/internal/pkg/qy-errors"
	"gitee.com/windcoder/qingyucms/internal/pkg/qy-middleware/auth"
	"gitee.com/windcoder/qingyucms/internal/qysystem/controller/v1/user"
	"gitee.com/windcoder/qingyucms/internal/qysystem/store/mysql"
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
