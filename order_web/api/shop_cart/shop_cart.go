package shop_cart

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"order_web/forms"
	"order_web/global"
	"order_web/proto"
	"order_web/utils"
	"strconv"
)

// List
// @Description: 获取购物车内商品列表
// @param ctx
//
func List(ctx *gin.Context) {
	// 获取购物车商品
	userId, _ := ctx.Get("userId")
	response, err := global.OrderClient.CartItemList(context.Background(), &proto.UserInfo{Id: int32(userId.(uint))})
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())
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
	// 请求商品服务商品信息
	goodsListResponse, err := global.GoodsClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{Id: ids})
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	reMap := gin.H{
		"total": response.Total,
	}
	goodsList := make([]interface{}, 0)
	for _, item := range response.Data { // 遍历购物车 获取商品ID
		for _, good := range goodsListResponse.Data { // 遍历商品 获得商品的详细信息 将两者进行组装
			if good.Id == item.GoodsId {
				tmpMap := map[string]interface{}{}
				tmpMap["id"] = item.Id
				tmpMap["goods_id"] = item.GoodsId
				tmpMap["good_name"] = good.Name
				tmpMap["good_image"] = good.GoodsFrontImage
				tmpMap["good_price"] = good.ShopPrice
				tmpMap["nums"] = item.Nums
				tmpMap["checked"] = item.Checked

				goodsList = append(goodsList, tmpMap)
			}
		}
	}
	reMap["data"] = goodsList
	ctx.JSON(http.StatusOK, reMap)
}

// New
// @Description: 添加商品到购物车
// @param ctx
//
func New(ctx *gin.Context) {
	// 创建购物车表单
	itemForm := forms.ShopCartItemForm{}
	err := ctx.ShouldBind(&itemForm)
	if err != nil {
		zap.S().Errorw("Error", "message", "购物车表单验证失败", "err", err.Error())
		utils.HandleValidatorError(ctx, err)
		return
	}

	// 添加之前检查商品是否存在
	_, err = global.GoodsClient.GetGoodsDetail(context.Background(), &proto.GoodsInfoRequest{
		Id: itemForm.GoodsId,
	})
	if err != nil {
		zap.S().Errorw("Error")
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}

	// 如果添加到购物车数量和商品库存不一致
	inventoryResp, err := global.InventoryClient.InvDetail(context.Background(), &proto.GoodsInvInfo{
		GoodsId: itemForm.GoodsId,
	})
	if err != nil {
		zap.S().Errorw("Error", "message", "商品库存查询失败", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}

	if inventoryResp.Num < itemForm.Nums {
		ctx.JSON(http.StatusOK, gin.H{
			"nums": "商品库存不足",
		})
	}

	userId, _ := ctx.Get("userId")
	orderResp, err := global.OrderClient.CreateCartItem(context.Background(), &proto.CartItemRequest{
		GoodsId: itemForm.GoodsId,
		UserId:  int32(userId.(uint)),
	})
	if err != nil {
		zap.S().Errorw("Error", "message", "添加到购物车失败", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": orderResp.Id,
	})
}

// Update
// @Description: 更新购物车商品信息
// @param ctx
//
func Update(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		zap.S().Errorw("Error", "message", "param参数id转换失败", "err", err.Error())
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "url格式错误",
		})
		return
	}

	itemForm := forms.ShopCartItemUpdateForm{}
	err = ctx.ShouldBind(&itemForm)
	if err != nil {
		zap.S().Errorw("Error", "message", "updateForm表单验证失败", "err", err.Error())
		utils.HandleValidatorError(ctx, err)
		return
	}

	userId, _ := ctx.Get("userId")
	_, err = global.OrderClient.UpdateCartItem(context.Background(), &proto.CartItemRequest{

		UserId:  int32(userId.(uint)),
		GoodsId: int32(idInt),
		Nums:    itemForm.Nums,
		Checked: *itemForm.Checked,
	})
	if err != nil {
		zap.S().Errorw("Error", "message", "更新", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
}

// Delete
// @Description: 删除购物车内商品
// @param ctx
//
func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		zap.S().Errorw("Error", "message", "param参数id转换失败", "err", err.Error())
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "url格式错误",
		})
		return
	}

	userId, _ := ctx.Get("userId")
	_, err = global.OrderClient.DeleteCartItem(context.Background(), &proto.CartItemRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: int32(idInt),
	})
	if err != nil {
		zap.S().Errorw("Error", "message", "删除购物车记录失败", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	ctx.Status(http.StatusOK)
}
