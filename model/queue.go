package model

import (
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/enum"
)

type BeeQueue struct {
	common.BaseModel
	Name      string              `gorm:"column:name;type:varchar(100);comment:排队名" json:"name"`
	CurNumber int64               `gorm:"column:cur_number;type:bigint(20);comment:当前叫号" json:"curNumber"`
	NumberGet int64               `gorm:"column:number_get;type:bigint(20);comment:当前取号" json:"numberGet"`
	MaxNumber int64               `gorm:"column:max_number;type:bigint(20);comment:最大叫号" json:"maxNumber"`
	Minutes   int64               `gorm:"column:minutes;type:bigint(20);comment:每个人等待的时间，分钟" json:"minutes"`
	Status    enum.BeeQueueStatus `gorm:"column:status;type:int(10);comment:状态" json:"status"`
}

type BeeUserQueue struct {
	common.BaseModel
	Uid     int64                   `gorm:"column:uid;type:bigint(20)" json:"uid"`
	QueueId int64                   `gorm:"column:queue_id;type:bigint(20)" json:"queueId"`
	Number  int64                   `gorm:"column:number;type:bigint(20)" json:"number"`
	Status  enum.BeeUserQueueStatus `gorm:"column:status;type:int(10);comment:状态" json:"status"`
}
