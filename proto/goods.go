package proto

import (
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/model"
	"github.com/shopspring/decimal"
)

type GoodsDetailResp struct {
	BasicInfo  *model.BeeShopGoods         `json:"basicInfo"`
	Category   *model.BeeShopGoodsCategory `json:"category"`
	Content    string                      `json:"content"`
	Logistics  *model.BeeLogistics         `json:"logistics"`
	Pics       []*model.BeeShopGoodsImages `json:"pics"`
	Pics2      []string                    `json:"pics2"`
	Properties []*model.BeeShopGoodsProp   `json:"properties"`
	SkuList    []*model.BeeShopGoodsSku    `json:"skuList"`
}

type GoodsPriceResp struct {
	FxType             enum.FxType     `json:"fxType"`
	GoodsId            int64           `json:"goodsId"`
	Id                 int64           `json:"id"`
	OriginalPrice      decimal.Decimal `json:"originalPrice"`
	PingtuanPrice      decimal.Decimal `json:"pingtuanPrice"`
	Price              decimal.Decimal `json:"price"`
	PropertyChildIds   string          `json:"propertyChildIds"`
	PropertyChildNames string          `json:"propertyChildNames"`
	Score              decimal.Decimal `json:"score"`
	Stores             int64           `json:"stores"`
}
