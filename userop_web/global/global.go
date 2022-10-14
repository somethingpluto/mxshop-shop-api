package global

import (
	ut "github.com/go-playground/universal-translator"
	"userop_web/config"
	"userop_web/proto"
)

var (
	FilePath           *config.FilePathConfig
	NacosConfig        *config.NacosConfig
	WebServiceConfig   *config.WebServiceConfig
	Translator         ut.Translator
	UserFavoriteClient proto.UserFavoriteClient
	MessageClient      proto.MessageClient
	AddressClient      proto.AddressClient
	GoodsClient        proto.GoodsClient
)
