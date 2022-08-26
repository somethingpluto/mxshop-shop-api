package main

import (
	"flag"
	"go.uber.org/zap"
	"goods_api/global"
	"goods_api/initialize"
	runMode "goods_api/mode"
)

func main() {
	portF := flag.Int("port", 8022, "服务启动端口")
	modeF := flag.String("mode", "debug", "开发模式debug / 服务注册release")
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
	// 初始化RPC连接
	initialize.InitRPC()
	// 初始化路由
	Router := initialize.InitRouter()
	// 判断启动模式
	if mode == "" || mode != "debug" || mode != "release" {
		mode = global.WebApiConfig.Mode
	}
	zap.S().Warnf("启动模式：%s", mode)
	if mode == "release" {
		runMode.ReleaseMode()
	} else if mode == "debug" {
		runMode.DebugMode()
	}
	zap.S().Infof("goods_web服务开启 端口%s", port)
	err := Router.Run(":8022")
	if err != nil {
		panic(err)
	}
}
