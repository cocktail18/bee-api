package model

import "gitee.com/stuinfer/bee-api/common"

// BeeCmsInfo
type BeeCmsInfo struct {
	common.BaseModel
	Content string `gorm:"column:content;type:varchar(5000);comment:内容" json:"content"`
	Key     string `gorm:"column:key;type:varchar(100);comment:key" json:"key"`
	Title   string `gorm:"column:title;type:varchar(100);comment:标题" json:"title"`
	Type    int64  `gorm:"column:type;type:bigint(11);comment:类型" json:"type"`
}

func (m *BeeCmsInfo) TableName() string {
	return "bee_cms_info"
}
