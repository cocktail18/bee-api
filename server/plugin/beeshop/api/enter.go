package api

type ApiGroup struct {
	BeeUserMapperApi
	BeeShopGoodsCategoryApi
	BeeShopInfoApi
	BeeUserBalanceLogApi
	BeeCouponApi
	BeeUserCouponApi
	BeeCashLogApi
	BeeConfigApi
	BeeShopGoodsPropApi
	BeeShopGoodsAdditionApi
	BeeShopGoodsApi
	BeeShopGoodsImagesApi
	BeeShopGoodsSkuApi
	BeeShopGoodsContentApi
	BeeWxConfigApi
	BeeCommentApi
	BeeOrderApi
	BeeOrderReputationApi
	BeeBannerApi
	BeeUserAmountApi
	BeeCyTableApi
	BeeRegionApi
	BeeCmsInfoApi
	BeeUserAddressApi
	BeePeiSongApi
	BeeUserApi
	BeeOrderGoodsApi
	BeeOrderLogisticsApi
	BeeUploadFileApi
	BeePayLogApi
	BeeWxPayConfigApi
	BeeNoticeApi
	BeeLogisticsApi
	BeeDeliveryApi
	BeeOrderLogApi
	BeePrinterApi
	BeeQueueApi
	BeeUserQueueApi
}

var ApiGroupApp = new(ApiGroup)
