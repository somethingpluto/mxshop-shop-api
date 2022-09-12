package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"userop_web/global"
	"userop_web/proto"
)

func InitService() {
	consulConfig := global.WebApiConfig.ConsulInfo
	useropConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulConfig.Host, consulConfig.Port, global.WebApiConfig.UseropService.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		zap.S().Fatalw("连接 【goods_service】商品服务失败", "err", err)
	}
	global.UserFavoriteClient = proto.NewUserFavoriteClient(useropConn)
	global.AddressClient = proto.NewAddressClient(useropConn)
	global.MessageClient = proto.NewMessageClient(useropConn)
}
