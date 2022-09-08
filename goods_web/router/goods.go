package router

import (
	"github.com/gin-gonic/gin"
	"goods_api/api/goods"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("goods")
	{
		// 获取商品列表
		GoodsRouter.GET("", goods.List)
		// 新建商品
		GoodsRouter.POST("", goods.New)
		// 获取商品详情
		GoodsRouter.GET("/:id", goods.Detail)
		// 删除商品
		GoodsRouter.DELETE("/:id", goods.Delete)
		// 获取商品库存
		//GoodsRouter.GET("/:id/stocks")
		// 更新商品
		GoodsRouter.PUT("/:id", goods.Update)
		// 更新商品状态
		GoodsRouter.PATCH("/:id", goods.UpdateStatus)
	}
}
