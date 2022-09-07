package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"order_web/global"
	"order_web/proto"
)

func InitService() {
	initGoodsService()
	initInventoryService()
}

func initGoodsService() {
	consulConfig := global.WebApiConfig.ConsulInfo

	goodsConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulConfig.Host, consulConfig.Port, global.WebApiConfig.GoodsService.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		zap.S().Fatalw("连接 【goods_service】商品服务失败", "err", err)
	}
	global.GoodsClient = proto.NewGoodsClient(goodsConn)
}

func initInventoryService() {
	consulConfig := global.WebApiConfig.ConsulInfo

	inventoryConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulConfig.Host, consulConfig.Port, global.WebApiConfig.InventoryService.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		zap.S().Fatalw("连接 【inventory_service】商品服务失败", "err", err)
	}
	global.InventoryClient = proto.NewInventoryClient(inventoryConn)
}
