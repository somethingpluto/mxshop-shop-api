package global

import (
	ut "github.com/go-playground/universal-translator"
	"user_web/config"
)

var (
	ServerConfig = &config.ServerConfig{}
	Translator   ut.Translator
)
