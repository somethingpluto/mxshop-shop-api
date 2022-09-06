package router

import (
	"github.com/gin-gonic/gin"
	"order_web/api/order"
	"order_web/api/pay"
)

func InitOrderRouter(Router *gin.RouterGroup) {
	OrderRouter := Router.Group("orders")
	{
		OrderRouter.GET("", order.List)
		OrderRouter.POST("", order.New)
		OrderRouter.GET("/:id", order.Detail)
	}
	PayRouter := Router.Group("pay")
	{
		PayRouter.POST("alipy/notify", pay.Notify)
	}
}
