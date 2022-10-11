package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"user_web/forms"
	"user_web/global"
	"user_web/utils"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	dysmsapi "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

var red *redis.Client

// SendNoteCode
// @Description: 发送验证码
// @param c
//
func SendNoteCode(c *gin.Context) {
	// 表单验证
	sendSmsForm := forms.SendSmsForm{}
	err := c.ShouldBind(&sendSmsForm)
	if err != nil {
		zap.S().Errorw("Error", "method", "SendNoteCode", "err", err.Error())
		utils.HandleValidatorError(c, err)
		return
	}

	config := sdk.NewConfig()
	credential := credentials.NewAccessKeyCredential(global.WebServiceConfig.AliSmsInfo.ApiKey, global.WebServiceConfig.AliSmsInfo.ApiSecret)
	/* use STS Token
	credential := credentials.NewStsTokenCredential("<your-access-key-id>", "<your-access-key-secret>", "<your-sts-token>")
	*/
	client, err := dysmsapi.NewClientWithOptions("cn-shenzhen", config, credential)
	if err != nil {
		panic(err)
	}
	smsCode := generateNoteCode(5)

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = global.WebServiceConfig.AliSmsInfo.SignName
	request.TemplateCode = global.WebServiceConfig.AliSmsInfo.TemplateCode
	request.PhoneNumbers = sendSmsForm.Mobile
	request.TemplateParam = "{\"code\":\"" + smsCode + "\"}"

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
	connectRedis()
	red.Set(context.WithValue(context.Background(), "ginContext", c), sendSmsForm.Mobile, smsCode, 300*time.Second)
	c.JSON(http.StatusOK, gin.H{
		"msg": "发送成功",
	})
}

// generateNoteCode
// @Description: 生成随机验证码
// @param width 验证码长度
// @return string
//
func generateNoteCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().Unix())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		_, _ = fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

// connectRedis
// @Description: 连接redis
//
func connectRedis() {
	red = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", global.WebServiceConfig.RedisInfo.Host, global.WebServiceConfig.RedisInfo.Port),
		Password: global.WebServiceConfig.RedisInfo.Password,
	})
}
