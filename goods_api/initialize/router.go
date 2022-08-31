package initialize

import (
	"github.com/gin-gonic/gin"
	"goods_api/router"
)

func InitRouter() *gin.Engine {
	Router := gin.Default()
	ApiGroup := Router.Group("/g/v1")
	router.InitBannerRouter(ApiGroup)
	router.InitGoodsRouter(ApiGroup)
	router.InitCategoryRouter(ApiGroup)
	router.InitBrandRouter(ApiGroup)
	return Router
}
