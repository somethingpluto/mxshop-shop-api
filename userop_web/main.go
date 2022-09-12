package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	"userop_web/initialize"
	"userop_web/mode"
)

func main() {
	Port := flag.Int("port", 8022, "服务启动端口")
	Mode := flag.String("mode", "debug", "开发模式debug / 服务注册release")
	flag.Parse()
	port := *Port
	// 初始化文件路径
	initialize.InitFileAbsPath()
	// 初始化配置
	initialize.InitConfig()
	// 初始化日志
	initialize.InitLogger()
	// 初始化翻译器
	initialize.InitTranslator("zh")
	initialize.InitValidator()
	if *Mode == "release" {
		mode.ReleaseMode()
	} else if *Mode == "debug" {
		mode.DebugMode()
	}
	Router := initialize.InitRouter()
	zap.S().Infof("userop_web服务开启")
	err := Router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
}
