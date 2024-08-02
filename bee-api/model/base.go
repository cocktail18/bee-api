package model

import (
	"bufio"
	"errors"
	config2 "gitee.com/stuinfer/bee-api/config"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/logger"
	"go.uber.org/zap"
	"io"
	"log"
	"os"
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
	logger.GetLogger().Info("开始初始化数据库，生产环境请禁用！！")
	if config2.AppConfigIns.DB != nil && config2.AppConfigIns.DB.Drop {
		logger.GetLogger().Info("清空数据库")
		if err := db.GetDB().Migrator().DropTable(AllModel...); err != nil {
			panic(err)
		}
		logger.GetLogger().Info("清空数据库成功")
	}
	if err := db.GetDB().AutoMigrate(
		AllModel...,
	); err != nil {
		panic(err)
	}
	// 清空数据库
	logger.GetLogger().Info("建表成功")

	logger.GetLogger().Info("开始初始化demo数据")
	if err := InitDemoData(); err != nil {
		panic(err)
	}
	logger.GetLogger().Info("初始化demo数据成功")

	// 导入地址库
	initBeeRegion()
}

func InitDemoData() error {
	//db.GetDB().CreateOrder()
	return nil
}

func initBeeRegion() {
	// 打开文件
	file, err := os.Open("data/bee_region.sql")
	if err != nil { //文件不存在之类的
		logger.GetLogger().Error("打开地址库文件失败", zap.Error(err))
		return
	}
	defer file.Close()

	// 创建一个 Scanner 来读取文件内容
	scanner := bufio.NewScanner(file)
	logger.GetLogger().Info("导入地址库中，请勿关闭或者停止")
	// 循环读取每一行
	for scanner.Scan() {
		// 获取当前行的内容
		line := scanner.Text()
		if err := db.GetDB().Exec(line).Error; err != nil {
			logger.GetLogger().Error("写入地址库失败", zap.Error(err))
			continue
		}
	}

	// 检查 Scan 是否发生错误
	if err := scanner.Err(); err != nil && !errors.Is(err, io.EOF) {
		log.Fatal(err)
	}
	logger.GetLogger().Info("导入地址库成功.")
}
