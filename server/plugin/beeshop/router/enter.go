package router

type RouterGroup struct {
	BeeUserMapperRouter
	BeeShopGoodsCategoryRouter
	BeeShopInfoRouter
	BeeUserBalanceLogRouter
	BeeCouponRouter
	BeeUserCouponRouter
	BeeCashLogRouter
	BeeConfigRouter
	BeeShopGoodsPropRouter
	BeeShopGoodsAdditionRouter
	BeeShopGoodsRouter
	BeeShopGoodsImagesRouter
	BeeShopGoodsSkuRouter
	BeeShopGoodsContentRouter
	BeeWxConfigRouter
	BeeCommentRouter
	BeeOrderRouter
	BeeOrderReputationRouter
	BeeBannerRouter
	BeeUserAmountRouter
	BeeCyTableRouter
	BeeRegionRouter
	BeeCmsInfoRouter
	BeeUserAddressRouter
	BeePeiSongRouter
	BeeUserRouter
	BeeOrderGoodsRouter
	BeeOrderLogisticsRouter
	BeeUploadFileRouter
	BeePayLogRouter
	BeeWxPayConfigRouter
	BeeNoticeRouter
	BeeLogisticsRouter
	BeeOrderLogRouter
	BeePrinterRouter
}

var RouterGroupApp = new(RouterGroup)
