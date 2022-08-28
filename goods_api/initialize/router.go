package initialize

import (
	"github.com/gin-gonic/gin"
	"goods_api/router"
)

func InitRouter() *gin.Engine {
	Router := gin.Default()
	ApiGroup := Router.Group("/g/v1")
	router.InitGoodsRouter(ApiGroup)
	router.InitCategoryRouter(ApiGroup)
	router.InitBrandRouter(ApiGroup)
	router.InitBannerRouter(ApiGroup)
	return Router
}
