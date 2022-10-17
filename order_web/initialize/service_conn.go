package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"order_web/global"
	"order_web/proto"
	"order_web/utils/otgrpc"
)

func InitService() {
	initGoodsService()
	initInventoryService()
}

func initGoodsService() {
	consulConfig := global.WebServiceConfig.ConsulInfo
	goodsConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulConfig.Host, consulConfig.Port, global.WebServiceConfig.GoodsService.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		zap.S().Fatalw("连接 【goods_service】商品服务失败", "err", err)
	}
	zap.S().Infof("goods_service 连接成功")
	global.GoodsClient = proto.NewGoodsClient(goodsConn)
}

func initInventoryService() {
	consulConfig := global.WebServiceConfig.ConsulInfo

	inventoryConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulConfig.Host, consulConfig.Port, global.WebServiceConfig.InventoryService.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		zap.S().Fatalw("连接 【inventory_service】商品服务失败", "err", err)
	}
	zap.S().Infof("inventory_service 连接成功")
	global.InventoryClient = proto.NewInventoryClient(inventoryConn)
}
