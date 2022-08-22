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

func New(ctx *gin.Context) {
	bannerForm := forms.BannerForm{}
	err := ctx.ShouldBind(bannerForm)
	if err != nil {
		utils.HandleGrpcErrorToHttpError(err, ctx)
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
