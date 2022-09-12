package initialize

import (
	"fmt"
	"path"
	"runtime"
	"userop_web/config"
	"userop_web/global"
)

func InitFileAbsPath() {
	basePath := getCurrentAbsolutePath()
	global.FilePath = &config.FilePathConfig{
		ConfigFile: basePath + "/config-debug.yaml",
		LogFile:    basePath + "/log",
	}
	fmt.Println("文件路径初始化成功:", basePath)
}

func getCurrentAbsolutePath() string {
	var abPath string
	_, fileName, _, ok := runtime.Caller(2)
	if ok {
		abPath = path.Dir(fileName)
	}
	return abPath
}
