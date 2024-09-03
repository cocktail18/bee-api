// 自动生成模板BeeUserQueue
package bee

import (
	"time"
)

// beeUserQueue表 结构体  BeeUserQueue
type BeeUserQueue struct {
	Id         *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`           //id字段
	UserId     *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`   //商店用户ID
	IsDeleted  *bool      `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`     //已删除
	DateAdd    *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`          //创建时间
	DateUpdate *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"` //更新时间
	DateDelete *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"` //删除时间
	Uid        *int       `json:"uid" form:"uid" gorm:"column:uid;comment:;size:19;"`                   //uid字段
	QueueId    *int       `json:"queueId" form:"queueId" gorm:"column:queue_id;comment:;size:19;"`      //queueId字段
	Number     *int       `json:"number" form:"number" gorm:"column:number;comment:;size:19;"`          //number字段
	Status     *int       `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`        //状态
}

// TableName beeUserQueue表 BeeUserQueue自定义表名 bee_user_queue
func (BeeUserQueue) TableName() string {
	return "bee_user_queue"
}
