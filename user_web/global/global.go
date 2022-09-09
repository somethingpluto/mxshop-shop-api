package global

import (
	ut "github.com/go-playground/universal-translator"
	"user_web/config"
	"user_web/proto"
)

var (
	Translator       ut.Translator      // 翻译器
	UserClient       proto.UserClient   // grpc客户端
	FileConfig       *config.FileConfig // 文件配置
	NacosConfig      *config.NacosConfig
	WebServiceConfig *config.WebServiceConfig
	Port             int
)
