package global

import (
	ut "github.com/go-playground/universal-translator"
	"order_web/config"
	"order_web/proto"
)

var (
	FilePath         *config.FilePathConfig
	NacosConfig      *config.NacosConfig
	WebServiceConfig *config.WebServiceConfig
	Translator       ut.Translator
	OrderClient      proto.OrderClient
	GoodsClient      proto.GoodsClient
	InventoryClient  proto.InventoryClient
)
