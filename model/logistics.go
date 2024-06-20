package model

import (
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/enum"
	"github.com/shopspring/decimal"
)

type BeeLogistics struct {
	common.BaseModel
	Name       string                `gorm:"column:name;type:varchar(100);comment:模板名" json:"name"`
	FeeType    enum.BeeGoodsFeeType  `gorm:"column:fee_type;type:int(11);comment:计价方式，1按件，2按重量，3按金额" json:"feeType"`
	FeeTypeStr string                `gorm:"-" json:"feeTypeStr"`
	IsFree     bool                  `gorm:"column:is_free;type:tinyint(1);comment:是否包邮" json:"isFree"`
	Details    []*BeeLogisticsDetail `gorm:"-" json:"details"`
}

func (m *BeeLogistics) TableName() string {
	return "bee_logistics"
}

type BeeLogisticsDetail struct {
	common.BaseModel
	LogisticsId int64           `gorm:"column:logistics_id;type:bigint(20)" json:"logisticsId"`
	AddAmount   decimal.Decimal `gorm:"column:add_amount;type:decimal(20,2);comment:条件包邮,满xx包邮" json:"addAmount"`
	AddNumber   decimal.Decimal `gorm:"column:add_number;type:decimal(20,2);comment:条件包邮,满xx件包邮" json:"addNumber"`
	FirstAmount decimal.Decimal `gorm:"column:first_amount;type:decimal(20,2);comment:首件价格" json:"firstAmount"`
	FirstNumber decimal.Decimal `gorm:"column:first_number;type:decimal(20,2);comment:首件数量" json:"firstNumber"`
	Type        int             `gorm:"column:type;type:int(11);comment:类型，1首件，2续件" json:"type"`
}

func (m *BeeLogisticsDetail) TableName() string {
	return "bee_logistics_detail"
}

// BeeShopLogisticsFreeShipping 条件包邮配置
type BeeShopLogisticsFreeShipping struct {
	common.BaseModel
	CityId      string          `gorm:"column:city_id;type:varchar(100)" json:"cityId"`
	ProvinceId  string          `gorm:"column:province_id;type:varchar(100)" json:"provinceId"`
	Number      decimal.Decimal `gorm:"column:number;type:decimal(20,2)" json:"number"`
	LogisticsId int64           `gorm:"column:logistics_id;type:bigint(20)" json:"logisticsId"`
	Type        int             `gorm:"column:type;type:int(10);comment:0件，1重量，3金额" json:"type"`
}

func (m *BeeShopLogisticsFreeShipping) TableName() string {
	return "bee_shop_logistics_free_shipping"
}

type BeeShopLogisticsException struct {
	common.BaseModel
	LogisticsId int64           `gorm:"column:logistics_id;type:bigint(20);comment:模板详情id" json:"logisticsId"`
	DetailId    int64           `gorm:"column:detail_id;type:bigint(20);comment:模板详情id" json:"detailId"`
	ProvinceId  string          `gorm:"column:province_id;type:varchar(100)" json:"provinceId"`
	CityId      string          `gorm:"column:city_id;type:varchar(100)" json:"cityId"`
	RegionId    string          `gorm:"column:region_id;type:varchar(100)" json:"regionId"`
	RegionStr   string          `gorm:"column:region_str;type:varchar(100)" json:"regionStr"`
	AddAmount   decimal.Decimal `gorm:"column:add_amount;type:decimal(20,2)" json:"addAmount"`
	AddNumber   decimal.Decimal `gorm:"column:add_number;type:decimal(20,2)" json:"addNumber"`
	FirstAmount decimal.Decimal `gorm:"column:first_amount;type:decimal(20,2)" json:"firstAmount"`
	FirstNumber decimal.Decimal `gorm:"column:first_number;type:decimal(20,2)" json:"firstNumber"`
}

func (m *BeeShopLogisticsException) TableName() string {
	return "bee_shop_logistics_exception"
}
