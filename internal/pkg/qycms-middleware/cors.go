package qycms_middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// 跨域访问
func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "POST", "DELETE", "GET", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           MaxAge * time.Hour,
		//  AllowOrigins 与 AllowOriginFunc 互斥,AllowOriginFunc 设置后 AllowOrigins 失效，AllowOriginFunc返回true时允许访问
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
	})
}
