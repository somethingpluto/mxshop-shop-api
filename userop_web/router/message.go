package router

import (
	"github.com/gin-gonic/gin"
	"userop_web/api/message"
	"userop_web/middlewares"
)

func InitMessageRouter(Router *gin.RouterGroup) {
	MessageRouter := Router.Group("message", middlewares.Trace())
	{
		MessageRouter.GET("", middlewares.JWTAuth(), message.List)
		MessageRouter.POST("", middlewares.JWTAuth(), message.New)
	}
}
