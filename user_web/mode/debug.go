package mode

import (
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"user_web/global"
	"user_web/proto"
	"user_web/utils/otgrpc"
)

func DebugMode() {
	target := "127.0.0.1:8000"
	userConn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())))
	if err != nil {
		panic(err)
	}
	global.UserClient = proto.NewUserClient(userConn)
	zap.S().Infof("RPC debugg模式 服务连接成功")
}
