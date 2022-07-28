package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
	"user_web/global"
)

// HandleValidatorError
// @Description:
// @param c
// @param err
//
func HandleValidatorError(c *gin.Context, err error) {
	errors, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": "--" + err.Error(),
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errors.Translate(global.Translator)),
	})
}

// removeTopStruct
// @Description: 移除信息中的结构体前缀
// @param fields
// @return map[string]string
//
func removeTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for filed, err := range fields {
		rsp[filed[strings.Index(filed, ".")+1:]] = err
	}
	return rsp
}
