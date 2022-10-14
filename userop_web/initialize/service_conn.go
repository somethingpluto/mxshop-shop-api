package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"userop_web/global"
	"userop_web/proto"
)

func InitUseropService() {
	cfg := api.DefaultConfig()
	fmt.Println(cfg)
	consulConfig := global.WebServiceConfig.ConsulInfo
	cfg.Address = fmt.Sprintf("%s:%d", consulConfig.Host, consulConfig.Port)

	var useropServiceHost string
	var useropServicePort int
	client, err := api.NewClient(cfg)
	if err != nil {
		zap.S().Errorw("Error", "message", "注册中心连接失败", "err", err.Error())
		return
	}
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == \" %s\"", global.WebServiceConfig.UseropService.Name))
	if err != nil {
		zap.S().Errorw("Error", "message", "查找服务失败", "err", err.Error())
		return
	}
	for _, value := range data {
		useropServiceHost = value.Address
		useropServicePort = value.Port
		break
	}
	if useropServiceHost == "" || useropServicePort == 0 {
		zap.S().Errorw("Error", "message", "获取服务IP/Port失败")
		return
	}
	zap.S().Infof("获得【%s】服务 %s:%d", global.WebServiceConfig.UseropService.Name, useropServiceHost, useropServicePort)
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

func InitGoodsService() {
	cfg := api.DefaultConfig()
	fmt.Println(cfg)
	consulConfig := global.WebServiceConfig.ConsulInfo
	cfg.Address = fmt.Sprintf("%s:%d", consulConfig.Host, consulConfig.Port)

	var goodsServiceHost string
	var goodsServicePort int
	client, err := api.NewClient(cfg)
	if err != nil {
		zap.S().Errorw("Error", "message", "注册中心连接失败", "err", err.Error())
		return
	}
	data, err := client.Agent().ServicesWithFilter(`Service == goods_service`)
	if err != nil {
		zap.S().Errorw("Error", "message", "查找服务失败", "err", err.Error())
		return
	}
	for _, value := range data {
		goodsServiceHost = value.Address
		goodsServicePort = value.Port
		break
	}
	if goodsServiceHost == "" || goodsServicePort == 0 {
		zap.S().Errorw("Error", "message", "获取服务IP/Port失败")
		return
	}
	zap.S().Infof("获得【%s】服务 %s:%d", global.WebServiceConfig.UseropService.Name, goodsServiceHost, goodsServicePort)
	target := fmt.Sprintf("%s:%d", goodsServiceHost, goodsServicePort)
	goodsConn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	if err != nil {
		zap.S().Errorw("grpc Dial错误", "err", err.Error())
		return
	}
	global.GoodsClient = proto.NewGoodsClient(goodsConn)
	zap.S().Infof("RPC release模式 服务连接成功")
}
