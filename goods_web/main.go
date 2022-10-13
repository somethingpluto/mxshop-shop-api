package main

import (
	"flag"
	"go.uber.org/zap"
	"goods_api/initialize"
	"goods_api/mode"
)

func main() {
	Port := flag.Int("port", 8022, "服务启动端口")
	Mode := flag.String("mode", "debug", "开发模式debug / 服务注册release")
	flag.Parse()
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
	initialize.InitSentinel()
	// 初始化路由
	Router := initialize.InitRouter()
	// 判断启动模式
	if *Mode == "debug" {
		zap.S().Warnf("debug本地调试模式 \n")
		mode.DebugMode()
	} else if *Mode == "release" {
		zap.S().Warnf("release服务注册模式 \n")
		mode.ReleaseMode()
	}
	zap.S().Infof("goods_web服务开启 端口%s", *Port)
	err := Router.Run(":8022")
	if err != nil {
		panic(err)
	}
}
