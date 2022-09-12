package mode

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"userop_web/global"
	"userop_web/proto"
)

func ReleaseMode() {
	cfg := api.DefaultConfig()
	fmt.Println(cfg)
	consulConfig := global.WebApiConfig.ConsulInfo
	cfg.Address = fmt.Sprintf("%s:%d", consulConfig.Host, consulConfig.Port)

	var useropServiceHost string
	var useropServicePort int
	client, err := api.NewClient(cfg)
	if err != nil {
		zap.S().Errorw("连接注册中心失败", "err", err.Error())
		return
	}
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == \" %s\"", global.WebApiConfig.UseropService.Name))
	if err != nil {
		zap.S().Errorw("查询 userop_service失败", "err", err.Error())
		return
	}
	for _, value := range data {
		useropServiceHost = value.Address
		useropServicePort = value.Port
		break
	}
	if useropServiceHost == "" || useropServicePort == 0 {
		zap.S().Fatal("InitRPC失败")
		return
	}
	zap.S().Infof("查询到order-service %s:%d", useropServiceHost, useropServicePort)
	target := fmt.Sprintf("%s:%d", useropServiceHost, useropServicePort)
	useropConn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	if err != nil {
		zap.S().Errorw("grpc Dial错误", "err", err.Error())
		return
	}
	global.UserFavoriteClient = proto.NewUserFavoriteClient(useropConn)
	global.AddressClient = proto.NewAddressClient(useropConn)
	global.MessageClient = proto.NewMessageClient(useropConn)
	zap.S().Infof("RPC release模式 服务连接成功")
}
