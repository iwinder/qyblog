package qycms

import (
	"gitee.com/windcoder/qingyucms/internal/qycms/controller/v1/admin/user"
	"gitee.com/windcoder/qingyucms/internal/qycms/store/mysql"
	"github.com/gin-gonic/gin"
)

//func InitDefaultAPIs() {
//	engine := gin.Default()
//	engine.GET("/healthz", func(context *gin.Context) {
//		context.JSON(http.StatusOK, map[string]string{"status": "ok"})
//	})
//
//	engine.Run()
//}

func initRouter(g *gin.Engine) {
	installController(g)
}

func installController(g *gin.Engine) *gin.Engine {
	storeIns, _ := mysql.GetMySQLFactoryOr(nil)
	v1 := g.Group("/v1")
	{
		adminV1 := v1.Group("/admin")
		{
			userv1 := adminV1.Group("/users")
			{
				userController := user.NewUserController(storeIns)
				userv1.POST("", userController.Create)                                  // 创建新用户
				userv1.PUT("/:username", userController.Update)                         // 更新新用户
				userv1.PUT("/:username/change-password", userController.ChangePassword) // 用户修改密码
				//userv1.DELETE("", userController.DeleteCollection)
				userv1.DELETE(":username", userController.Delete)     // 删除用户
				userv1.GET("", userController.List)                   // 获取用户列表
				userv1.GET(":username", userController.GetByUserName) // 获取用户详情
			}
		}

	}
	return g
}
