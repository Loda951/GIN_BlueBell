package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
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
	api := router.Group("/api")
	{
		// 注册相关路由
		api.POST("/signup", controller.SignUpHandler)
		api.POST("/login", controller.LogInHandler)

		// 需要JWT认证的路由
		api.GET("/ping", middlewares.JWTAuthMiddleware(), func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "ok")
		})
	}
	return router
}
