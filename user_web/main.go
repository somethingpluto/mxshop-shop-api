package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	"user_web/initialize"
	"user_web/mode"
)

func main() {
	Port := flag.Int("port", 8022, "port: 服务端口(release模式下随机获取)")
	Mode := flag.String("mode", "release", "mode: 服务启动模式 debug 开发模式 / release 服务注册模式")
	flag.Parse()
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
	if *Mode == "release" {
		mode.ReleaseMode()
	} else if *Mode == "debug" {
		mode.DebugMode()
	}
	// 5.初始化router
	Router := initialize.InitRouters()
	zap.S().Infof("user_web服务开启")
	err := Router.Run(fmt.Sprintf(":%d", *Port))
	if err != nil {
		panic(err)
	}
}
