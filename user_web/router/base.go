package router

import (
	"github.com/gin-gonic/gin"
	"user_web/api"
	"user_web/middlewares"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("base", middlewares.Trace())
	{
		BaseRouter.GET("/captcha", api.GetCaptcha)
		BaseRouter.POST("/note_code", api.SendNoteCode)
	}

}
