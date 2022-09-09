package mode

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"user_web/global"
	"user_web/proto"
)

func ReleaseMode() {
	cfg := api.DefaultConfig()
	fmt.Println(cfg)
	consulConfig := global.WebServiceConfig.ConsulInfo
	cfg.Address = fmt.Sprintf("%s:%d", consulConfig.Host, consulConfig.Port)

	var userServiceHost string
	var userServicePort int
	client, err := api.NewClient(cfg)
	if err != nil {
		zap.S().Errorw("连接注册中心失败", "err", err.Error())
		return
	}
	data, err := client.Agent().ServicesWithFilter(`Service == "user_service"`)
	if err != nil {
		zap.S().Errorw("查询 user-service失败", "err", err.Error())
		return
	}
	for _, value := range data {
		userServiceHost = value.Address
		userServicePort = value.Port
		break
	}
	if userServiceHost == "" || userServicePort == 0 {
		zap.S().Fatal("InitRPC失败")
		return
	}
	zap.S().Infof("查询到user-service %s:%d", userServiceHost, userServicePort)
	target := fmt.Sprintf("%s:%d", userServiceHost, userServicePort)
	userConn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("grpc Dial错误", "err", err.Error())
		return
	}
	global.UserClient = proto.NewUserClient(userConn)
	zap.S().Infof("RPC release模式 服务连接成功")
}
