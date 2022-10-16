package mode

import "userop_web/initialize"

func ReleaseMode() {
	//initialize.InitUseropService()
	initialize.InitUseropService()
	initialize.InitGoodsService()
}
