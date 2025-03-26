package service

import (
	"context"
	"encoding/json"
	"fmt"
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"gitee.com/stuinfer/bee-api/util"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"sync"
	"time"
)

type ShoppingCartSrv struct {
	BaseSrv
}

var shoppingCartSrvOnce sync.Once
var shoppingCartSrvInstance *ShoppingCartSrv

func GetShoppingCartSrv() *ShoppingCartSrv {
	shoppingCartSrvOnce.Do(func() {
		shoppingCartSrvInstance = &ShoppingCartSrv{}
	})
	return shoppingCartSrvInstance
}

func (srv *ShoppingCartSrv) Add(c context.Context, goodsId int64, num int64, sku string, addition string) (*proto.ShoppingCartInfo, error) {
	uid := kit.GetUid(c)
	goodsSkuList := make([]*proto.ShoppingCartGoodsSku, 0, 10)
	err := json.Unmarshal([]byte(sku), &goodsSkuList)
	if err != nil {
		return nil, errors.Wrap(err, "解析sku字段失败")
	}
	propertyChildIds := ""
	for _, goodsSku := range goodsSkuList {
		propertyChildIds = propertyChildIds + fmt.Sprintf("%d:%d,", goodsSku.OptionId, goodsSku.OptionValueId)
	}
	goodsInfo, skuInfo, err := GetGoodsSrv().GetGoodsWithSku(c, goodsId, propertyChildIds)
	if err != nil {
		return nil, errors.Wrap(err, "获取商品信息失败")
	}
	// 获取已有购物车，检查是否有添加过该sku
	var curSkuInCart model.BeeShoppingCart
	err = db.GetDB().Where("uid = ? and sku_id = ?", uid, skuInfo.Id).Take(&curSkuInCart).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrap(err, "获取购物车信息失败")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) { // 新增
		now := time.Now()
		cartInfo := model.BeeShoppingCart{
			BaseModel: common.BaseModel{
				UserId:     kit.GetUserId(c),
				IsDeleted:  false,
				DateAdd:    common.JsonTime(now),
				DateUpdate: common.JsonTime(now),
			},
			Uid:              uid,
			GoodsId:          goodsId,
			CategoryId:       goodsInfo.CategoryId,
			Key:              util.GetUUIDStr(),
			LogisticsId:      0,
			Name:             goodsInfo.Name,
			Number:           num,
			Price:            skuInfo.Price,
			Selected:         true,
			ShopId:           goodsInfo.ShopId,
			SkuId:            skuInfo.Id,
			PropertyChildIds: kit.GetDBPropertyChildIds(propertyChildIds),
		}
		level, err := GetUserSrv().GetUserLevel(c, uid)
		if err != nil {
			level = 0
		}
		if level > 0 && skuInfo.VipPrice.GreaterThan(decimal.Zero) {
			cartInfo.Price = skuInfo.VipPrice
		}
		err = db.GetDB().Create(&cartInfo).Error
		if err != nil {
			return nil, errors.Wrap(err, "保存到数据库失败")
		}
		return srv.GetShoppingCart(c, uid)
	}

	return srv.ModifyNumber(c, uid, curSkuInCart.Key, curSkuInCart.Number+num)
}

func (srv *ShoppingCartSrv) GetShoppingCart(c context.Context, userId int64) (*proto.ShoppingCartInfo, error) {
	var cartList []*model.BeeShoppingCart
	err := db.GetDB().Where("uid = ? and is_deleted =0", userId).Find(&cartList).Error
	if err != nil {
		return nil, errors.Wrap(err, "获取购物车失败")
	}
	resp := &proto.ShoppingCartInfo{
		Number:      0,
		Price:       decimal.NewFromInt(0),
		Score:       decimal.NewFromInt(0),
		ShopList:    make([]*proto.ShoppingCartShop, len(cartList)),
		Items:       make([]*proto.ShoppingCartGoodsItem, len(cartList)),
		GoodsStatus: make([]*proto.ShoppingCartGoodsStatus, len(cartList)),
	}
	level, err := GetUserSrv().GetUserLevel(c, userId)
	if err != nil {
		level = 0
	}
	for i, _ := range cartList {
		cart := cartList[i]
		resp.Number = resp.Number + cart.Number
		resp.Price = resp.Price.Add(cart.Price.Mul(decimal.NewFromInt(cart.Number)))
		shopInfo, err := GetShopSrv().GetShopInfo(c, cart.ShopId, 0, 0)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			shopInfo = &model.BeeShopInfo{}
		} else if err != nil {
			return nil, err
		}
		resp.ShopList[i] = &proto.ShoppingCartShop{
			Id:              cart.ShopId,
			Name:            shopInfo.Name,
			ServiceDistance: shopInfo.ServiceDistance,
		}
		goodsInfo, skuInfo, err := GetGoodsSrv().GetGoodsWithSku(c, cart.GoodsId, cart.PropertyChildIds)
		if err != nil {
			return nil, err
		}
		if skuInfo == nil {
			skuInfo = &model.BeeShopGoodsSku{}
		}
		skuList := proto.NewShoppingCartGoodsSkuList(skuInfo)
		price := skuInfo.Price
		if level > 0 && skuInfo.VipPrice.GreaterThan(decimal.Zero) {
			price = skuInfo.VipPrice
		}
		resp.Items[i] = &proto.ShoppingCartGoodsItem{
			CategoryId:   cart.CategoryId,
			GoodsId:      cart.GoodsId,
			Key:          cart.Key,
			LogisticsId:  cart.LogisticsId,
			MinBuyNumber: goodsInfo.MinBuyNumber,
			Name:         goodsInfo.Name,
			Number:       cart.Number,
			Overseas:     goodsInfo.Overseas,
			Pic:          goodsInfo.Pic,
			Price:        price,
			Score:        skuInfo.Score,
			Selected:     cart.Selected,
			ShopId:       cart.ShopId,
			Sku:          skuList,
			Status:       enum.CartStatusNormal,
			StatusStr:    enum.CartStatusStrMap[enum.CartStatusNormal],
			Stores:       skuInfo.Stores,
			Type:         goodsInfo.Type,
			Weight:       skuInfo.Weight,
		}

		resp.GoodsStatus[i] = proto.NewShoppingCartGoodsStatus(goodsInfo)
	}
	return resp, err
}

func (srv *ShoppingCartSrv) ModifyNumber(c context.Context, userId int64, key string, number int64) (*proto.ShoppingCartInfo, error) {
	var info model.BeeShoppingCart
	err := db.GetDB().Where("uid = ? and `key` = ?", userId, key).Take(&info).Error
	if err != nil {
		return nil, errors.Wrap(err, "获取购物车失败")
	}
	info.Number = number
	info.DateUpdate = common.JsonTime(time.Now())
	err = db.GetDB().Where("uid = ? and `key` = ?", userId, key).Save(&info).Error
	if err != nil {
		return nil, errors.Wrap(err, "更新购物车失败")
	}
	return srv.GetShoppingCart(c, userId)
}

func (srv *ShoppingCartSrv) Remove(c context.Context, key string) (*proto.ShoppingCartInfo, error) {
	var info model.BeeShoppingCart
	var uid = kit.GetUid(c)
	err := db.GetDB().Where("uid = ? and `key` = ?", uid, key).Delete(&info).Error
	if err != nil {
		return nil, errors.Wrap(err, "删除购物车失败")
	}
	return srv.GetShoppingCart(c, uid)
}

func (srv *ShoppingCartSrv) Empty(c context.Context) (*proto.ShoppingCartInfo, error) {
	var info model.BeeShoppingCart
	var uid = kit.GetUid(c)
	err := db.GetDB().Where("uid = ?", uid).Delete(&info).Error
	if err != nil {
		return nil, errors.Wrap(err, "清空购物车失败")
	}
	return srv.GetShoppingCart(c, uid)
}
