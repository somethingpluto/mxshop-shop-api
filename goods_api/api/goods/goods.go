package goods

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
// @Description: 获取商品列表
// @param ctx
//
func List(ctx *gin.Context) {
	zap.S().Infof("goods 【List】获取商品列表 request:%v", ctx.Request.Host)
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

// New
// @Description: 创建商品
// @param ctx
//
func New(ctx *gin.Context) {
	zap.S().Infof("goods 【New】新建商品 request:%v", ctx.Request.Host)
	goodsForm := forms.GoodsForm{}
	if err := ctx.ShouldBindJSON(&goodsForm); err != nil {
		zap.S().Errorw("goods 创建商品格式错误", "err", err.Error())
		utils.HandleValidatorError(ctx, err)
		return
	}
	goodsClient := global.GoodsClient
	rsp, err := goodsClient.CreateGoods(context.Background(), &proto.CreateGoodsInfo{
		Name:            goodsForm.Name,
		GoodsSn:         goodsForm.GoodsSn,
		Stocks:          goodsForm.Stocks,
		MarketPrice:     goodsForm.MarketPrice,
		ShopPrice:       goodsForm.ShopPrice,
		GoodsBrief:      goodsForm.GoodsBrief,
		ShipFree:        *goodsForm.ShipFree,
		Images:          goodsForm.Images,
		DescImages:      goodsForm.DescImages,
		GoodsFrontImage: goodsForm.FrontImage,
		CategoryId:      goodsForm.CategoryId,
		BrandId:         goodsForm.Brand,
	})
	if err != nil {
		zap.S().Errorw("goods 创建商品失败", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, rsp)
}

// Detail
// @Description: 获取商品详情
// @param ctx
//
func Detail(ctx *gin.Context) {
	zap.S().Infof("goods 【Detail】获取商品详情 request:%v", ctx.Request.Host)
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	response, err := global.GoodsClient.GetGoodsDetail(context.WithValue(context.Background(), "ginContext", ctx), &proto.GoodsInfoRequest{Id: int32(i)})
	if err != nil {
		zap.S().Errorw("goods 获取商品详情失败", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	rep := map[string]interface{}{
		"id":          response.Id,
		"name":        response.Name,
		"goods_brief": response.GoodsBrief,
		"desc":        response.GoodsDesc,
		"ship_free":   response.ShipFree,
		"images":      response.Images,
		"desc_images": response.DescImages,
		"front_image": response.GoodsFrontImage,
		"shop_price":  response.ShopPrice,
		"category": map[string]interface{}{
			"id":   response.Category.Id,
			"name": response.Category.Name,
		},
		"brand": map[string]interface{}{
			"id":   response.Brand.Id,
			"name": response.Brand.Name,
			"logo": response.Brand.Logo,
		},
		"is_hot":  response.IsHot,
		"is_new":  response.IsNew,
		"on_sale": response.OnSale,
	}
	ctx.JSON(http.StatusOK, rep)
}

func Delete(ctx *gin.Context) {
	zap.S().Infof("goods 【Delelte】删除商品 request:%v", ctx.Request.Host)
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	response, err := global.GoodsClient.DeleteGoods(context.Background(), &proto.DeleteGoodsInfo{Id: int32(i)})
	if err != nil {
		zap.S().Errorf("goods 删除商品 id:%v", i)
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": response.Success,
	})
}

func UpdateStatus(ctx *gin.Context) {
	zap.S().Infof("goods 【Update】更新商品信息 request:%v", ctx.Request.Host)
	goodsStatusForm := forms.GoodsStatusForm{}
	err := ctx.ShouldBind(&goodsStatusForm)
	if err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	response, err := global.GoodsClient.UpdateGoods(context.Background(), &proto.CreateGoodsInfo{
		Id:     int32(i),
		IsHot:  *goodsStatusForm.IsHot,
		IsNew:  *goodsStatusForm.IsNew,
		OnSale: *goodsStatusForm.OnSale,
	})
	if err != nil {
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":          response.Id,
		"name":        response.Name,
		"goods_brief": response.GoodsBrief,
		"desc":        response.GoodsDesc,
		"ship_free":   response.ShipFree,
		"images":      response.Images,
		"desc_images": response.DescImages,
		"front_image": response.GoodsFrontImage,
		"shop_price":  response.ShopPrice,
		"category": map[string]interface{}{
			"id":   response.Category.Id,
			"name": response.Category.Name,
		},
		"brand": map[string]interface{}{
			"id":   response.Brand.Id,
			"name": response.Brand.Name,
			"logo": response.Brand.Logo,
		},
		"is_hot":  response.IsHot,
		"is_new":  response.IsNew,
		"on_sale": response.OnSale,
	})
}

func Update(ctx *gin.Context) {
	goodsForm := forms.GoodsForm{}
	err := ctx.ShouldBind(&goodsForm)
	if err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	response, err := global.GoodsClient.UpdateGoods(context.Background(), &proto.CreateGoodsInfo{
		Id:              int32(i),
		Name:            goodsForm.Name,
		GoodsSn:         goodsForm.GoodsSn,
		Stocks:          goodsForm.Stocks,
		MarketPrice:     goodsForm.MarketPrice,
		ShopPrice:       goodsForm.ShopPrice,
		GoodsBrief:      goodsForm.GoodsBrief,
		ShipFree:        *goodsForm.ShipFree,
		Images:          goodsForm.Images,
		DescImages:      goodsForm.DescImages,
		GoodsFrontImage: goodsForm.FrontImage,
		CategoryId:      goodsForm.CategoryId,
		BrandId:         goodsForm.Brand,
	})
	if err != nil {
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":          response.Id,
		"name":        response.Name,
		"goods_brief": response.GoodsBrief,
		"desc":        response.GoodsDesc,
		"ship_free":   response.ShipFree,
		"images":      response.Images,
		"desc_images": response.DescImages,
		"front_image": response.GoodsFrontImage,
		"shop_price":  response.ShopPrice,
		"category": map[string]interface{}{
			"id":   response.Category.Id,
			"name": response.Category.Name,
		},
		"brand": map[string]interface{}{
			"id":   response.Brand.Id,
			"name": response.Brand.Name,
			"logo": response.Brand.Logo,
		},
		"is_hot":  response.IsHot,
		"is_new":  response.IsNew,
		"on_sale": response.OnSale,
	})
}
