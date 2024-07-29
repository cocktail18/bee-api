// 自动生成模板BeeBanner
package bee

import (
	"time"
)

// 商店banner 结构体  BeeBanner
type BeeBanner struct {
	Id         *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`                   //id字段
	UserId     *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`           //商店用户ID
	IsDeleted  *bool      `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`             //已删除
	DateAdd    *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`                  //创建时间
	DateUpdate *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`         //更新时间
	DateDelete *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`         //删除时间
	BusinessId *int       `json:"businessId" form:"businessId" gorm:"column:business_id;comment:业务id;size:19;"` //业务id
	LinkType   *int       `json:"linkType" form:"linkType" gorm:"column:link_type;comment:链接类型;size:19;"`       //链接类型
	Paixu      *int       `json:"paixu" form:"paixu" gorm:"column:paixu;comment:排序;size:19;"`                   //排序
	PicUrl     string     `json:"picUrl" form:"picUrl" gorm:"column:pic_url;comment:图片;size:100;"`              //图片
	ShopId     *int       `json:"shopId" form:"shopId" gorm:"column:shop_id;comment:店铺id;size:19;"`             //店铺id
	Status     *int       `json:"status" form:"status" gorm:"column:status;comment:状态;size:19;"`                //状态
	Title      string     `json:"title" form:"title" gorm:"column:title;comment:标题;size:100;"`                  //标题
	Type       string     `json:"type" form:"type" gorm:"column:type;comment:类型;size:100;"`                     //类型
}

// TableName 商店banner BeeBanner自定义表名 bee_banner
func (BeeBanner) TableName() string {
	return "bee_banner"
}
