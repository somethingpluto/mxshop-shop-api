package brand

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
// @Description: 获取品牌列表
// @param ctx
//
func List(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}

	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)

	pSize := ctx.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)

	response, err := global.GoodsClient.BrandList(context.WithValue(context.Background(), "ginContext", ctx), &proto.BrandFilterRequest{
		Pages:       int32(pnInt),
		PagePerNums: int32(pSizeInt),
	})
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	responseMap := make(map[string]interface{})
	brandList := make([]interface{}, 0)
	responseMap["total"] = response.Total
	for _, value := range response.Data {
		brandList = append(brandList, map[string]interface{}{
			"id":   value.Id,
			"name": value.Name,
			"logo": value.Logo,
		})
	}
	responseMap["data"] = brandList
	ctx.JSON(http.StatusOK, responseMap)
	entry.Exit()
}

func New(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}

	brandForm := forms.BrandForm{}
	err := ctx.ShouldBind(&brandForm)
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())
		utils.HandleValidatorError(ctx, err)
		return
	}
	response, err := global.GoodsClient.CreateBrand(context.WithValue(context.Background(), "ginContext", ctx), &proto.BrandRequest{
		Name: brandForm.Name,
		Logo: brandForm.Logo,
	})
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	responseMap := make(map[string]interface{})
	responseMap["id"] = response.Id
	responseMap["name"] = response.Name
	responseMap["logo"] = response.Logo
	ctx.JSON(http.StatusOK, responseMap)
	entry.Exit()
}

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
	response, err := global.GoodsClient.DeleteBrand(context.WithValue(context.Background(), "ginContext", ctx), &proto.BrandRequest{
		Id: int32(idInt),
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

func Update(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}
	brandForm := forms.BrandForm{}
	err := ctx.ShouldBind(&brandForm)
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
	response, err := global.GoodsClient.UpdateBrand(context.WithValue(context.Background(), "ginContext", ctx), &proto.BrandRequest{
		Id:   int32(idInt),
		Name: brandForm.Name,
		Logo: brandForm.Logo,
	})
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())

		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":   response.Id,
		"name": response.Name,
		"logo": response.Logo,
	})
	entry.Exit()
}
