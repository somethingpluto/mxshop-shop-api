package mode

import "goods_api/initialize"

func ReleaseMode() {
	initialize.InitGoodsServiceConn()
}
