package proto

import (
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/util"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
	"strings"
	"time"
)

type ShoppingCartInfo struct {
	Number      int64                      `json:"number"`
	Price       decimal.Decimal            `json:"price"`
	Score       decimal.Decimal            `json:"score"`
	ShopList    []*ShoppingCartShop        `json:"shopList"`
	Items       []*ShoppingCartGoodsItem   `json:"items"`
	GoodsStatus []*ShoppingCartGoodsStatus `json:"goodsStatus"`
}

type ShoppingCartGoodsStatus struct {
	Id        int64            `json:"id"`
	SellEnd   bool             `json:"sellEnd"`
	SellStart bool             `json:"sellStart"`
	Status    enum.GoodsStatus `json:"status"`
	StatusStr string           `json:"statusStr"`
	Stores    int64            `json:"stores"`
}

func NewShoppingCartGoodsStatus(goods *model.BeeShopGoods) *ShoppingCartGoodsStatus {
	now := time.Now()
	sellEnd := util.IF(now.After(time.Time(goods.SellEndTime)), true, false)
	sellStart := util.IF(goods.Status == enum.GoodsStatusUp && now.After(time.Time(goods.SellBeginTime)), true, false)
	return &ShoppingCartGoodsStatus{
		Id:        goods.Id,
		SellEnd:   sellEnd,
		SellStart: sellStart,
		Status:    goods.Status,
		StatusStr: enum.GoodsStatusStrMap[goods.Status],
		Stores:    goods.Stores,
	}
}

type ShoppingCartGoodsItem struct {
	CategoryId   int64                   `json:"categoryId"`
	GoodsId      int64                   `json:"goodsId"`
	Key          string                  `json:"key"` // 唯一key
	LogisticsId  int64                   `json:"logisticsId"`
	MinBuyNumber int64                   `json:"minBuyNumber"`
	Name         string                  `json:"name"`
	Number       int64                   `json:"number"`
	Overseas     bool                    `json:"overseas"`
	Pic          string                  `json:"pic"`
	Price        decimal.Decimal         `json:"price"`
	Score        decimal.Decimal         `json:"score"`
	Selected     bool                    `json:"selected"`
	ShopId       int64                   `json:"shopId"`
	Sku          []*ShoppingCartGoodsSku `json:"sku"`
	Status       enum.CartStatus         `json:"status"`
	StatusStr    string                  `json:"statusStr"`
	Stores       int64                   `json:"stores"`
	Type         int64                   `json:"type"` //可不传该参数，用来区分多个购物车；业务场景中需要区分购物车的场景下使用
	Weight       decimal.Decimal         `json:"weight"`
}

type ShoppingCartGoodsSku struct {
	OptionId        int64  `json:"optionId"`
	OptionName      string `json:"optionName"`
	OptionValueId   int64  `json:"optionValueId"`
	OptionValueName string `json:"optionValueName"`
}

type ShoppingCartShop struct {
	Id              int64   `json:"id"`
	Name            string  `json:"name"`
	ServiceDistance float64 `json:"serviceDistance"`
}

func NewShoppingCartGoodsSkuList(sku *model.BeeShopGoodsSku) []*ShoppingCartGoodsSku {
	list := make([]*ShoppingCartGoodsSku, 0)
	childIdStrArr := strings.Split(sku.PropertyChildIds, ",")
	childNameStrArr := strings.Split(sku.PropertyChildNames, ",")
	for i := 0; i < len(childIdStrArr); i++ {
		if childIdStrArr[i] == "" {
			continue
		}
		childIds := strings.Split(childIdStrArr[i], ":")
		names := strings.Split(childNameStrArr[i], ":")
		list = append(list, &ShoppingCartGoodsSku{
			OptionId:        cast.ToInt64(childIds[0]),
			OptionName:      names[0],
			OptionValueId:   cast.ToInt64(childIds[1]),
			OptionValueName: names[1],
		})
	}
	return list
}
