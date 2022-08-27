package router

import (
	"github.com/gin-gonic/gin"
	"goods_api/api/category"
)

func InitCategoryRouter(Router *gin.RouterGroup) {
	CategoryRouter := Router.Group("categorys")
	{
		// 商品类别  列表
		CategoryRouter.GET("", category.List)
		// 删除分类
		CategoryRouter.DELETE("/:id", category.Delete)
		//// 分类详情
		CategoryRouter.GET("/:id", category.Detail)
		//// 新建分类
		CategoryRouter.POST("", category.New)
		//// 修改分类信息
		CategoryRouter.PUT("/:id", category.Update)
	}
}
