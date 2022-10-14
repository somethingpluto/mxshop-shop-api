package utils

import (
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/gin-gonic/gin"
	"net/http"
)

//TODO: Entry resource
func SentinelEntry(ctx *gin.Context) (*base.SentinelEntry, *base.BlockError) {
	entry, blockError := sentinel.Entry("goods_web", sentinel.WithTrafficType(base.Inbound))
	if blockError != nil {
		ctx.JSON(http.StatusTooManyRequests, gin.H{
			"msg": "请求过于频繁",
		})
	}
	return entry, blockError
}
