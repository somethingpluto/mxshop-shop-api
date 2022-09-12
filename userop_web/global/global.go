package global

import (
	ut "github.com/go-playground/universal-translator"
	"userop_web/config"
	"userop_web/proto"
)

var (
	FilePath           *config.FilePathConfig
	NacosConfig        *config.NacosConfig
	WebApiConfig       *config.WebApiConfig
	Translator         ut.Translator
	UserFavoriteClient proto.UserFavoriteClient
	MessageClient      proto.MessageClient
	AddressClient      proto.AddressClient
)
