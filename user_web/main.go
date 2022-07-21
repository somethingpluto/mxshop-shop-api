package main

import (
	"fmt"
	"go.uber.org/zap"
	"user_web/global"
	"user_web/initialize"
)

func main() {
	// 0.初始化配置
	initialize.InitConfig()
	// 1.初始化logger
	initialize.InitLogger()
	// 2.初始化router
	Router := initialize.InitRouters()

	zap.S().Debugf("gin listen port %d", global.ServerConfig.UserServer.Port)
	err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.UserServer.Port))
	if err != nil {
		panic(err)
	}
}
