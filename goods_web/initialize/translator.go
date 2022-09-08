package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
	"goods_api/global"
	"reflect"
	"strings"
)

func InitTranslator(locale string) {
	err := initTrans(locale)
	if err != nil {
		zap.S().Errorw("初始化 翻译器失败")
		return
	}
	zap.S().Infow("翻译器加载成功")
}

func initTrans(locale string) (err error) {
	// 修改gin框架中的validator引擎属性 定制规则
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New()
		enT := en.New()
		// 第一个参数是备用语言环境 后面的参数是应该支持的语言环境
		universalTranslator := ut.New(enT, zhT, enT)
		global.Translator, ok = universalTranslator.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}

		switch locale {
		case "en":
			err := en_translations.RegisterDefaultTranslations(v, global.Translator)
			if err != nil {
				zap.S().Errorw("初始化 英文翻译器 失败")
				return err
			}
		case "zh":
			err := zh_translations.RegisterDefaultTranslations(v, global.Translator)
			if err != nil {
				zap.S().Errorw("初始化 中文翻译器 失败")
				return err
			}
		default:
			err := zh_translations.RegisterDefaultTranslations(v, global.Translator)
			if err != nil {
				zap.S().Errorw("初始化 中文翻译器 失败")
				return err
			}
		}
	}
	return
}
