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

}

func New(ctx *gin.Context) {
	userFavFrom := forms.UserFavForm{}
	err := ctx.ShouldBind(&userFavFrom)
	if err != nil {
		zap.S().Errorw("Error", "message", "添加地址表单验证失败", "err", err.Error())
		utils.HandleValidatorError(ctx, err)
		return
	}

	userId, _ := ctx.Get("userId")
	_, err = global.UserFavoriteClient.AddUserFavorite(context.Background(), &proto.UserFavoriteRequest{
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
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	userId, _ := ctx.Get("userId")
	_, err = global.UserFavoriteClient.DeleteUserFavorite(context.Background(), &proto.UserFavoriteRequest{
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
}

func Detail(ctx *gin.Context) {

}
