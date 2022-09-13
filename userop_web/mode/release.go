package mode

import "userop_web/initialize"

func ReleaseMode() {
	DebugMode()
	initialize.InitGoodsService()
}
