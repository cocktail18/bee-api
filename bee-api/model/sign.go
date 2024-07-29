package model

import "gitee.com/stuinfer/bee-api/common"

// root
type BeeSignLog struct {
	common.BaseModel
	Uid       int64 `gorm:"column:uid;type:bigint(11);comment:用户id" json:"uid"`
	Continues int   `gorm:"column:continuous;type:bigint(11);comment:连续签到次数" json:"continuous"`
	Score     int64 `gorm:"column:score;type:bigint(20);comment:获得积分" json:"score"`
}

func (m *BeeSignLog) TableName() string {
	return "bee_sign_log"
}
