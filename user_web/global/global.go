package global

import (
	ut "github.com/go-playground/universal-translator"
	"user_web/config"
	"user_web/proto"
)

var (
	ServerConfig = &config.ServerConfig{} // 服务配置
	Translator   ut.Translator            // 翻译器
	UserClient   proto.UserClient         // grpc客户端
	FileConfig   = &config.FileConfig{}   // 文件配置
)
