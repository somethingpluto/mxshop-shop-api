package main

import (
	"flag"
	"go.uber.org/zap"
	"order_web/initialize"
	"order_web/runMode"
)

func main() {
	portF := flag.Int("port", 8022, "服务启动端口")
	modeF := flag.String("runMode", "debug", "开发模式debug / 服务注册release")
	flag.Parse()
	port := *portF
	mode := *modeF
	// 初始化文件路径
	initialize.InitFileAbsPath()
	// 初始化配置
	initialize.InitConfig()
	// 初始化日志
	initialize.InitLogger()
	// 初始化翻译器
	initialize.InitTranslator("zh")
	initialize.InitService()
	if mode == "release" {
		runMode.ReleaseMode()
	} else if mode == "debug" {
		runMode.DebugMode()
	}
	Router := initialize.InitRouter()
	zap.S().Infof("goods_web服务开启 端口%s", port)
	err := Router.Run(":8022")
	if err != nil {
		panic(err)
	}
}
