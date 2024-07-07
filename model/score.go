package model

import (
	"gitee.com/stuinfer/bee-api/common"
	"github.com/shopspring/decimal"
)

type BeeScoreLog struct {
	common.BaseModel
	Uid    int64           `gorm:"column:uid;type:bigint(11);comment:用户id" json:"uid"`
	Score  decimal.Decimal `gorm:"column:score;type:bigint(20);comment:获得积分" json:"score"`
	Remark string          `gorm:"column:remark;type:varchar(100);comment:备注" json:"remark"`
}

func (m *BeeScoreLog) TableName() string {
	return "bee_score_log"
}
