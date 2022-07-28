package middlewares

import (
	"github.com/gin-gonic/gin"
)

// AdminAuth
// @Description: 登录管理员权限验证
// @return gin.HandlerFunc
//
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//claims, _ := c.Get("claims")
		//currentUser := claims.(*models.CustomClaims)

		//if currentUser.AuthorityId != 2 {
		//	c.JSON(http.StatusForbidden, gin.H{
		//		"msg": "没有权限",
		//	})
		//	c.Abort()
		//	return
		//}
		c.Next()
	}
}
