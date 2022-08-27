package router

import (
	"github.com/gin-gonic/gin"
	"goods_api/api/brand"
	"goods_api/api/categoryBrand"
)

func InitBrandRouter(Router *gin.RouterGroup) {
	BrandRouter := Router.Group("brands")
	{
		// 获取品牌列表页
		BrandRouter.GET("", brand.List)
		// 删除品牌
		BrandRouter.DELETE("/:id", brand.Delete)
		//// 创建品牌
		BrandRouter.POST("", brand.New)
		//// 更新品牌
		BrandRouter.PUT("/:id", brand.Update)
	}
	CategoryBrandRouter := Router.Group("categorybrands")
	{
		// 类别品牌页
		CategoryBrandRouter.GET("", categoryBrand.List)
		// 删除类别品牌
		//CategoryBrandRouter.DELETE("/:id")
		//// 新建类别品牌
		CategoryBrandRouter.POST("", categoryBrand.New)
		//// 修改类别品牌
		//CategoryBrandRouter.PUT("/:id")
		//// 获取分类的品牌
		CategoryBrandRouter.GET("/:id", categoryBrand.Detail)
	}
}
