package initialize

import (
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"order_web/global"
	myValidate "order_web/validator"
)

var validate *validator.Validate
var ok bool

func InitValidator() {
	validate, ok = binding.Validator.Engine().(*validator.Validate)
	if !ok {
		zap.S().Errorw("绑定自定义验证器失败")
	}
	initMobileValidator()
}

func initMobileValidator() {
	err := validate.RegisterValidation("mobile", myValidate.ValidateMobile)
	if err != nil {
		zap.S().Errorw("mobile 验证绑定失败", "错误", err.Error())
	}
	err = validate.RegisterTranslation("mobile", global.Translator, func(ut ut.Translator) error {
		return ut.Add("mobile", "{0} 非法的手机号码", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, err := ut.T("mobile", fe.Field())
		if err != nil {
			zap.S().Errorw("mobile 错误 翻译注册失败", "err", err.Error())
		}
		return t
	})
	if err != nil {
		zap.S().Errorw("mobile 注册翻译失败", "err", err.Error())
	}
	zap.S().Infow("mobile 验证器加载成功")
}
