package router

import (
	"github.com/gin-gonic/gin"
	"user_web/api"
	"user_web/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	{
		userRouter.GET("list", middlewares.JWTAuth(), middlewares.AdminAuth(), api.GetUserList)
		userRouter.POST("pwd_login", api.PasswordLogin)
	}
}
