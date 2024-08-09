package enum

type GoodsStatus int
type GoodsAfterSale int // 售后类型 1仅退款，2退款退货，3换货
type BeeShopGoodsAdditionType int

const (
	GoodsStatusUp   GoodsStatus = 0
	GoodsStatusDown GoodsStatus = 1

	BeeShopGoodsAdditionTypeSingle BeeShopGoodsAdditionType = 0 //单选
	BeeShopGoodsAdditionTypeMulti  BeeShopGoodsAdditionType = 1 //多选

	GoodsAfterSaleNone      GoodsAfterSale = 0
	GoodsAfterSaleRefund    GoodsAfterSale = 1
	GoodsAfterSaleRefundRet GoodsAfterSale = 2
	GoodsAfterSaleExchange  GoodsAfterSale = 3
)

var (
	GoodsStatusStrMap = map[GoodsStatus]string{
		GoodsStatusUp:   "上架",
		GoodsStatusDown: "下架",
	}

	GoodsAfterSaleStrMap = map[GoodsAfterSale]string{
		GoodsAfterSaleNone:      "无售后",
		GoodsAfterSaleRefund:    "仅退款",
		GoodsAfterSaleRefundRet: "退款退货",
		GoodsAfterSaleExchange:  "换货",
	}
)
