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
	BeeDeliveryRouter
	BeeRegionRouter
	BeeCmsInfoRouter
	BeeUserAddressRouter
	BeePeiSongRouter
	BeeUserRouter
	BeeOrderGoodsRouter
	BeeOrderLogisticsRouter
	BeeOrderLogRouter
	BeeOrderPeisongRouter
	BeeOrderPeisongLogRouter
	BeeUploadFileRouter
	BeePayLogRouter
	BeeWxPayConfigRouter
	BeeNoticeRouter
	BeeLogisticsRouter
	BeePrinterRouter
	BeeQueueRouter
	BeeUserQueueRouter
}

var RouterGroupApp = new(RouterGroup)
