package mode

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"user_web/global"
	"user_web/proto"
)

func DebugMode() {
	target := "http://127.0.0.1:8000"
	userConn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	global.UserClient = proto.NewUserClient(userConn)
	zap.S().Infof("RPC debugg模式 服务连接成功")
}
