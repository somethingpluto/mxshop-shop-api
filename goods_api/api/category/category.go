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

func List(ctx *gin.Context) {
	response, err := global.GoodsClient.GetAllCategoriesList(context.Background(), &empty.Empty{})
	if err != nil {
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	data := make([]interface{}, 0)
	err = json.Unmarshal([]byte(response.JsonData), &data)
	if err != nil {
		zap.S().Errorw("Category 【list】获取商品分类失败", "err", err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func Detail(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	responseMap := make(map[string]interface{})
	subCategorys := make([]interface{}, 0)
	response, err := global.GoodsClient.GetSubCategory(context.Background(), &proto.CategoryListRequest{
		Id: int32(i),
	})
	if err != nil {
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
}

// TODO:返回数据无效
func New(ctx *gin.Context) {
	categoryForm := forms.CategoryForm{}
	err := ctx.ShouldBind(&categoryForm)
	if err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	response, err := global.GoodsClient.CreateCategory(context.Background(), &proto.CategoryInfoRequest{
		Name:           categoryForm.Name,
		ParentCategory: categoryForm.ParentCategory,
		Level:          categoryForm.Level,
		IsTab:          *categoryForm.IsTab,
	})
	if err != nil {
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
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	response, err := global.GoodsClient.DeleteCategory(context.Background(), &proto.DeleteCategoryRequest{Id: int32(i)})
	if err != nil {
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": response.Success,
	})
}

// Update
// @Description: 更新目录信息
// @param ctx
//
// TODO:目录信息不更新
func Update(ctx *gin.Context) {
	categoryForm := forms.UpdateCategoryForm{}
	err := ctx.ShouldBind(&categoryForm)
	if err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
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
	response, err := global.GoodsClient.UpdateCategory(context.Background(), request)
	if err != nil {
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":     response.Id,
		"name":   response.Name,
		"is_tab": response.IsTab,
		"level":  response.Level,
	})
}
