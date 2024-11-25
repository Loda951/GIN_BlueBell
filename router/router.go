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

	// 注册相关路由
	api.POST("/signup", controller.SignUpHandler)
	// 登陆
	api.POST("/login", controller.LogInHandler)

	api.Use(middlewares.JWTAuthMiddleware())

	{
		api.GET("/community", controller.CommunityHandler)
		api.GET("/community/:id", controller.CommunityDetailHandler)

		api.POST("/post", controller.CreatPostHandler)
		api.GET("/post/:id", controller.GetPostDetailHandler)
		api.GET("/posts", controller.GetPostListHandler)

		api.POST("/vote", controller.PostVoteController)
	}

	// 需要JWT认证的路由
	api.GET("/ping", middlewares.JWTAuthMiddleware(), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})

	return router
}
