package router

import "github.com/gin-gonic/gin"

func InitBrandRouter(Router *gin.RouterGroup) {
	BrandRouter := Router.Group("brands")
	{
		// 获取品牌列表页
		BrandRouter.GET("")
		// 删除品牌
		BrandRouter.DELETE("/:id")
		// 创建品牌
		BrandRouter.POST("")
		// 更新品牌
		BrandRouter.PUT("/id")
	}
	CategoryBrandRouter := Router.Group("categorybrands")
	{
		// 类别品牌页
		CategoryBrandRouter.GET("")
		// 删除类别品牌
		CategoryBrandRouter.DELETE("/:id")
		// 新建类别品牌
		CategoryBrandRouter.POST("")
		// 修改类别品牌
		CategoryBrandRouter.PUT("/:id")
		// 获取分类的品牌
		CategoryBrandRouter.GET("/:id")
	}
}
