package router

import (
	"github.com/gin-gonic/gin"
	"goods_api/api/banner"
	"goods_api/middlewares"
)

func InitBannerRouter(Router *gin.RouterGroup) {
	BannerRouter := Router.Group("banners", middlewares.Trace())
	{
		BannerRouter.GET("", banner.List)
		BannerRouter.DELETE("/:id", banner.Delete)
		BannerRouter.POST("", banner.New)
		BannerRouter.PUT("/:id", banner.Update)
	}
}
