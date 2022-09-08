package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"goods_api/global"
	"goods_api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitRPC() {
	cfg := api.DefaultConfig()
	fmt.Println(cfg)
	consulConfig := global.WebApiConfig.ConsulInfo
	cfg.Address = fmt.Sprintf("%s:%d", consulConfig.Host, consulConfig.Port)

	var goodsServiceHost string
	var goodsServicePort int
	client, err := api.NewClient(cfg)
	if err != nil {
		zap.S().Errorw("连接注册中心失败", "err", err.Error())
		return
	}
	data, err := client.Agent().ServicesWithFilter(`Service == "goods_service"`)
	if err != nil {
		zap.S().Errorw("查询 goods-service失败", "err", err.Error())
		return
	}
	for _, value := range data {
		goodsServiceHost = value.Address
		goodsServicePort = value.Port
		break
	}
	if goodsServiceHost == "" || goodsServicePort == 0 {
		zap.S().Fatal("InitRPC失败")
		return
	}
	zap.S().Infof("查询到goods-service %s:%d", goodsServiceHost, goodsServicePort)
	target := fmt.Sprintf("%s:%d", goodsServiceHost, goodsServicePort)
	goodsConn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("grpc Dial错误", "err", err.Error())
		return
	}
	global.GoodsClient = proto.NewGoodsClient(goodsConn)
	zap.S().Infow("RPC 服务连接成功")
}
