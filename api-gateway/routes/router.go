package routes

import (
	"api-gateway/internal/handler"
	"api-gateway/middleware"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors(), middleware.InitMiddleware(service), middleware.ErrorMiddleware())
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mysession", store))
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

		}
	}
	return ginRouter
}
