package model

import "gitee.com/stuinfer/bee-api/common"

// BeeUserAddress
type BeeUserAddress struct {
	common.BaseModel
	Address     string  `gorm:"column:address;type:varchar(100);comment:详细地址" json:"address"`
	AreaStr     string  `gorm:"column:area_str;type:varchar(100);comment:地区" json:"areaStr"`
	CityId      string  `gorm:"column:city_id;type:varchar(100);comment:地区码" json:"cityId"`
	CityStr     string  `gorm:"column:city_str;type:varchar(100);comment:城市" json:"cityStr"`
	DistrictId  string  `gorm:"column:district_id;type:varchar(100);comment:地区id" json:"districtId"`
	IsDefault   bool    `gorm:"column:is_default;type:tinyint(1);comment:默认地址" json:"isDefault"`
	Latitude    float64 `gorm:"column:latitude;type:double(9,6);comment:纬度" json:"latitude"`
	LinkMan     string  `gorm:"column:link_man;type:varchar(100);comment:联系人" json:"linkMan"`
	Longitude   float64 `gorm:"column:longitude;type:double(9,6);comment:经度" json:"longitude"`
	Mobile      string  `gorm:"column:mobile;type:varchar(100);comment:联系电话" json:"mobile"`
	ProvinceId  string  `gorm:"column:province_id;type:varchar(100);comment:省代码" json:"provinceId"`
	ProvinceStr string  `gorm:"column:province_str;type:varchar(100);comment:省份" json:"provinceStr"`
	Status      int64   `gorm:"column:status;type:bigint(11);comment:状态" json:"status"`
	Uid         int64   `gorm:"column:uid;type:bigint(11);comment:用户id" json:"uid"`
}

func (m *BeeUserAddress) TableName() string {
	return "bee_user_address"
}
