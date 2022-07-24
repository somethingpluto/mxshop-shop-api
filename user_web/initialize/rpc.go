package initialize

import (
	"fmt"
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
	host := global.ServerConfig.UserService.Host
	port := global.ServerConfig.UserService.Port
	target := fmt.Sprintf("%s:%d", host, port)
	userConn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("RPC 服务连接失败")
		return
	}
	global.UserClient = proto.NewUserClient(userConn)
	zap.S().Infow("RPC 服务连接成功")
}
