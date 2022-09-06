package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order_web/middlewares"
	"order_web/router"
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

	ApiGroup := Router.Group("/o/v1")
	router.InitOrderRouter(ApiGroup)
	router.InitShopCartRouter(ApiGroup)
	return Router
}
