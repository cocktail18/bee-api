package sys

import (
	_ "embed"
	"encoding/json"
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"github.com/samber/lo"
	"gorm.io/gorm/clause"
	"reflect"
	"time"
)

//go:embed demo/bee_config.json
var beeConfigStr string

//go:embed demo/bee_sub_shop.json
var beeSubShopStr string

//go:embed demo/bee_shop_category.json
var beeShopCategoryStr string

//go:embed demo/bee_notice.json
var beeNoticeStr string

//go:embed demo/bee_banner.json
var beeBannerStr string

//go:embed demo/bee_coupons.json
var beeCouponStr string

//go:embed demo/bee_shop_goods.json
var beeShopGoodsStr string

//go:embed demo/bee_shop_goods_detail.json
var beeShopGoodsDetailStr string

//go:embed demo/bee_shop_goods_addition.json
var beeShopGoodsAdditionStr string

func (srv *UserSrv) InitBeeData(userId int64) error {
	beeConfigArr := make([]*model.BeeConfig, 0)
	if err := initDbData(userId, beeConfigArr, beeConfigStr); err != nil {
		return err
	}

	beeShopArr := make([]*model.BeeShopInfo, 0)
	if err := initDbData(userId, beeShopArr, beeSubShopStr); err != nil {
		return err
	}

	beeCategoryArr := make([]*model.BeeShopGoodsCategory, 0)
	if err := initDbData(userId, beeCategoryArr, beeShopCategoryStr); err != nil {
		return err
	}

	beeNoticeArr := make([]*model.BeeNotice, 0)
	if err := initDbData(userId, beeNoticeArr, beeNoticeStr); err != nil {
		return err
	}

	beeBannerArr := make([]*model.BeeBanner, 0)
	if err := initDbData(userId, beeBannerArr, beeBannerStr); err != nil {
		return err
	}

	beeCouponArr := make([]*model.BeeCoupon, 0)
	if err := initDbData(userId, beeCouponArr, beeCouponStr); err != nil {
		return err
	}

	beeShopGoodsArr := make([]*model.BeeShopGoods, 0)
	if err := initDbData(userId, beeShopGoodsArr, beeShopGoodsStr); err != nil {
		return err
	}

	if err := initBeeShopGoodsDetail(userId); err != nil {
		return err
	}
	return nil
}

func initDbData[T any](userId int64, slices []T, dataStr string) error {
	if err := json.Unmarshal([]byte(dataStr), &slices); err != nil {
		return err
	}
	lo.ForEach(slices, func(item T, _ int) {
		userValue := reflect.ValueOf(item).Elem()          // 获取user对象的反射值
		userIdField := userValue.FieldByName("UserId")     // 根据字段名获取UserId字段的反射值
		if userIdField.IsValid() && userIdField.CanSet() { // 检查字段是否有效且可设置
			userIdField.SetInt(userId) // 使用SetInt方法设置字段的值
		}
	})
	if err := db.GetDB().CreateInBatches(slices, 100).Error; err != nil {
		return err
	}
	return nil
}

func initBeeShopGoodsDetail(userId int64) error {
	shopList := make([]*model.BeeShopInfo, 0)
	if err := db.GetDB().Find(&shopList).Limit(1).Error; err != nil {
		return err
	}
	beeShopGoodsDetailList := make([]*proto.GoodsDetailResp, 0)
	if err := json.Unmarshal([]byte(beeShopGoodsDetailStr), &beeShopGoodsDetailList); err != nil {
		return err
	}
	for _, beeShopGoodsDetail := range beeShopGoodsDetailList {
		baseInfo := beeShopGoodsDetail.BasicInfo
		baseInfo.ShopId = shopList[0].Id
		baseInfo.SellBeginTime = common.JsonTime(time.Now().AddDate(0, 0, -1))
		baseInfo.SellEndTime = common.JsonTime(time.Now().AddDate(10, 0, 0))
		if err := db.GetDB().Clauses(clause.OnConflict{UpdateAll: true}).Create(baseInfo).Error; err != nil {
			return err
		}
		contentModel := &model.BeeShopGoodsContent{
			BaseModel: *kit.GetInsertBaseModelWithUserId(userId),
			GoodsId:   baseInfo.Id,
			Content:   beeShopGoodsDetail.Content,
		}
		if err := db.GetDB().Create(contentModel).Error; err != nil {
			return err
		}
		lo.ForEach(beeShopGoodsDetail.Pics, func(item *model.BeeShopGoodsImages, _ int) {
			item.GoodsId = baseInfo.Id
		})
		if err := db.GetDB().Create(beeShopGoodsDetail.Pics).Error; err != nil {
			return err
		}

		//运费
		beeShopGoodsDetail.Logistics.Id = baseInfo.LogisticsId
		if err := db.GetDB().Clauses(clause.OnConflict{DoNothing: true}).Create(beeShopGoodsDetail.Logistics).Error; err != nil {
			return err
		}
		for _, logisticsDetail := range beeShopGoodsDetail.Logistics.Details {
			logisticsDetail.LogisticsId = beeShopGoodsDetail.Logistics.Id
			logisticsDetail.BaseModel = *kit.GetInsertBaseModelWithUserId(userId)
			if err := db.GetDB().Clauses(clause.OnConflict{DoNothing: true}).Create(logisticsDetail).Error; err != nil {
				return err
			}
		}

		for _, beeShopGoodsProp := range beeShopGoodsDetail.Properties {
			if err := initGoodsProp(userId, beeShopGoodsProp); err != nil {
				return err
			}
		}

		for _, sku := range beeShopGoodsDetail.SkuList {
			sku.PropertyChildIds = kit.GetDBPropertyChildIds(sku.PropertyChildIds)
			if err := db.GetDB().Clauses(clause.OnConflict{DoNothing: true}).Create(sku).Error; err != nil {
				return err
			}
		}

		var addition []*model.BeeShopGoodsAddition
		if err := json.Unmarshal([]byte(beeShopGoodsAdditionStr), &addition); err != nil {
			return err
		}
		for _, goodsAddition := range addition {
			if err := initGoodsAddition(userId, baseInfo.Id, goodsAddition); err != nil {
				return err
			}
		}
	}

	return nil
}

func initGoodsProp(userId int64, prop *model.BeeShopGoodsProp) error {
	prop.UserId = userId
	prop.DateAdd = common.JsonTime(time.Now())
	prop.DateUpdate = common.JsonTime(time.Now())
	if err := db.GetDB().Clauses(clause.OnConflict{DoNothing: true}).Create(prop).Error; err != nil {
		return err
	}
	for _, childProp := range prop.ChildsCurGoods {
		if err := initGoodsProp(userId, childProp); err != nil {
			return err
		}
	}
	return nil
}

func initGoodsAddition(userId int64, goodsId int64, prop *model.BeeShopGoodsAddition) error {
	prop.GoodsId = goodsId
	prop.BaseModel = *kit.GetInsertBaseModelWithUserId(userId)
	if err := db.GetDB().Clauses(clause.OnConflict{DoNothing: true}).Create(prop).Error; err != nil {
		return err
	}
	for _, childProp := range prop.Items {
		if err := initGoodsAddition(userId, goodsId, childProp); err != nil {
			return err
		}
	}
	return nil
}
