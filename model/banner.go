package model

import "gitee.com/stuinfer/bee-api/common"

// bee_banner
type BeeBanner struct {
	common.BaseModel
	BusinessId int64  `gorm:"column:business_id;type:bigint(11);comment:业务id" json:"businessId"`
	LinkType   int64  `gorm:"column:link_type;type:bigint(11);comment:链接类型" json:"linkType"`
	Paixu      int64  `gorm:"column:paixu;type:bigint(11);comment:排序" json:"paixu"`
	PicUrl     string `gorm:"column:pic_url;type:varchar(100);comment:图片" json:"picUrl"`
	ShopId     int64  `gorm:"column:shop_id;type:bigint(11);comment:店铺id" json:"shopId"`
	Status     int64  `gorm:"column:status;type:bigint(11);comment:状态" json:"status"`
	Title      string `gorm:"column:title;type:varchar(100);comment:标题" json:"title"`
	Type       string `gorm:"column:type;type:varchar(100);comment:类型" json:"type"`
}

func (m *BeeBanner) TableName() string {
	return "bee_banner"
}
