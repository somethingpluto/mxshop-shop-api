package shop_cart

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"order_web/global"
	"order_web/proto"
	"order_web/utils"
)

func List(ctx *gin.Context) {
	// 获取购物车商品
	userId, _ := ctx.Get("userId")
	response, err := global.OrderClient.CartItemList(context.Background(), &proto.UserInfo{Id: int32(userId.(uint))})
	if err != nil {
		zap.S().Errorf("[CartItemList] 查询失败")
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
		zap.S().Errorf("[CartItemList] 商品列表查询失败")
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	reMap := gin.H{
		"total": response.Total,
	}
	goodsList := make([]interface{}, 0)
	for _, item := range response.Data {
		for _, good := range goodsListResponse.Data {
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

func New(ctx *gin.Context) {

}

func Update(ctx *gin.Context) {

}

func Delete(ctx *gin.Context) {

}
