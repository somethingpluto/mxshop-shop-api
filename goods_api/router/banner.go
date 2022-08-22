package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitBannerRouter(Router *gin.RouterGroup) {
	BannerRouter := Router.Group("banners")
	{
		BannerRouter.GET("", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "hhh",
			})
		})
		BannerRouter.DELETE("/:id")
		BannerRouter.POST("")
		BannerRouter.PUT("/:id")
	}
}
