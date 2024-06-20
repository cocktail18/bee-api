package model

import "gitee.com/stuinfer/bee-api/common"

// root
type BeeSignLog struct {
	common.BaseModel
	DateAdd    common.JsonTime `gorm:"column:date_add;type:datetime(3);comment:签到时间" json:"date_add"`
	Uid        int64           `gorm:"column:uid;type:bigint(11);comment:用户id" json:"uid"`
	Continuous int             `gorm:"column:continuous;type:bigint(11);comment:连续签到次数" json:"continuous"`
}

func (m *BeeSignLog) TableName() string {
	return "bee_sign_log"
}
