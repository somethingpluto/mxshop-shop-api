package router

import (
	"github.com/gin-gonic/gin"
	"userop_web/api/userFavorite"
	"userop_web/middlewares"
)

func InitUserFavoriteRouter(Router *gin.RouterGroup) {
	userFavoriteRouter := Router.Group("userfavs")
	{
		userFavoriteRouter.GET("", middlewares.JWTAuth(), userFavorite.List)
		userFavoriteRouter.POST("", middlewares.JWTAuth(), userFavorite.New)
		userFavoriteRouter.GET("/:id", middlewares.JWTAuth(), userFavorite.Detail)
		userFavoriteRouter.DELETE("/:id", middlewares.JWTAuth(), userFavorite.Delete)
	}
}
