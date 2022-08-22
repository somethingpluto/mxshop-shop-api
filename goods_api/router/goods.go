package router

import "github.com/gin-gonic/gin"

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("goods")
	{
		// 获取商品列表
		GoodsRouter.GET("")
		// 获取商品详情
		GoodsRouter.POST("/:id")
		// 删除商品
		GoodsRouter.DELETE("/:id")
		// 获取商品库存
		GoodsRouter.GET("/:id/stocks")
		// 更新商品
		GoodsRouter.PUT("/:id")
		GoodsRouter.PATCH("/:id")
	}
}
