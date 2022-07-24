package global

import (
	ut "github.com/go-playground/universal-translator"
	"user_web/config"
	"user_web/proto"
)

var (
	ServerConfig = &config.ServerConfig{}
	Translator   ut.Translator
	UserClient   proto.UserClient
)
