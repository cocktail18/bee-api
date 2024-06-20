package model

import "gitee.com/stuinfer/bee-api/common"

// BeeRegion
// 参考：https://api.it120.cc/common/region/v2/province
// https://api.it120.cc/common/region/v2/child?pid=110100000000
type BeeRegion struct {
	common.BaseModel
	FirstLetter string `gorm:"column:first_letter;type:varchar(10)" json:"firstLetter"`
	Jianping    string `gorm:"column:jianping;type:varchar(100)" json:"jianping"`
	Level       int    `gorm:"column:level;type:int(11)" json:"level"`
	Name        string `gorm:"column:name;type:varchar(100)" json:"name"`
	NameEn      string `gorm:"column:name_en;type:varchar(200)" json:"nameEn"`
	Pingyin     string `gorm:"column:pingyin;type:varchar(100)" json:"pingyin"`
	Pid         string `gorm:"column:pid;type:varchar(100)" json:"pid"`
}

func (m *BeeRegion) TableName() string {
	return "bee_region"
}
