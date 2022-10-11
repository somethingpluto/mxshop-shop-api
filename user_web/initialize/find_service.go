package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"user_web/global"
)

func InitUserService() string {
	cfg := api.DefaultConfig()
	consulConfig := global.WebServiceConfig.ConsulInfo
	cfg.Address = fmt.Sprintf("%s:%d", consulConfig.Host, consulConfig.Port)

	var userServiceHost string
	var userServicePort int
	client, err := api.NewClient(cfg)
	if err != nil {
		zap.S().Errorw("连接注册中心失败", "err", err.Error())
		return ""
	}
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == \"%s\"", global.WebServiceConfig.UserServiceInfo.Name))
	if err != nil {
		zap.S().Errorw("查询 user-service失败", "err", err.Error())
		return ""
	}
	for _, value := range data {
		userServiceHost = value.Address
		userServicePort = value.Port
		break
	}
	if userServiceHost == "" || userServicePort == 0 {
		zap.S().Fatal("InitRPC失败")
		return ""
	}
	zap.S().Infof("查询到user-service %s:%d", userServiceHost, userServicePort)
	target := fmt.Sprintf("%s:%d", userServiceHost, userServicePort)
	return target
}
