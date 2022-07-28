package initialize

import (
	"fmt"
	"path"
	"runtime"
	"user_web/config"
	"user_web/global"
)

func InitFilePath() {
	basePath := getCurrentAbPathByCaller()
	global.FileConfig = &config.FileConfig{
		ConfigFile: basePath + "/config-debug.yaml",
		LogFile:    basePath + "/log",
	}
	fmt.Println(global.FileConfig)
}

func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(2)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
