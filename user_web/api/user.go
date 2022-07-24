package api

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
	"user_web/forms"
	"user_web/global"
	"user_web/global/response"
	"user_web/middlewares"
	"user_web/models"
	"user_web/proto"
	"user_web/utils"
)

//var userClient proto.UserClient
//
//func init() {
//	// 1.grpc Dial
//	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", "127.0.0.1", 8000), grpc.WithTransportCredentials(insecure.NewCredentials()))
//	if err != nil {
//		zap.S().Errorw("[GetUserList] 连接 【用户服务】失败", "msg", err.Error())
//		return
//	}
//	// 2.实例化客户端
//	userClient = proto.NewUserClient(userConn)
//}

// GetUserList
// @Description: 获取用户列表
// @param c
//
func GetUserList(c *gin.Context) {
	// 0.从http请求中获取参数
	pageNum := c.DefaultQuery("page", "0")
	pageNumInt, _ := strconv.Atoi(pageNum)
	pageSize := c.DefaultQuery("size", "10")
	pageSizeInt, _ := strconv.Atoi(pageSize)
	// 1.调用rpc服务
	resp, err := global.UserClient.GetUserList(context.Background(), &proto.PageInfoRequest{
		PageNum:  uint32(pageNumInt),
		PageSize: uint32(pageSizeInt),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList] 查询 【用户列表】 失败", "msg", err.Error())
		utils.HandleGrpcErrorToHttpError(err, c)
		return
	}
	// 2.返回查询结果
	result := make([]interface{}, 0)
	for _, value := range resp.Data {
		user := response.UserResponse{
			Id:       value.Id,
			Name:     value.NickName,
			Gender:   value.Gender,
			Mobile:   value.Mobile,
			Birthday: time.Time(time.Unix(int64(value.Birthday), 0)),
		}
		result = append(result, user)
	}
	c.JSON(http.StatusOK, result)
}

// PasswordLogin
// @Description: 手机密码登录
// @param c
//
func PasswordLogin(c *gin.Context) {
	// 1.实例化验证对象
	passwordLoginForm := forms.PasswordLoginForm{}
	// 2.判断是否有错误
	if err := c.ShouldBind(&passwordLoginForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	verify := store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, true)
	if !verify {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
		return
	}

	// 3.登录
	// 3.1获取用户加密后的密码
	userInfoResponse, err := global.UserClient.GetUserByMobile(context.Background(), &proto.MobileRequest{Mobile: passwordLoginForm.Mobile})
	if err != nil {
		zap.S().Errorw("[GetUserByMobiles] 查询失败", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, c)
	}
	// 4.密码进行验证比对
	checkPasswordResponse, err := global.UserClient.CheckPassword(context.Background(), &proto.CheckPasswordRequest{
		Password:          passwordLoginForm.Password,
		EncryptedPassword: userInfoResponse.Password,
	})
	if err != nil {
		zap.S().Errorw("[CheckPassword] 密码验证失败")
		//utils.HandleGrpcErrorToHttpError(err, c)
	}
	// 5.根据获取的结果返回
	if checkPasswordResponse.Success {
		j := middlewares.NewJWT()
		claims := models.CustomClaims{
			ID:          uint(userInfoResponse.Id),
			NickName:    userInfoResponse.NickName,
			AuthorityId: uint(userInfoResponse.Role),
			StandardClaims: jwt.StandardClaims{
				NotBefore: time.Now().Unix(),               // 签名的生效时间
				ExpiresAt: time.Now().Unix() + 60*60*24*30, // 设置30天过期
				Issuer:    "pluto",
			},
		}
		token, err := j.CreateToken(claims)
		if err != nil {
			zap.S().Errorw("生成token失败", "err:", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "生成token失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"id":        userInfoResponse.Id,
			"nickName":  userInfoResponse.NickName,
			"token":     token,
			"expiresAt": (time.Now().Unix() + 60*60*24*30) * 1000,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "登录失败",
		})
	}
}
