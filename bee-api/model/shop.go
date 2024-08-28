package model

import (
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/enum"
	"github.com/golang/geo/s2"
	"github.com/shopspring/decimal"
)

// shop_info
type BeeShopInfo struct {
	common.BaseModel
	Address              string             `gorm:"column:address;type:varchar(100);comment:地址" json:"address"`
	BindUid              int64              `gorm:"column:bind_uid;type:bigint(11);comment:绑定用户id" json:"bindUid"`
	CityId               string             `gorm:"column:city_id;type:varchar(100);comment:城市id" json:"cityId"`
	CyTablePayMod        enum.CyTablePayMod `gorm:"column:cy_table_pay_mod;type:bigint(11);comment:支付方式" json:"cyTablePayMod"`
	DistrictId           string             `gorm:"column:district_id;type:varchar(100);comment:地区id" json:"districtId"`
	ExpressType          string             `gorm:"column:express_type;type:varchar(100);comment:快递类型" json:"expressType"`
	GoodsNeedCheck       bool               `gorm:"column:goods_need_check;type:tinyint(1);comment:接单需商家确认" json:"goodsNeedCheck"`
	Latitude             float64            `gorm:"column:latitude;type:double(9,6);comment:纬度" json:"latitude"`
	LinkPhone            string             `gorm:"column:link_phone;type:varchar(100);comment:联系人手机号" json:"linkPhone"`
	LinkMan              string             `gorm:"column:link_man;type:varchar(100);comment:联系人" json:"linkMan"`
	Longitude            float64            `gorm:"column:longitude;type:double(9,6);comment:经度" json:"longitude"`
	Name                 string             `gorm:"column:name;type:varchar(100);comment:店名" json:"name"`
	Number               string             `gorm:"column:number;type:varchar(100);comment:座机号码" json:"number"`
	NumberFav            int64              `gorm:"column:number_fav;type:bigint(11);comment:喜欢人数" json:"numberFav"`
	NumberGoodReputation int64              `gorm:"column:number_good_reputation;type:bigint(11);comment:商品评分" json:"numberGoodReputation"`
	NumberOrder          int64              `gorm:"column:number_order;type:bigint(11);comment:下单数" json:"numberOrder"`
	NumberReputation     int64              `gorm:"column:number_reputation;type:bigint(11);comment:店铺评分" json:"numberReputation"`
	OpenScan             bool               `gorm:"column:open_scan;type:tinyint(1);comment:开启点餐" json:"openScan"`
	OpenWaimai           bool               `gorm:"column:open_waimai;type:tinyint(1);comment:开启外卖" json:"openWaimai"`
	OpenZiqu             bool               `gorm:"column:open_ziqu;type:tinyint(1);comment:开启自取" json:"openZiqu"`
	OpeningHours         string             `gorm:"column:opening_hours;type:varchar(100);comment:开店时间" json:"openingHours"`
	Paixu                int64              `gorm:"column:paixu;type:bigint(11);comment:排序" json:"paixu"`
	ProvinceId           string             `gorm:"column:province_id;type:varchar(100);comment:省id" json:"provinceId"`
	RecommendStatus      int64              `gorm:"column:recommend_status;type:bigint(11);comment:推荐状态" json:"recommendStatus"`
	ServiceDistance      float64            `gorm:"column:service_distance;type:float(10,2);comment:服务距离" json:"serviceDistance"`
	Status               enum.BeeShopStatus `gorm:"column:status;type:bigint(11);comment:状态" json:"status"`
	StatusStr            string             `gorm:"-" json:"statusStr"`
	TaxGst               decimal.Decimal    `gorm:"column:tax_gst;type:bigint(11);comment:发票" json:"taxGst"`
	TaxService           decimal.Decimal    `gorm:"column:tax_service;type:bigint(11);comment:发票服务" json:"taxService"`
	Type                 string             `gorm:"column:type;type:varchar(100);comment:店铺类型" json:"type"`
	Characteristic       string             `gorm:"column:type;type:varchar(1000);comment:特色" json:"characteristic"`
	BusinessScope        string             `gorm:"column:type;type:varchar(1000);comment:经营范围" json:"businessScope"`
	Introduce            string             `gorm:"column:type;type:varchar(1000);comment:介绍" json:"introduce"`
	WorkStatus           int64              `gorm:"column:work_status;type:bigint(11);comment:营业状态，0正常" json:"workStatus"`
	DadaShopNo           string             `gorm:"column:dada_shop_no;type:varchar(100);comment:达达门店编号" json:"dadaShopNo"`

	Distance float64 `gorm:"-" json:"distance"`
}

func (m *BeeShopInfo) TableName() string {
	return "bee_shop_info"
}

func (m *BeeShopInfo) FillData(lat, lon float64) {
	m.StatusStr = enum.BeeShopStatusMap[m.Status]

	// 创建 S2 点
	p1 := s2.PointFromLatLng(s2.LatLngFromDegrees(lat, lon))
	p2 := s2.PointFromLatLng(s2.LatLngFromDegrees(m.Latitude, m.Longitude))

	// 计算球面距离
	distance := s2.LatLngFromPoint(p1).Distance(s2.LatLngFromPoint(p2)).Degrees() * 111.32 // 1度约等于111.32公里
	m.Distance = distance
}
