package global

import (
	ut "github.com/go-playground/universal-translator"
	"order_web/config"
	"order_web/proto"
)

var (
	FilePath     *config.FilePathConfig
	NacosConfig  *config.NacosConfig
	WebApiConfig *config.WebApiConfig
	Translator   ut.Translator
	OrderClient  proto.OrderClient
)
