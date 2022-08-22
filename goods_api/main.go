package main

import (
	"go.uber.org/zap"
	"goods_api/initialize"
)

func main() {
	initialize.InitFileAbsPath()
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitTranslator("zh")
	initialize.InitRPC()
	Router := initialize.InitRouter()
	zap.S().Infof("goods_web服务开启 端口8081")
	err := Router.Run(":8081")
	if err != nil {
		panic(err)
	}
}
