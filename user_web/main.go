package main

import (
	"fmt"
	"go.uber.org/zap"
	"user_web/global"
	"user_web/initialize"
)

func main() {
	//0.初始化文件路径
	initialize.InitFilePath()
	// 1.初始化配置文件
	initialize.InitConfig()
	// 2.初始化日志器
	initialize.InitLogger()
	// 3.初始化翻译器
	initialize.InitTranslator("zh")
	// 4.初始化验证器
	initialize.InitValidator()
	// 初始化rpc
	initialize.InitRPC()
	// 5.初始化router
	Router := initialize.InitRouters()
	zap.S().Warnf("--------------user-web服务开启gin listen port %d", global.ServerConfig.UserServer.Port)
	err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.UserServer.Port))
	if err != nil {
		panic(err)
	}
}
