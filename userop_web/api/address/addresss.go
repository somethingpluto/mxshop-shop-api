package address

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"userop_web/forms"
	"userop_web/global"
	"userop_web/proto"
	"userop_web/utils"
)

func List(ctx *gin.Context) {

}

func New(ctx *gin.Context) {
	addressForm := forms.AddressForm{}
	err := ctx.ShouldBind(&addressForm)
	if err != nil {
		zap.S().Errorw("Error", "message", "添加地址表单验证失败")
		utils.HandleValidatorError(ctx, err)
		return
	}

	userId, _ := ctx.Get("userId")

	response, err := global.AddressClient.CreateAddress(context.Background(), &proto.AddressRequest{
		UserId:       int32(userId.(uint)),
		Province:     addressForm.Province,
		City:         addressForm.City,
		District:     addressForm.District,
		Address:      addressForm.Address,
		SignerName:   addressForm.SignerName,
		SignerMobile: addressForm.SignerMobile,
	})
	if err != nil {
		zap.S().Errorw("Error", "message", "添加地址服务失败", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":           response.Id,
		"province":     response.Province,
		"city":         response.City,
		"district":     response.District,
		"address":      response.Address,
		"singerName":   response.SignerName,
		"singerMobile": response.SignerMobile,
	})
}

func Delete(ctx *gin.Context) {

}

func Update(ctx *gin.Context) {

}
