package qycms

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitDefaultAPIs() {
	engine := gin.Default()
	engine.GET("/healthz", func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})
	engine.Run()
}
