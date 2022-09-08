package mode

import (
	"github.com/nacos-group/nacos-sdk-go/inner/uuid"
	"go.uber.org/zap"
	"goods_api/global"
	"goods_api/utils/register/consul"
)

func ReleaseMode() {
	consulInfo := global.WebApiConfig.ConsulInfo
	registryClient := consul.NewRegistry(consulInfo.Host, consulInfo.Port)
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	cfg := &consul.RegistryConfig{
		Address: "127.0.0.1",
		Port:    8022,
		Name:    global.WebApiConfig.Name,
		Tags:    consulInfo.Tags,
		Id:      id.String(),
	}
	err = registryClient.Register(cfg)
	if err != nil {
		zap.S().Errorw("服务注册失败", "err", err.Error())
		panic(err)
	}
}
