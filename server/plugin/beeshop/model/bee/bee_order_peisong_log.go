// 自动生成模板BeeOrderPeisongLog
package bee

import (
	"time"
)

// beeOrderPeisongLog表 结构体  BeeOrderPeisongLog
type BeeOrderPeisongLog struct {
	Id             *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`                                  //id字段
	UserId         *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`                          //商店用户ID
	IsDeleted      *bool      `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`                            //已删除
	DateAdd        *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`                                 //创建时间
	DateUpdate     *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`                        //更新时间
	DateDelete     *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`                        //删除时间
	OrderId        *int       `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单id;size:19;"`                         //订单id
	PeisongOrderNo string     `json:"peisongOrderNo" form:"peisongOrderNo" gorm:"column:peisong_order_no;comment:配送订单号;size:100;"` //配送订单号
	Remark         string     `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:500;"`                              //备注
	Log            string     `json:"log" form:"log" gorm:"column:log;comment:日志;"`                                                //日志
}

// TableName beeOrderPeisongLog表 BeeOrderPeisongLog自定义表名 bee_order_peisong_log
func (BeeOrderPeisongLog) TableName() string {
	return "bee_order_peisong_log"
}
