package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"userop_web/middlewares"
	"userop_web/router"
)

func InitRouter() *gin.Engine {
	Router := gin.Default()
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})
	Router.Use(middlewares.Cors())

	ApiGroup := Router.Group("/up/v1")
	router.InitMessageRouter(ApiGroup)
	router.InitAddressRouter(ApiGroup)
	router.InitUserFavoriteRouter(ApiGroup)
	return Router
}
