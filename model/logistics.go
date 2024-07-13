package model

import (
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/enum"
	"github.com/shopspring/decimal"
)

type BeeLogistics struct {
	common.BaseModel
	Name                string                `gorm:"column:name;type:varchar(100);comment:模板名" json:"name"`
	FeeType             enum.BeeLogisticsType `gorm:"column:fee_type;type:int(11);comment:计价方式，1按件，2按重量，3按金额" json:"feeType"`
	FeeTypeStr          string                `gorm:"-" json:"feeTypeStr"`
	IsFree              bool                  `gorm:"column:is_free;type:tinyint(1);comment:是否包邮" json:"isFree"`
	DetailsJsonStr      string                `gorm:"column:details_json;type:text" json:"detailsJson"`
	FreeShippingSetting string                `gorm:"column:free_shipping_setting;type:text" json:"freeShippingSetting"`

	Details       []*BeeLogisticsDetail           `gorm:"-" json:"details"`
	FreeShippings []*BeeShopLogisticsFreeShipping `gorm:"-" json:"freeShipping"`
}

func (m *BeeLogistics) TableName() string {
	return "bee_logistics"
}

type BeeLogisticsDetail struct {
	Default     bool                  `json:"default,omitempty"`
	Remark      string                `json:"remark"`
	Names       []string              `json:"names" json:"names,omitempty"`
	CityIds     []string              `json:"cityIds" json:"cityIds,omitempty"`
	FirstNumber decimal.Decimal       `json:"firstNumber" json:"firstNumber,omitempty"`
	FirstAmount decimal.Decimal       `json:"firstAmount" json:"firstAmount"`
	AddNumber   decimal.Decimal       `json:"addNumber" json:"addNumber,omitempty"`
	AddAmount   decimal.Decimal       `json:"addAmount" json:"addAmount"`
	Type        enum.BeeLogisticsType `json:"type" json:"type,omitempty"`
	CityId      string                `json:"cityId" json:"cityId,omitempty"`
}

// BeeShopLogisticsFreeShipping 条件包邮配置
type BeeShopLogisticsFreeShipping struct {
	Default  bool                  `json:"default,omitempty"`
	Remark   string                `json:"remark"`
	Names    []string              `json:"names"`
	CityIds  []string              `json:"cityIds"`
	Type     enum.BeeLogisticsType `json:"type"`
	Number   decimal.Decimal       `json:"number"`
	RegionId string                `json:"regionId"`
}
