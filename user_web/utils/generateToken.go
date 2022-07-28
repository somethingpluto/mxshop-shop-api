package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"user_web/middlewares"
	"user_web/models"
)

// GenerateToken
// @Description: 生成Token
// @param Id
// @param NickName
// @param Role
// @return string
// @return error
//
func GenerateToken(Id uint, NickName string, Role uint) (string, error) {
	j := middlewares.NewJWT()
	claims := models.CustomClaims{
		ID:          Id,
		NickName:    NickName,
		AuthorityId: Role,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),               // 签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30, // 设置30天过期
			Issuer:    "pluto",
		},
	}
	token, err := j.CreateToken(claims)
	return token, err
}
