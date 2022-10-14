package address

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"userop_web/forms"
	"userop_web/global"
	"userop_web/models"
	"userop_web/proto"
	"userop_web/utils"
)

func List(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}
	request := &proto.AddressRequest{}

	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)

	if currentUser.AuthorityId != 2 {
		userId, _ := ctx.Get("userId")
		request.UserId = int32(userId.(uint))
	}

	response, err := global.AddressClient.GetAddressList(context.WithValue(context.Background(), "ginContext", ctx), request)
	if err != nil {
		zap.S().Errorw("Error", "message", "获取地址列表失败", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}

	responseMap := gin.H{
		"total": response.Total,
	}

	result := make([]interface{}, 0)
	for _, item := range response.Data {
		temp := make(map[string]interface{})
		temp["id"] = item.Id
		temp["user_id"] = item.Id
		temp["province"] = item.Province
		temp["city"] = item.City
		temp["district"] = item.District
		temp["address"] = item.Address
		temp["singer_name"] = item.SignerName
		temp["singer_mobile"] = item.SignerMobile
		result = append(result, temp)
	}
	responseMap["data"] = result
	ctx.JSON(http.StatusOK, responseMap)
	entry.Exit()
}

func New(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}
	addressForm := forms.AddressForm{}
	err := ctx.ShouldBind(&addressForm)
	if err != nil {
		zap.S().Errorw("Error", "message", "添加地址表单验证失败")
		utils.HandleValidatorError(ctx, err)
		return
	}

	userId, _ := ctx.Get("userId")

	response, err := global.AddressClient.CreateAddress(context.WithValue(context.Background(), "ginContext", ctx), &proto.AddressRequest{
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

	_, err = global.AddressClient.DeleteAddress(context.WithValue(context.Background(), "ginContext", ctx), &proto.AddressRequest{Id: int32(idInt), UserId: int32(userId.(uint))})
	if err != nil {
		zap.S().Errorw("Error", "message", "删除地址服务失败", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
	entry.Exit()
}

func Update(ctx *gin.Context) {
	entry, blockError := utils.SentinelEntry(ctx)
	if blockError != nil {
		return
	}
	addressForm := forms.AddressForm{}
	err := ctx.ShouldBind(&addressForm)
	if err != nil {
		zap.S().Errorw("Error", "message", "地址表单验证失败", "err", err.Error())
		utils.HandleValidatorError(ctx, err)
		return
	}

	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	userId, _ := ctx.Get("userId")
	_, err = global.AddressClient.UpdateAddress(context.WithValue(context.Background(), "ginContext", ctx), &proto.AddressRequest{
		Id:           int32(idInt),
		UserId:       int32(userId.(uint)),
		Province:     addressForm.Province,
		City:         addressForm.City,
		District:     addressForm.District,
		Address:      addressForm.Address,
		SignerName:   addressForm.SignerName,
		SignerMobile: addressForm.SignerMobile,
	})
	if err != nil {
		zap.S().Errorw("Error", "message", "更新地址服务失败", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
	})
	entry.Exit()
}
