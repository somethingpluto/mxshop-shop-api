package router

import "github.com/gin-gonic/gin"

func InitCategoryRouter(Router *gin.RouterGroup) {
	CategoryRouter := Router.Group("categorys")
	{
		// 商品类别  列表
		CategoryRouter.GET("")
		// 删除分类
		CategoryRouter.DELETE("/:id")
		// 分类详情
		CategoryRouter.GET("/:id")
		// 新建分类
		CategoryRouter.POST("")
		// 修改分类信息
		CategoryRouter.PUT("/:id")
	}
}
