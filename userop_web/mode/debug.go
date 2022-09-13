package mode

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"userop_web/global"
	"userop_web/proto"
)

func DebugMode() {
	zap.S().Warnf("<<<<<debug模式仅启动Userop服务 关联服务未启动>>>>>")
	target := "127.0.0.1:8000"
	useropConn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	global.UserFavoriteClient = proto.NewUserFavoriteClient(useropConn)
	global.AddressClient = proto.NewAddressClient(useropConn)
	global.MessageClient = proto.NewMessageClient(useropConn)
	zap.S().Infof("RPC debugg模式 服务连接成功")
}
