package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"time"
	"user_web/global"
	"user_web/global/response"
	"user_web/proto"
	"user_web/utils"
)

var userClient proto.UserClient

func init() {
	// 1.grpc Dial
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserService.Host, global.ServerConfig.UserService.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
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
	// 1.调用rpc服务
	resp, err := userClient.GetUserList(context.Background(), &proto.PageInfoRequest{
		PageNum:  0,
		PageSize: 10,
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

func GetUserByMobile(c *gin.Context) {
	// 1.调用rpc服务
	resp, err := userClient.GetUserByMobile(context.Background(), &proto.MobileRequest{Mobile: "1234561"})
	if err != nil {
		zap.S().Error("[GetUserByMobile] 查询 【用户】失败", "msg", err.Error())
		utils.HandleGrpcErrorToHttpError(err, c)
		return
	}

	// 2. 返回结果查询
	user := response.UserResponse{
		Id:       resp.Id,
		Name:     resp.NickName,
		Gender:   resp.Gender,
		Mobile:   resp.Mobile,
		Birthday: time.Time(time.Unix(int64(resp.Birthday), 0)),
	}
	c.JSON(http.StatusOK, user)

}
