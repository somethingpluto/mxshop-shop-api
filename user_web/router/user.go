package router

import (
	"github.com/gin-gonic/gin"
	"user_web/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	{
		userRouter.GET("list", api.GetUserList)
		userRouter.POST("pwd_login", api.PasswordLogin)
	}
}
