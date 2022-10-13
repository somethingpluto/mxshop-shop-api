package global

import (
	ut "github.com/go-playground/universal-translator"
	"goods_api/config"
	"goods_api/proto"
)

var (
	FilePath         *config.FilePathConfig
	NacosConfig      *config.NacosConfig
	WebServiceConfig *config.WebServiceConfig
	Translator       ut.Translator
	GoodsClient      proto.GoodsClient
)
