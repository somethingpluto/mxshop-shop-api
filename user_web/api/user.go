package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
	"user_web/forms"
	"user_web/global"
	"user_web/global/response"
	"user_web/proto"
	"user_web/utils"
)

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
	resp, err := global.UserClient.GetUserList(context.WithValue(context.Background(), "ginContext", c), &proto.PageInfoRequest{
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
	userInfoResponse, err := global.UserClient.GetUserByMobile(context.WithValue(context.Background(), "ginContext", c), &proto.MobileRequest{Mobile: passwordLoginForm.Mobile})
	if err != nil {
		zap.S().Errorw("[GetUserByMobiles] 查询失败", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, c)
	}
	// 4.密码进行验证比对
	checkPasswordResponse, err := global.UserClient.CheckPassword(context.WithValue(context.Background(), "ginContext", c), &proto.CheckPasswordRequest{
		Password:          passwordLoginForm.Password,
		EncryptedPassword: userInfoResponse.Password,
	})
	if err != nil {
		zap.S().Errorw("[CheckPassword] 密码验证失败")
		utils.HandleGrpcErrorToHttpError(err, c)
	}
	// 5.根据获取的结果返回
	if checkPasswordResponse.Success {
		token, err := utils.GenerateToken(uint(userInfoResponse.Id), userInfoResponse.NickName, uint(userInfoResponse.Role))
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

func Register(c *gin.Context) {
	// 1.表单认证
	registerForm := forms.RegisterForm{}
	err := c.ShouldBind(&registerForm)
	if err != nil {

		fmt.Println("c.ShouldBind error", err.Error())
		utils.HandleValidatorError(c, err)
		return
	}
	// 2.通过redis 验证 验证码是否正确
	connectRedis()
	value, err := red.Get(context.Background(), registerForm.Mobile).Result()
	if err == redis.Nil { // redis中没有验证码
		zap.S().Warnw("验证码发送/redis存储失败", "用户手机号", registerForm.Mobile)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
		return
	} else { // 验证码错误
		if value != registerForm.Code {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "验证码错误",
			})
			return
		}
	}
	userResponse, err := global.UserClient.CreateUser(context.Background(), &proto.CreateUserInfoRequest{
		NickName: registerForm.Mobile,
		Password: registerForm.Password,
		Mobile:   registerForm.Mobile,
	})
	if err != nil {
		zap.S().Errorw("[CreateUser] 失败", "err", err.Error())
		utils.HandleGrpcErrorToHttpError(err, c)
		return
	}
	token, err := utils.GenerateToken(uint(userResponse.Id), userResponse.NickName, uint(userResponse.Role))
	if err != nil {
		zap.S().Errorw("生成token失败", "err:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":        userResponse.Id,
		"nickName":  userResponse.NickName,
		"token":     token,
		"expiresAt": (time.Now().Unix() + 60*60*24*30) * 1000,
	})
}
