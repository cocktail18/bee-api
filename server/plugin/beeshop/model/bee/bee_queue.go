// 自动生成模板BeeQueue
package bee

import (
	"time"
)

// beeQueue表 结构体  BeeQueue
type BeeQueue struct {
	Id         *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`                //id字段
	UserId     *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`        //商店用户ID
	IsDeleted  *bool      `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`          //已删除
	DateAdd    *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`               //创建时间
	DateUpdate *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`      //更新时间
	DateDelete *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`      //删除时间
	Name       string     `json:"name" form:"name" gorm:"column:name;comment:排队名;size:100;"`                 //排队名
	CurNumber  *int       `json:"curNumber" form:"curNumber" gorm:"column:cur_number;comment:当前叫号;size:19;"` //当前叫号
	NumberGet  *int       `json:"numberGet" form:"numberGet" gorm:"column:number_get;comment:当前取号;size:19;"` //当前取号
	MaxNumber  *int       `json:"maxNumber" form:"maxNumber" gorm:"column:max_number;comment:最大叫号;size:19;"` //最大叫号
	Minutes    *int       `json:"minutes" form:"minutes" gorm:"column:minutes;comment:每个人等待的时间，分钟;size:19;"` //每个号预计等待的时间，分钟
	Status     *int       `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`             //状态
}

// TableName beeQueue表 BeeQueue自定义表名 bee_queue
func (BeeQueue) TableName() string {
	return "bee_queue"
}
