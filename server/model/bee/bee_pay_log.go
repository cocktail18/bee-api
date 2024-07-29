// 自动生成模板BeePayLog
package bee

import (
	"time"
)

// 支付流水 结构体  BeePayLog
type BeePayLog struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`  //id字段 
    UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`  //商店用户ID 
    IsDeleted  *bool `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`  //已删除 
    DateAdd  *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`  //创建时间 
    DateUpdate  *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`  //更新时间 
    DateDelete  *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`  //删除时间 
    Money  *float64 `json:"money" form:"money" gorm:"column:money;comment:总金额;size:10;"`  //总金额 
    NextAction  string `json:"nextAction" form:"nextAction" gorm:"column:next_action;comment:next_action;size:100;"`  //next_action 
    OrderNo  string `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:order_no;size:100;"`  //order_no 
    PayGate  string `json:"payGate" form:"payGate" gorm:"column:pay_gate;comment:pay_gate;size:100;"`  //pay_gate 
    Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:remark;size:100;"`  //remark 
    Status  string `json:"status" form:"status" gorm:"column:status;comment:status;size:10;"`  //支付状态 
    Uid  *int `json:"uid" form:"uid" gorm:"column:uid;comment:用户id;size:19;"`  //用户id 
}


// TableName 支付流水 BeePayLog自定义表名 bee_pay_log
func (BeePayLog) TableName() string {
    return "bee_pay_log"
}

