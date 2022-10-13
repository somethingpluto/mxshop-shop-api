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
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}
	response, err := global.GoodsClient.BannerList(context.WithValue(context.Background(), "ginContext", ctx), &empty.Empty{})
	zap.S().Infof("List 触发 request:%v", ctx.Request.Host)
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())
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
	entry.Exit()
}

// New
// @Description: 创建轮播图
// @param ctx
//
func New(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}
	bannerForm := forms.BannerForm{}
	err := ctx.ShouldBind(&bannerForm)
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())
		utils.HandleValidatorError(ctx, err)
		return
	}

	response, err := global.GoodsClient.CreateBanner(context.WithValue(context.Background(), "ginContext", ctx), &proto.BannerRequest{
		Index: int32(bannerForm.Index),
		Image: bannerForm.Image,
		Url:   bannerForm.Url,
	})
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":    response.Id,
		"index": response.Index,
		"image": response.Image,
		"url":   response.Url,
	})
	entry.Exit()
}

// Update
// @Description: 更新轮播图信息
// @param ctx
//
func Update(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}
	bannerForm := forms.BannerForm{}
	err := ctx.ShouldBind(&bannerForm)
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

	response, err := global.GoodsClient.UpdateBanner(context.WithValue(context.Background(), "ginContext", ctx), &proto.BannerRequest{
		Id:    int32(idInt),
		Index: int32(bannerForm.Index),
		Url:   bannerForm.Url,
		Image: bannerForm.Image,
	})
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	entry.Exit()
	ctx.JSON(http.StatusOK, gin.H{
		"id":    response.Id,
		"index": response.Index,
		"url":   response.Url,
		"image": response.Image,
	})
	entry.Exit()
}

// Delete
// @Description: 删除轮播图
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

	response, err := global.GoodsClient.DeleteBanner(context.WithValue(context.Background(), "ginContext", ctx), &proto.BannerRequest{Id: int32(idInt)})
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
