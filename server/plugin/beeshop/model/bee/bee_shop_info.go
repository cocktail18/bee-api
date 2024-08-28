// 自动生成模板BeeShopInfo
package bee

import (
	"time"
)

// 商店信息 结构体  BeeShopInfo
type BeeShopInfo struct {
	Id                   *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`                                                  //id字段
	UserId               *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`                                          //商店用户ID
	IsDeleted            *bool      `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`                                            //已删除
	DateAdd              *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`                                                 //创建时间
	DateUpdate           *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`                                        //更新时间
	DateDelete           *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`                                        //删除时间
	Address              string     `json:"address" form:"address" gorm:"column:address;comment:地址;size:100;"`                                           //地址
	BindUid              *int       `json:"bindUid" form:"bindUid" gorm:"column:bind_uid;comment:绑定用户id;size:19;"`                                       //绑定用户id
	CityId               string     `json:"cityId" form:"cityId" gorm:"column:city_id;comment:城市id;size:100;"`                                           //城市id
	CyTablePayMod        *int       `json:"cyTablePayMod" form:"cyTablePayMod" gorm:"column:cy_table_pay_mod;comment:支付方式;size:19;"`                     //支付方式
	DistrictId           string     `json:"districtId" form:"districtId" gorm:"column:district_id;comment:唯一id;size:100;"`                               //唯一id
	ExpressType          string     `json:"expressType" form:"expressType" gorm:"column:express_type;comment:快递类型;size:100;"`                            //快递类型
	GoodsNeedCheck       *bool      `json:"goodsNeedCheck" form:"goodsNeedCheck" gorm:"column:goods_need_check;comment:商品需要检查;"`                         //商品需要检查
	Latitude             *float64   `json:"latitude" form:"latitude" gorm:"column:latitude;comment:纬度;size:20;"`                                         //纬度
	LinkPhone            string     `json:"linkPhone" form:"linkPhone" gorm:"column:link_phone;comment:绑定手机;size:100;"`                                  //绑定手机
	Longitude            *float64   `json:"longitude" form:"longitude" gorm:"column:longitude;comment:经度;size:20;"`                                      //经度
	Name                 string     `json:"name" form:"name" gorm:"column:name;comment:店名;size:100;"`                                                    //店名
	Number               string     `json:"number" form:"number" gorm:"column:number;comment:座机号码;size:100;"`                                            //座机号码
	NumberFav            *int       `json:"numberFav" form:"numberFav" gorm:"column:number_fav;comment:喜欢人数;size:19;"`                                   //喜欢人数
	NumberGoodReputation *int       `json:"numberGoodReputation" form:"numberGoodReputation" gorm:"column:number_good_reputation;comment:商品评分;size:19;"` //商品评分
	NumberOrder          *int       `json:"numberOrder" form:"numberOrder" gorm:"column:number_order;comment:下单数;size:19;"`                              //下单数
	NumberReputation     *int       `json:"numberReputation" form:"numberReputation" gorm:"column:number_reputation;comment:店铺评分;size:19;"`              //店铺评分
	OpenScan             *bool      `json:"openScan" form:"openScan" gorm:"column:open_scan;comment:开启扫描;"`                                              //开启扫描
	OpenWaimai           *bool      `json:"openWaimai" form:"openWaimai" gorm:"column:open_waimai;comment:开启外卖;"`                                        //开启外卖
	OpenZiqu             *bool      `json:"openZiqu" form:"openZiqu" gorm:"column:open_ziqu;comment:开启自取;"`                                              //开启自取
	OpeningHours         string     `json:"openingHours" form:"openingHours" gorm:"column:opening_hours;comment:开店时间;size:100;"`                         //开店时间
	Paixu                *int       `json:"paixu" form:"paixu" gorm:"column:paixu;comment:排序;size:19;"`                                                  //排序
	ProvinceId           string     `json:"provinceId" form:"provinceId" gorm:"column:province_id;comment:省id;size:100;"`                                //省id
	RecommendStatus      *int       `json:"recommendStatus" form:"recommendStatus" gorm:"column:recommend_status;comment:推荐状态;size:19;"`                 //推荐状态
	ServiceDistance      *float64   `json:"serviceDistance" form:"serviceDistance" gorm:"column:service_distance;comment:服务距离;"`                         //服务距离
	Status               *int       `json:"status" form:"status" gorm:"column:status;comment:状态;size:19;"`                                               //状态
	TaxGst               *int       `json:"taxGst" form:"taxGst" gorm:"column:tax_gst;comment:发票;size:19;"`                                              //发票
	TaxService           *int       `json:"taxService" form:"taxService" gorm:"column:tax_service;comment:发票服务;size:19;"`                                //发票服务
	Type                 string     `json:"type" form:"type" gorm:"column:type;comment:店铺类型;size:100;"`                                                  //店铺类型
	WorkStatus           *bool      `json:"workStatus" form:"workStatus" gorm:"column:work_status;comment:营业状态;"`                                        //营业状态
	DadaShopNo           string     `json:"dadaShopNo" form:"dadaShopNo" gorm:"column:dada_shop_no;comment:达达门店编号;size:100;"`
}

// TableName 商店信息 BeeShopInfo自定义表名 bee_shop_info
func (BeeShopInfo) TableName() string {
	return "bee_shop_info"
}
