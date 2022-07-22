package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"strconv"
	"time"
	"user_web/forms"
	"user_web/global/response"
	"user_web/proto"
	"user_web/utils"
)

var userClient proto.UserClient

func init() {
	// 1.grpc Dial
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", "127.0.0.1", 8000), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 【用户服务】失败", "msg", err.Error())
		return
	}
	// 2.实例化客户端
	userClient = proto.NewUserClient(userConn)
}

// GetUserList
// @Description: 获取用户列表
// @param c
//
func GetUserList(c *gin.Context) {
	// 0.从http请求中获取参数
	pageNum := c.DefaultQuery("page", "0")
	pageNumInt, _ := strconv.Atoi(pageNum)
	pageSize := c.DefaultQuery("size", "10")
	pageSizeInt, _ := strconv.Atoi(pageSize)
	// 1.调用rpc服务
	resp, err := userClient.GetUserList(context.Background(), &proto.PageInfoRequest{
		PageNum:  uint32(pageNumInt),
		PageSize: uint32(pageSizeInt),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList] 查询 【用户列表】 失败", "msg", err.Error())
		utils.HandleGrpcErrorToHttpError(err, c)
		return
	}
	// 2.返回查询结果
	result := make([]interface{}, 0)
	for _, value := range resp.Data {
		user := response.UserResponse{
			Id:       value.Id,
			Name:     value.NickName,
			Gender:   value.Gender,
			Mobile:   value.Mobile,
			Birthday: time.Time(time.Unix(int64(value.Birthday), 0)),
		}
		result = append(result, user)
	}
	c.JSON(http.StatusOK, result)
}

// PasswordLogin
// @Description: 手机密码登录
// @param c
//
func PasswordLogin(c *gin.Context) {
	// 1.实例化验证对象
	passwordLoginForm := forms.PasswordLoginForm{}
	// 2.判断是否有错误
	if err := c.ShouldBind(&passwordLoginForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	// 3.登录
	// 3.1获取用户加密后的密码
	userInfoResponse, err := userClient.GetUserByMobile(context.Background(), &proto.MobileRequest{Mobile: passwordLoginForm.Mobile})
	if err != nil {
		zap.S().Errorw("[GetUserByMobiles] 查询失败", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, c)
	}
	// 4.密码进行验证比对
	checkPasswordResponse, err := userClient.CheckPassword(context.Background(), &proto.CheckPasswordRequest{
		Password:          passwordLoginForm.Password,
		EncryptedPassword: userInfoResponse.Password,
	})
	if err != nil {
		zap.S().Errorw("[CheckPassword] 密码验证失败")
		//utils.HandleGrpcErrorToHttpError(err, c)
	}
	// 5.根据获取的结果返回
	if checkPasswordResponse.Success {
		c.JSON(http.StatusOK, gin.H{
			"msg": "登录成功",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "登录失败",
		})
	}
}
