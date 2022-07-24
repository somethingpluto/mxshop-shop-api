package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"net/http"
)

var store = base64Captcha.DefaultMemStore

// GetCaptcha
// @Description: 生成图形验证码
// @param c
//
func GetCaptcha(c *gin.Context) {
	// 生成base64 数字验证码
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()
	if err != nil {
		zap.S().Errorw("生成验证码失败：", "err", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "生成验证码错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId": id,
		"picPath":   b64s,
	})
}
