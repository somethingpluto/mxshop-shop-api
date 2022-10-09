package router

import (
	"github.com/gin-gonic/gin"
	"order_web/api/order"
	"order_web/api/pay"
	"order_web/middlewares"
)

func InitOrderRouter(Router *gin.RouterGroup) {
	OrderRouter := Router.Group("orders", middlewares.Trace())
	{
		OrderRouter.GET("", middlewares.JWTAuth(), order.List)
		OrderRouter.POST("", middlewares.JWTAuth(), order.New)
		OrderRouter.GET("/:id", middlewares.JWTAuth(), order.Detail)
	}
	PayRouter := Router.Group("pay")
	{
		PayRouter.POST("alipy/notify", pay.Notify)
	}
}
