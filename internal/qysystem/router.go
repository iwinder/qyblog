package qysystem

import (
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
	storeIns, _ := mysql.GetMySQLFactoryOr(nil)
	v1 := g.Group("/v1")
	{
		userv1 := v1.Group("/users")
		{
			userController := user.NewUserController(storeIns)
			userv1.POST("", userController.Create)
			userv1.GET(":username", userController.Get)
		}
	}

	return g
}
