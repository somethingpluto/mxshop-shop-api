package goods

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goods_api/global"
	"goods_api/proto"
	"goods_api/utils"
	"net/http"
	"strconv"
)

// List
// @Description: 获取商品列表
// @param ctx
//
func List(ctx *gin.Context) {
	zap.S().Infof("goods List request:%v", ctx.Request.Host)
	request := &proto.GoodsFilterRequest{}

	priceMin := ctx.DefaultQuery("pmin", "0")
	priceMinInt, _ := strconv.Atoi(priceMin)
	request.PriceMin = int32(priceMinInt)

	priceMax := ctx.DefaultQuery("pmax", "0")
	priceMaxInt, _ := strconv.Atoi(priceMax)
	request.PriceMax = int32(priceMaxInt)

	isHot := ctx.DefaultQuery("ih", "0")
	if isHot == "1" {
		request.IsHot = true
	}
	isNew := ctx.DefaultQuery("in", "0")
	if isNew == "1" {
		request.IsNew = true
	}
	isTab := ctx.DefaultQuery("it", "0")
	if isTab == "1" {
		request.IsTab = true
	}
	categoryId := ctx.DefaultQuery("c", "0")
	categoryIdInt, _ := strconv.Atoi(categoryId)
	request.TopCategory = int32(categoryIdInt)

	pages := ctx.DefaultQuery("p", "0")
	pagesInt, _ := strconv.Atoi(pages)
	request.Pages = int32(pagesInt)

	perNums := ctx.DefaultQuery("pnum", "0")
	perNumsInt, _ := strconv.Atoi(perNums)
	request.PagePerNums = int32(perNumsInt)

	keywords := ctx.DefaultQuery("q", "")
	request.KeyWords = keywords

	brandId := ctx.DefaultQuery("b", "0")
	brandIdInt, _ := strconv.Atoi(brandId)
	request.Brand = int32(brandIdInt)

	response, err := global.GoodsClient.GoodsList(context.Background(), request)
	if err != nil {
		zap.S().Errorw("goods 查询商品列表失败", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	responseMap := map[string]interface{}{
		"total": response.Total,
	}
	goodsList := make([]interface{}, 0)
	for _, goods := range response.Data {
		goodsMap := map[string]interface{}{
			"id":          goods.Id,
			"name":        goods.Name,
			"goods_brief": goods.GoodsBrief,
			"desc":        goods.GoodsDesc,
			"ship_free":   goods.ShipFree,
			"desc_image":  goods.DescImages,
			"front_image": goods.GoodsFrontImage,
			"shop_price":  goods.ShopPrice,
			"category": map[string]interface{}{
				"id":   goods.Category.Id,
				"name": goods.Category.Name,
			},
			"brand": map[string]interface{}{
				"id":   goods.Brand.Id,
				"name": goods.Brand.Name,
				"logo": goods.Brand.Logo,
			},
			"is_host": goods.IsHot,
			"is_new":  goods.IsNew,
			"on_sale": goods.OnSale,
		}
		goodsList = append(goodsList, goodsMap)
	}
	responseMap["data"] = goodsList
	ctx.JSON(http.StatusOK, responseMap)
}
