package userFavorite

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"userop_web/forms"
	"userop_web/global"
	"userop_web/proto"
	"userop_web/utils"
)

func List(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}
	userId, _ := ctx.Get("userId")
	response, err := global.UserFavoriteClient.GetFavoriteList(context.WithValue(context.Background(), "ginContext", ctx), &proto.UserFavoriteRequest{
		UserId: int32(userId.(uint)),
	})
	if err != nil {
		zap.S().Errorw("Error", "message", "获取收藏列表失败", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	ids := make([]int32, 0)
	for _, item := range response.Data {
		ids = append(ids, item.GoodsId)
	}
	if len(ids) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"total": 0,
		})
		return
	}

	goodsResponse, err := global.GoodsClient.BatchGetGoods(context.WithValue(context.Background(), "ginContext", ctx), &proto.BatchGoodsIdInfo{Id: ids})
	if err != nil {
		zap.S().Errorw("Error", "message", "商品服务批量获取商品失败", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}

	responseMap := gin.H{
		"total": goodsResponse.Total,
	}
	goodsList := make([]interface{}, 0)
	for _, item := range response.Data {
		data := gin.H{
			"id": item.GoodsId,
		}

		for _, goods := range goodsResponse.Data {
			if item.GoodsId == goods.Id {
				data["name"] = goods.Name
				data["shop_price"] = goods.ShopPrice
			}
		}
		goodsList = append(goodsList, data)
	}
	responseMap["data"] = goodsList
	ctx.JSON(http.StatusOK, responseMap)
	entry.Exit()
}

func New(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}
	userFavFrom := forms.UserFavForm{}
	err := ctx.ShouldBind(&userFavFrom)
	if err != nil {
		zap.S().Errorw("Error", "message", "添加地址表单验证失败", "err", err.Error())
		utils.HandleValidatorError(ctx, err)
		return
	}

	userId, _ := ctx.Get("userId")
	_, err = global.UserFavoriteClient.AddUserFavorite(context.WithValue(context.Background(), "ginContext", ctx), &proto.UserFavoriteRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: userFavFrom.GoodsId,
	})
	if err != nil {
		zap.S().Errorw("Error", "message", "添加收藏失败", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "收藏成功",
	})
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
		ctx.Status(http.StatusNotFound)
		return
	}
	userId, _ := ctx.Get("userId")
	_, err = global.UserFavoriteClient.DeleteUserFavorite(context.WithValue(context.Background(), "ginContext", ctx), &proto.UserFavoriteRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: int32(idInt),
	})
	if err != nil {
		zap.S().Errorw("Error", "message", "删除收藏失败", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "删除收藏成功",
	})
	entry.Exit()
}

func Detail(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}
	goodsId := ctx.Param("id")
	goodsIdInt, err := strconv.ParseInt(goodsId, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	userId, _ := ctx.Get("userId")
	_, err = global.UserFavoriteClient.GetUserFavoriteDetail(context.WithValue(context.Background(), "ginContext", ctx), &proto.UserFavoriteRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: int32(goodsIdInt),
	})
	if err != nil {
		zap.S().Errorw("Error", "message", "获取收藏详情失败", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	ctx.Status(http.StatusOK)
	entry.Exit()
}
