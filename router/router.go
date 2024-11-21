package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.Use(logger.GinLogger(), logger.GinRecovery(true))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	// 注册
	router.POST("/signup", controller.SignUpHandler)
	router.POST("/login", controller.LogInHandler)
	return router
}
