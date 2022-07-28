package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"user_web/global"
	"user_web/proto"
)

// InitRPC
// @Description: 初始化rpc
//
func InitRPC() {
	cfg := api.DefaultConfig()
	consulInfo := global.ServerConfig.ConsulInfo
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)

	var userServiceHost string
	var userServicePort int
	client, err := api.NewClient(cfg)
	if err != nil {
		zap.S().Errorw("连接注册中心失败", "err", err.Error())
		return
	}
	data, err := client.Agent().ServicesWithFilter(`Service == "user-service"`)
	if err != nil {
		zap.S().Errorw("查询 user-service失败", "err", err.Error())
		return
	}
	for _, value := range data {
		userServiceHost = value.Address
		userServicePort = value.Port
		break
	}
	zap.S().Infof("查询到 user-service address:%s:%d", userServiceHost, userServicePort)

	target := fmt.Sprintf("%s:%d", userServiceHost, userServicePort)
	userConn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("RPC 服务连接失败")
		return
	}
	global.UserClient = proto.NewUserClient(userConn)
	zap.S().Infow("RPC 服务连接成功")
}
