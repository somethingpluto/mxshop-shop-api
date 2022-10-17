package pay

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"go.uber.org/zap"
	"net/http"
	"order_web/global"
	"order_web/proto"
	"order_web/utils"
)

func Notify(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}
	aliPayInfo := global.WebServiceConfig.AlipayInfo
	client, err := alipay.New(aliPayInfo.AppID, aliPayInfo.PrivateKey, false)
	if err != nil {
		zap.S().Errorw("Error", "message", "支付宝支付实例化初始化失败", "err", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = client.LoadAliPayPublicKey(aliPayInfo.AliPublicKey)
	if err != nil {
		zap.S().Errorw("Error", "message", "加载支付宝公钥失败", "err", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	notification, err := client.GetTradeNotification(ctx.Request)
	if err != nil {
		zap.S().Errorw("Error", "message", "获取Notification失败", "err", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Println(notification)
	_, err = global.OrderClient.UpdateOrderStatus(context.WithValue(context.Background(), "ginContext", ctx), &proto.OrderStatus{
		OrderSn: notification.OutTradeNo,
		Status:  string(notification.TradeStatus),
	})
	if err != nil {
		zap.S().Errorw("Error", "message", "调用订单服务失败", "err", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.String(http.StatusOK, "success")
	entry.Exit()
}
