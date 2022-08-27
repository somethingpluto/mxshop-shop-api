package brand

import (
	"context"
	"github.com/gin-gonic/gin"
	"goods_api/forms"
	"goods_api/global"
	"goods_api/proto"
	"goods_api/utils"
	"net/http"
	"strconv"
)

func List(ctx *gin.Context) {
	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)

	pSize := ctx.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)

	response, err := global.GoodsClient.BrandList(context.Background(), &proto.BrandFilterRequest{
		Pages:       int32(pnInt),
		PagePerNums: int32(pSizeInt),
	})
	if err != nil {
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
}

func New(ctx *gin.Context) {
	brandForm := forms.BrandForm{}
	err := ctx.ShouldBind(&brandForm)
	if err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	response, err := global.GoodsClient.CreateBrand(context.Background(), &proto.BrandRequest{
		Name: brandForm.Name,
		Logo: brandForm.Logo,
	})
	if err != nil {
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	responseMap := make(map[string]interface{})
	responseMap["id"] = response.Id
	responseMap["name"] = response.Name
	responseMap["logo"] = response.Logo
	ctx.JSON(http.StatusOK, responseMap)
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	response, err := global.GoodsClient.DeleteBrand(context.Background(), &proto.BrandRequest{
		Id: int32(idInt),
	})
	if err != nil {
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": response.Success,
	})
}

func Update(ctx *gin.Context) {
	brandForm := forms.BrandForm{}
	err := ctx.ShouldBind(&brandForm)
	if err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	response, err := global.GoodsClient.UpdateBrand(context.Background(), &proto.BrandRequest{
		Id:   int32(idInt),
		Name: brandForm.Name,
		Logo: brandForm.Logo,
	})
	if err != nil {
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":   response.Id,
		"name": response.Name,
		"logo": response.Logo,
	})
}
