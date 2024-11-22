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
	router.POST("/signup", controller.SignUpHandler)
	router.POST("/login", controller.LogInHandler)

	router.GET("/ping", middlewares.JWTAuthMiddleware(), func(ctx *gin.Context) {
		// 如果是登陆用户 判断请求头是否有 有效的JWT
		ctx.JSON(http.StatusOK, "ok")
	})
	return router
}
