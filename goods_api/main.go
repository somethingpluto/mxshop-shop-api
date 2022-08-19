package main

import "goods_api/initialize"

func main() {
	initialize.InitFileAbsPath()
	initialize.InitLogger()
	initialize.InitConfig()
}
