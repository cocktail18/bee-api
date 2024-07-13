package model

import (
	config2 "gitee.com/stuinfer/bee-api/config"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/logger"
)

var AllModel = []interface{}{
	&BeeUserAddress{},
	&BeeUserBalanceLog{},
	&BeeBanner{},
	&BeeCmsInfo{},
	&BeeComment{},
	&BeeConfig{},
	&BeeCoupon{},
	&BeeUserCoupon{},
	&BeeUserCouponLog{},
	&BeeCyTable{},
	&BeeUploadFile{},
	&BeePeiSong{},
	&BeeShopGoodsCategory{},
	&BeeShopGoods{},
	&BeeShopGoodsImages{},
	&BeeShopGoodsProp{},
	&BeeShopGoodsContent{},
	&BeeShopGoodsAddition{},
	&BeeLogistics{},
	&BeeNotice{},
	&BeeOrder{},
	&BeeOrderGoods{},
	&BeeOrderLogistics{},
	&BeeOrderReputation{},
	&BeeOrderQuDanHao{},
	&BeeOrderCoupon{},
	&BeeOrderLog{},
	&BeeOrderPeisong{},
	&BeePayLog{},
	//&BeeRegion{},
	&BeeShopInfo{},
	&BeeShopGoodsSku{},
	&BeeShoppingCart{},
	&BeeSignLog{},
	&BeeScoreLog{},
	&BeeUser{},
	&BeeLevel{},
	&BeeUserLevel{},
	&BeeUserMobile{},
	&BeeUserAmount{},
	&BeeUserMapper{},
	&BeeUserDynamicCode{},
	&BeeCashLog{},
	&BeeToken{},
	&BeeWxConfig{},
	&BeeWxPayConfig{},
	&BeeQueue{},
	&BeeUserQueue{},
}

func InitDB() {
	logger.GetLogger().Infof("开始初始化数据库，生产环境请禁用！！")
	if config2.AppConfigIns.DB.Drop {
		logger.GetLogger().Infof("清空数据库")
		if err := db.GetDB().Migrator().DropTable(AllModel...); err != nil {
			panic(err)
		}
		logger.GetLogger().Infof("清空数据库成功")
	}
	if err := db.GetDB().AutoMigrate(
		AllModel...,
	); err != nil {
		panic(err)
	}
	// 清空数据库
	logger.GetLogger().Infof("建表成功")

	logger.GetLogger().Infof("开始初始化demo数据")
	if err := InitDemoData(); err != nil {
		panic(err)
	}
	logger.GetLogger().Infof("初始化demo数据成功")
}

func InitDemoData() error {
	//db.GetDB().CreateOrder()
	return nil
}
