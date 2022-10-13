package categoryBrand

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goods_api/forms"
	"goods_api/global"
	"goods_api/proto"
	"goods_api/utils"
	"net/http"
	"strconv"
)

// List
// @Description: 获取商品目录品牌列表
// @param ctx
//
func List(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}

	response, err := global.GoodsClient.CategoryBrandList(context.WithValue(context.Background(), "ginContext", ctx), &proto.CategoryBrandFilterRequest{})
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())

		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	responseMap := make(map[string]interface{})
	responseMap["total"] = response.Total
	list := make([]interface{}, 0)
	for _, value := range response.Data {
		listMap := make(map[string]interface{})
		listMap["id"] = value.Id
		listMap["category"] = map[string]interface{}{
			"id":     value.Category.Id,
			"name":   value.Category.Name,
			"level":  value.Category.Level,
			"is_tab": value.Category.IsTab,
		}
		listMap["brand"] = map[string]interface{}{
			"id":   value.Brand.Id,
			"name": value.Brand.Name,
			"logo": value.Brand.Logo,
		}
		list = append(list, listMap)
	}
	responseMap["data"] = list
	ctx.JSON(http.StatusOK, responseMap)
	entry.Exit()
}

// Detail
// @Description: 根据id获取目录下的全部品牌
// @param ctx
//
func Detail(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())

		ctx.Status(http.StatusNotFound)
		return
	}
	response, err := global.GoodsClient.GetCategoryBrandList(context.WithValue(context.Background(), "ginContext", ctx), &proto.CategoryInfoRequest{
		Id: int32(idInt),
	})
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())

		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}

	result := make([]interface{}, 0)
	for _, value := range response.Data {
		responseMap := make(map[string]interface{})
		responseMap["id"] = value.Id
		responseMap["name"] = value.Name
		responseMap["logo"] = value.Logo
		result = append(result, responseMap)
	}
	ctx.JSON(http.StatusOK, result)
	entry.Exit()
}

// New
// @Description: 创建目录商品分类
// @param ctx
//
func New(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}
	categoryBrandForm := forms.CategoryBrandForm{}
	err := ctx.ShouldBind(&categoryBrandForm)
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())

		utils.HandleValidatorError(ctx, err)
		return
	}

	response, err := global.GoodsClient.CreateCategoryBrand(context.WithValue(context.Background(), "ginContext", ctx), &proto.CategoryBrandRequest{
		CategoryId: int32(categoryBrandForm.CategoryId),
		BrandId:    int32(categoryBrandForm.BrandId),
	})
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())

		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	responseMap := make(map[string]interface{})
	responseMap["id"] = response.Id
	responseMap["category"] = response.Category
	responseMap["brand"] = response.Brand

	ctx.JSON(http.StatusOK, responseMap)
	entry.Exit()
}

// Update
// @Description: 更新商品目录-品牌
// @param ctx
//
func Update(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}
	categoryBrandForm := forms.CategoryBrandForm{}
	err := ctx.ShouldBind(&categoryBrandForm)
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())

		utils.HandleValidatorError(ctx, err)
		return
	}

	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())

		ctx.Status(http.StatusNotFound)
		return
	}

	response, err := global.GoodsClient.UpdateCategoryBrand(context.WithValue(context.Background(), "ginContext", ctx), &proto.CategoryBrandRequest{
		Id:         int32(idInt),
		CategoryId: int32(categoryBrandForm.CategoryId),
		BrandId:    int32(categoryBrandForm.BrandId),
	})
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())

		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": response.Success,
	})
	entry.Exit()
}

// Delete
// @Description: 删除商品目录-品牌
// @param ctx
//
func Delete(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())

		ctx.Status(http.StatusNotFound)
		return
	}

	response, err := global.GoodsClient.DeleteCategoryBrand(context.WithValue(context.Background(), "ginContext", ctx), &proto.CategoryBrandRequest{Id: int32(idInt)})
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())

		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": response.Success,
	})
	entry.Exit()
}
