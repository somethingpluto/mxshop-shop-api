package main

import (
	"fmt"
	"go.uber.org/zap"
	"user_web/global"
	"user_web/initialize"
)

func main() {

	initialize.InitConfig()
	// 1.初始化logger
	initialize.InitLogger()
	// 2.初始化router
	Router := initialize.InitRouters()
	// 3.初始化 translator
	initialize.InitTranslator("zh")

	zap.S().Debugf("gin listen port %d", global.ServerConfig.UserServer.Port)
	err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.UserServer.Port))
	if err != nil {
		panic(err)
	}
}
