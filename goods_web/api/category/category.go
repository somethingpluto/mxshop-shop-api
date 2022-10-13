package category

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"goods_api/forms"
	"goods_api/global"
	"goods_api/proto"
	"goods_api/utils"
	"net/http"
	"strconv"
)

// List
// @Description: 获取商品目录列表
// @param ctx
//
func List(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}
	response, err := global.GoodsClient.GetAllCategoriesList(context.WithValue(context.Background(), "ginContext", ctx), &empty.Empty{})
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())

		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	data := make([]interface{}, 0)
	err = json.Unmarshal([]byte(response.JsonData), &data)
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())

		zap.S().Errorw("Category 【list】获取商品分类失败", "err", err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
	entry.Exit()
}

// Detail
// @Description: 获取商品目录详情信息
// @param ctx
//
func Detail(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}

	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())

		ctx.Status(http.StatusNotFound)
		return
	}
	responseMap := make(map[string]interface{})
	subCategorys := make([]interface{}, 0)
	response, err := global.GoodsClient.GetSubCategory(context.WithValue(context.Background(), "ginContext", ctx), &proto.CategoryListRequest{
		Id: int32(i),
	})
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())

		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	for _, value := range response.SubCategory {
		subCategorys = append(subCategorys, map[string]interface{}{
			"id":              value.Id,
			"name":            value.Name,
			"level":           value.Level,
			"parent_category": value.ParentCategory,
			"is_tab":          value.IsTab,
		})
	}
	responseMap["id"] = response.Info.Id
	responseMap["name"] = response.Info.Name
	responseMap["level"] = response.Info.Level
	responseMap["parent_category"] = response.Info.ParentCategory
	responseMap["is_tab"] = response.Info.IsTab
	responseMap["sub_categorys"] = subCategorys
	ctx.JSON(http.StatusOK, responseMap)
	entry.Exit()
}

// New
// @Description: 创建商品目录
// @param ctx
//
func New(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}

	categoryForm := forms.CategoryForm{}
	err := ctx.ShouldBind(&categoryForm)
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())

		utils.HandleValidatorError(ctx, err)
		return
	}
	response, err := global.GoodsClient.CreateCategory(context.WithValue(context.Background(), "ginContext", ctx), &proto.CategoryInfoRequest{
		Name:           categoryForm.Name,
		ParentCategory: categoryForm.ParentCategory,
		Level:          categoryForm.Level,
		IsTab:          *categoryForm.IsTab,
	})
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())

		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	responseMap := map[string]interface{}{
		"id":     response.Id,
		"name":   response.Name,
		"parent": response.ParentCategory,
		"level":  response.Level,
		"is_tab": response.IsTab,
	}
	ctx.JSON(http.StatusOK, responseMap)
	entry.Exit()
}

func Delete(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}

	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())

		ctx.Status(http.StatusNotFound)
		return
	}
	response, err := global.GoodsClient.DeleteCategory(context.WithValue(context.Background(), "ginContext", ctx), &proto.DeleteCategoryRequest{Id: int32(i)})
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

// Update
// @Description: 更新目录信息
// @param ctx
//
func Update(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}

	categoryForm := forms.UpdateCategoryForm{}
	err := ctx.ShouldBind(&categoryForm)
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())

		utils.HandleValidatorError(ctx, err)
		return
	}
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())

		ctx.Status(http.StatusNotFound)
		return
	}
	request := &proto.CategoryInfoRequest{
		Id:   int32(i),
		Name: categoryForm.Name,
	}
	if categoryForm.IsTab != nil {
		request.Name = categoryForm.Name
		request.IsTab = *categoryForm.IsTab
	}
	response, err := global.GoodsClient.UpdateCategory(context.WithValue(context.Background(), "ginContext", ctx), request)
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())

		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":     response.Id,
		"name":   response.Name,
		"is_tab": response.IsTab,
		"level":  response.Level,
	})
	entry.Exit()
}
