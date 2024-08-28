// 自动生成模板BeeOrderLogistics
package bee

import (
	"time"
)

// 用户订单地址 结构体  BeeOrderLogistics
type BeeOrderLogistics struct {
	Id          *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`                     //id字段
	UserId      *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`             //商店用户ID
	IsDeleted   *bool      `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`               //已删除
	DateAdd     *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`                    //创建时间
	DateUpdate  *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`           //更新时间
	DateDelete  *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`           //删除时间
	Address     string     `json:"address" form:"address" gorm:"column:address;comment:详细地址;size:100;"`            //详细地址
	AreaStr     string     `json:"areaStr" form:"areaStr" gorm:"column:area_str;comment:地区;size:100;"`             //地区
	CityId      string     `json:"cityId" form:"cityId" gorm:"column:city_id;comment:地区码;size:100;"`               //地区码
	CityStr     string     `json:"cityStr" form:"cityStr" gorm:"column:city_str;comment:城市;size:100;"`             //城市
	DistrictId  string     `json:"districtId" form:"districtId" gorm:"column:district_id;comment:地区id;size:100;"`  //地区id
	IsDefault   *bool      `json:"isDefault" form:"isDefault" gorm:"column:is_default;comment:默认地址;"`              //默认地址
	Latitude    *float64   `json:"latitude" form:"latitude" gorm:"column:latitude;comment:纬度;size:22;"`            //纬度
	LinkMan     string     `json:"linkMan" form:"linkMan" gorm:"column:link_man;comment:联系人;size:100;"`            //联系人
	Longitude   *float64   `json:"longitude" form:"longitude" gorm:"column:longitude;comment:经度;size:22;"`         //经度
	Mobile      string     `json:"mobile" form:"mobile" gorm:"column:mobile;comment:联系电话;size:100;"`               //联系电话
	ProvinceId  string     `json:"provinceId" form:"provinceId" gorm:"column:province_id;comment:省代码;size:100;"`   //省代码
	ProvinceStr string     `json:"provinceStr" form:"provinceStr" gorm:"column:province_str;comment:省份;size:100;"` //省份
	Status      *int       `json:"status" form:"status" gorm:"column:status;comment:状态;size:19;"`                  //状态
	Uid         *int       `json:"uid" form:"uid" gorm:"column:uid;comment:用户id;size:19;"`                         //用户id
	OrderId     *int       `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单id;size:19;"`            //订单id
	DadaShopNo  string     `json:"dadaShopNo" form:"dadaShopNo" gorm:"column:dada_shop_no;comment:达达店铺号;size:100;"`
	ShopId      *int       `json:"shopId" form:"shopId" gorm:"column:shop_id;comment:商店id;size:19;"`
}

// TableName 用户订单地址 BeeOrderLogistics自定义表名 bee_order_logistics
func (BeeOrderLogistics) TableName() string {
	return "bee_order_logistics"
}
