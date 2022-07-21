package initialize

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

// InitLogger
// @Description: 初始化logger
//
func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
}

// getEncoder
// @Description: 配置编码
// @return zapcore.Encoder
//
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// getLogWriter
// @Description: 设置文件输出位置
// @return zapcore.WriteSyncer
//
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    1,
		MaxAge:     5,
		MaxBackups: 30,
		Compress:   false,
	}
	dest := io.MultiWriter(lumberJackLogger, os.Stdout)
	return zapcore.AddSync(dest)
}
