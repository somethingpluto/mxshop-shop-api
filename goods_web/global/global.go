package global

import (
	ut "github.com/go-playground/universal-translator"
	"goods_api/config"
	"goods_api/proto"
)

var (
	FilePath     *config.FilePathConfig
	NacosConfig  *config.NacosConfig
	WebApiConfig *config.WebApiConfig
	Translator   ut.Translator
	GoodsClient  proto.GoodsClient
)
