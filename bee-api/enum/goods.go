package enum

type GoodsStatus int
type BeeShopGoodsAdditionType int

const (
	GoodsStatusUp   GoodsStatus = 0
	GoodsStatusDown GoodsStatus = 1

	BeeShopGoodsAdditionTypeSingle BeeShopGoodsAdditionType = 0 //单选
	BeeShopGoodsAdditionTypeMulti  BeeShopGoodsAdditionType = 1 //多选
)

var (
	GoodsStatusStrMap = map[GoodsStatus]string{
		GoodsStatusUp:   "上架",
		GoodsStatusDown: "下架",
	}
)
