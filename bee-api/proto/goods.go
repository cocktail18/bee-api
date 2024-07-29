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
	Logistics  *GoodsLogistics             `json:"logistics"`
	Pics       []*model.BeeShopGoodsImages `json:"pics"`
	Pics2      []string                    `json:"pics2"`
	Properties []*model.BeeShopGoodsProp   `json:"properties"`
	SkuList    []*model.BeeShopGoodsSku    `json:"skuList"`
}

type GoodsLogistics struct {
	IsFree     bool                    `json:"isFree"`
	FeeType    enum.BeeLogisticsType   `json:"feeType"`
	FeeTypeStr string                  `json:"feeTypeStr"`
	Details    []*GoodsLogisticsDetail `json:"details"`
}

func (m *GoodsLogistics) FillData() {
	m.FeeTypeStr = enum.BeeLogisticsTypeMap[m.FeeType]
}

type GoodsLogisticsDetail struct {
	AddAmount   decimal.Decimal       `json:"addAmount"`
	AddNumber   decimal.Decimal       `json:"addNumber"`
	FirstAmount decimal.Decimal       `json:"firstAmount"`
	FirstNumber decimal.Decimal       `json:"firstNumber"`
	Type        enum.BeeLogisticsType `json:"type"`
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
