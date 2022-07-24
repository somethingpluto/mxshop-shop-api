package initialize

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"user_web/global"
)

// InitConfig
// @Description:  初始化配置
//
func InitConfig() {
	v := viper.New()
	// 文件路径设置
	v.SetConfigFile("./config-debug.yaml")
	err := v.ReadInConfig()
	if err != nil {
		zap.S().Errorw("读取config-debug.yaml配置文件失败")
		panic(err)
	}
	err = v.Unmarshal(global.ServerConfig)
	if err != nil {
		zap.S().Errorw("解析config-debug.yaml配置文件失败")
		panic(err)
	}
	zap.S().Infof("配置文件加载成功 config:%#v \n", global.ServerConfig)
}
