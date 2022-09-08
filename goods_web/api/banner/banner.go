package banner

import (
	"context"
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
	response, err := global.GoodsClient.BannerList(context.Background(), &empty.Empty{})
	zap.S().Infof("List 触发 request:%v", ctx.Request.Host)
	if err != nil {
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	result := make([]interface{}, 0)
	for _, value := range response.Data {
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["index"] = value.Index
		reMap["image"] = value.Image
		result = append(result, reMap)
	}
	ctx.JSON(http.StatusOK, result)
}

// New
// @Description: 创建轮播图
// @param ctx
//
func New(ctx *gin.Context) {
	bannerForm := forms.BannerForm{}
	err := ctx.ShouldBind(&bannerForm)
	if err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	response, err := global.GoodsClient.CreateBanner(context.Background(), &proto.BannerRequest{
		Index: int32(bannerForm.Index),
		Image: bannerForm.Image,
		Url:   bannerForm.Url,
	})
	if err != nil {
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":    response.Id,
		"index": response.Index,
		"image": response.Image,
		"url":   response.Url,
	})
}

// Update
// @Description: 更新轮播图信息
// @param ctx
//
func Update(ctx *gin.Context) {
	bannerForm := forms.BannerForm{}
	err := ctx.ShouldBind(&bannerForm)
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
	response, err := global.GoodsClient.UpdateBanner(context.Background(), &proto.BannerRequest{
		Id:    int32(idInt),
		Index: int32(bannerForm.Index),
		Url:   bannerForm.Url,
		Image: bannerForm.Image,
	})
	if err != nil {
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":    response.Id,
		"index": response.Index,
		"url":   response.Url,
		"image": response.Image,
	})
}

// Delete
// @Description: 删除轮播图
// @param ctx
//
func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	response, err := global.GoodsClient.DeleteBanner(context.Background(), &proto.BannerRequest{Id: int32(idInt)})
	if err != nil {
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": response.Success,
	})
}
