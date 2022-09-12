package validator

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"regexp"
)

// ValidateMobile
// @Description: 验证手机号是否符合规则
// @param fl
// @return bool
//
func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	ok, _ := regexp.MatchString(`^1(3[0-9]|5[0-3,5-9]|7[1-3,5-8]|8[0-9])\d{8}$`, mobile)
	if !ok {
		zap.S().Errorw("手机号验证失败", "手机号:", mobile)
		return false
	}
	return true
}
