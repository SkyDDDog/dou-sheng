package routes

import (
	"api-gateway/internal/handler"
	"api-gateway/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors(), middleware.InitMiddleware(service), middleware.ErrorMiddleware())
	//store := cookie.NewStore([]byte("something-very-secret"))
	//ginRouter.Use(sessions.Sessions("mysession", store))
	ginRouter.Static("/static", "./public")
	v1 := ginRouter.Group("/douyin")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "pong")
		})
		// 用户服务
		v1.POST("/user/register/", handler.UserRegister)
		v1.POST("/user/login/", handler.UserLogin)
		v1.GET("/user/", handler.UserShow)

		// 需要登录保护
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			v1.GET("/feed/", handler.VideoFeed)
			v1.POST("/publish/action/", handler.ActionVideo)
			v1.GET("/publish/list/", handler.VideoList)
		}
	}
	return ginRouter
}
