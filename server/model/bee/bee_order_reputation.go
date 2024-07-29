// 自动生成模板BeeOrderReputation
package bee

import (
	"time"
)

// 商品评价 结构体  BeeOrderReputation
type BeeOrderReputation struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`  //id字段 
    UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`  //商店用户ID 
    IsDeleted  *bool `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`  //已删除 
    DateAdd  *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`  //创建时间 
    DateUpdate  *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`  //更新时间 
    DateDelete  *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`  //删除时间 
    Uid  *int `json:"uid" form:"uid" gorm:"column:uid;comment:用户id;size:19;"`  //用户id 
    OrderId  string `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单id;size:100;"`  //订单id 
    Reputation  *int `json:"reputation" form:"reputation" gorm:"column:reputation;comment:评价;size:19;"`  //评价 
    Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:100;"`  //备注 
    PicsStr  string `json:"picsStr" form:"picsStr" gorm:"column:pics_str;comment:图片列表;size:100;"`  //图片列表 
}


// TableName 商品评价 BeeOrderReputation自定义表名 bee_order_reputation
func (BeeOrderReputation) TableName() string {
    return "bee_order_reputation"
}

