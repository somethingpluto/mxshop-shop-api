package order

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"go.uber.org/zap"
	"net/http"
	"order_web/forms"
	"order_web/global"
	"order_web/models"
	"order_web/proto"
	"order_web/utils"
	"strconv"
)

func List(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}
	userId, _ := ctx.Get("userId")
	claims, _ := ctx.Get("claims")

	request := proto.OrderFilterRequest{}
	// 管理员特殊处理
	model := claims.(*models.CustomClaims)
	if model.AuthorityId == 1 {
		request.UserId = int32(userId.(uint))
	}

	pages := ctx.DefaultQuery("p", "0")
	pagesInt, _ := strconv.Atoi(pages)
	request.Pages = int32(pagesInt)

	perNums := ctx.DefaultQuery("pnum", "0")
	perNumsInt, _ := strconv.Atoi(perNums)
	request.PagePerNums = int32(perNumsInt)

	response, err := global.OrderClient.OrderList(context.WithValue(context.Background(), "ginContext", ctx), &request)
	if err != nil {
		zap.S().Errorw("Error", "message", "获取订单列表失败", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	responseMap := gin.H{
		// TODO:返回的总数为0
		"total": response.Total,
	}

	orderList := make([]interface{}, 0)
	for _, item := range response.Data {
		tempMap := map[string]interface{}{}

		tempMap["id"] = item.Id
		tempMap["status"] = item.Status
		tempMap["pay_type"] = item.PayType
		tempMap["user"] = item.UserId
		tempMap["post"] = item.Post
		tempMap["total"] = item.Total
		tempMap["address"] = item.Address
		tempMap["name"] = item.Name
		tempMap["mobile"] = item.Mobile
		tempMap["order_sn"] = item.OrderSn
		tempMap["id"] = item.Id
		tempMap["add_time"] = item.AddTime

		orderList = append(orderList, tempMap)
	}
	responseMap["data"] = orderList
	ctx.JSON(http.StatusOK, responseMap)
	entry.Exit()
}

func New(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}
	orderForm := forms.CreateOrderForm{}
	err := ctx.ShouldBind(&orderForm)
	if err != nil {
		zap.S().Errorw("Error", "message", "创建订单验证失败", "err", err.Error())
		utils.HandleValidatorError(ctx, err)
		return
	}
	userId, _ := ctx.Get("userId")
	response, err := global.OrderClient.CreateOrder(context.WithValue(context.Background(), "ginContext", ctx), &proto.OrderRequest{
		UserId:  int32(userId.(uint)),
		Address: orderForm.Address,
		Name:    orderForm.Name,
		Mobile:  orderForm.Mobile,
		Post:    orderForm.Post,
	})
	if err != nil {
		zap.S().Errorw("Error", "message", "创建订单服务失败", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	// 生成支付宝支付URL
	client, err := alipay.New(global.WebServiceConfig.AlipayInfo.AppID, global.WebServiceConfig.AlipayInfo.PrivateKey, false)
	if err != nil {
		zap.S().Errorw("Error", "message", "实例化支付宝失败", "err", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = client.LoadAliPayPublicKey(global.WebServiceConfig.AlipayInfo.AliPublicKey)
	if err != nil {
		zap.S().Errorw("Error", "message", "加载支付宝公钥失败", "err", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var p = alipay.TradePagePay{}
	p.NotifyURL = global.WebServiceConfig.AlipayInfo.NotifyURL
	p.ReturnURL = global.WebServiceConfig.AlipayInfo.ReturnURL
	p.Subject = "慕学生鲜订单-" + response.OrderSn
	p.OutTradeNo = response.OrderSn
	p.TotalAmount = strconv.FormatFloat(float64(response.Total), 'f', 2, 64)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err := client.TradePagePay(p)
	if err != nil {
		zap.S().Errorw("生成支付url失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         response.Id,
		"alipay_url": url.String(),
	})
	entry.Exit()
}

func Detail(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}
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

	request := proto.OrderRequest{
		Id: int32(idInt),
	}
	claims, _ := ctx.Get("claims")
	model := claims.(*models.CustomClaims)
	if model.AuthorityId == 1 {
		request.UserId = int32(userId.(uint))
	}
	response, err := global.OrderClient.OrderDetail(context.WithValue(context.Background(), "ginContext", ctx), &request)
	if err != nil {
		zap.S().Errorw("Error", "message", "获取用户详情失败", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	responseMap := gin.H{}
	responseMap["id"] = response.OrderInfo.Id
	responseMap["status"] = response.OrderInfo.Status
	responseMap["user"] = response.OrderInfo.UserId
	responseMap["post"] = response.OrderInfo.Post
	responseMap["total"] = response.OrderInfo.Total
	responseMap["address"] = response.OrderInfo.Address
	responseMap["name"] = response.OrderInfo.Name
	responseMap["mobile"] = response.OrderInfo.Mobile
	responseMap["pay_type"] = response.OrderInfo.PayType
	responseMap["order_sn"] = response.OrderInfo.OrderSn

	goodsList := make([]interface{}, 0)
	//TODO:goods列表为空
	for _, goods := range response.Goods {
		tempMap := gin.H{
			"id":    goods.GoodsId,
			"name":  goods.GoodsName,
			"image": goods.GoodsImage,
			"price": goods.GoodsPrice,
			"nums":  goods.Nums,
		}
		goodsList = append(goodsList, tempMap)
	}
	responseMap["goods"] = goodsList
	ctx.JSON(http.StatusOK, responseMap)
	entry.Exit()
}
