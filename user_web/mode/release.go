package mode

import (
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"user_web/global"
	"user_web/initialize"
	"user_web/proto"
	"user_web/utils/otgrpc"
)

func ReleaseMode() {
	getUserService()
}

// getUserService
// @Description: 连接user_service
//
func getUserService() {
	target := initialize.InitUserService()
	if target == "" {
		panic("初始化用户服务失败")
	}
	userConn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())))
	if err != nil {
		zap.S().Errorw("grpc Dial错误", "err", err.Error())
		return
	}
	global.UserClient = proto.NewUserClient(userConn)
	zap.S().Infof("RPC release模式 服务连接成功")
}
