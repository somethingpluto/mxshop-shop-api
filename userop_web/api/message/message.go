package message

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"userop_web/forms"
	"userop_web/global"
	"userop_web/models"
	"userop_web/proto"
	"userop_web/utils"
)

func List(ctx *gin.Context) {
	request := &proto.MessageRequest{}
	userId, _ := ctx.Get("userId")
	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	if currentUser.AuthorityId == 1 {
		request.UserId = int32(userId.(uint))
	}

	response, err := global.MessageClient.MessageList(context.Background(), request)
	if err != nil {
		zap.S().Errorw("Error", "message", "查询message列表失败", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	responseMap := gin.H{
		"total": response.Total,
	}

	result := make([]interface{}, 0)
	for _, item := range response.Data {
		temp := map[string]interface{}{}
		temp["id"] = item.Id
		temp["user_id"] = item.UserId
		temp["type"] = item.MessageType
		temp["subject"] = item.Subject
		temp["message"] = item.Message
		temp["file"] = item.File
		result = append(result, temp)
	}
	responseMap["data"] = result
	ctx.JSON(http.StatusOK, responseMap)
}

func New(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")

	messageForm := forms.MessageForm{}
	err := ctx.ShouldBind(&messageForm)
	if err != nil {
		zap.S().Errorw("Error", "message", "新建信息表单验证失败", "err", err.Error())
		utils.HandleValidatorError(ctx, err)
		return
	}

	response, err := global.MessageClient.CreateMessage(context.Background(), &proto.MessageRequest{
		UserId:      int32(userId.(uint)),
		MessageType: messageForm.MessageType,
		Subject:     messageForm.Subject,
		Message:     messageForm.Message,
		File:        messageForm.File,
	})
	if err != nil {
		zap.S().Errorw("Error", "message", "新建message失败", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Id": response.Id,
	})
}
