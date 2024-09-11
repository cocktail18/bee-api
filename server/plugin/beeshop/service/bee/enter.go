package bee

import (
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/sys"
	"golang.org/x/net/context"
)

type ServiceGroup struct {
	BeeUserMapperService
	BeeShopGoodsCategoryService
	BeeShopInfoService
	BeeUserBalanceLogService
	BeeCouponService
	BeeUserCouponService
	BeeCashLogService
	BeeConfigService
	BeeShopGoodsPropService
	BeeShopGoodsAdditionService
	BeeShopGoodsService
	BeeShopGoodsImagesService
	BeeShopGoodsSkuService
	BeeShopGoodsContentService
	BeeWxConfigService
	BeeCommentService
	BeeOrderService
	BeeOrderReputationService
	BeeBannerService
	BeeUserAmountService
	BeeCyTableService
	BeeDeliveryService
	BeeRegionService
	BeeCmsInfoService
	BeeUserAddressService
	BeePeiSongService
	BeeOrderPeisongLogService
	BeeUserService
	BeeOrderGoodsService
	BeeOrderLogisticsService
	BeeOrderPeisongService
	BeeUploadFileService
	BeePayLogService
	BeeWxPayConfigService
	BeeNoticeService
	BeeLogisticsService
	BeeOrderLogService
	BeePrinterService
	BeeQueueService
	BeeUserQueueService
	BeeRechargeSendRuleService
}

func getContext(shopUserId int) (context.Context, error) {
	ctx := context.Background()
	sysUser, err := sys.GetUserSrv().GetById(int64(shopUserId))
	if err != nil {
		return nil, err
	}
	ctx = kit.WithSysUser(ctx, sysUser)
	return ctx, nil
}
