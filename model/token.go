package model

import "gitee.com/stuinfer/bee-api/common"

type BeeToken struct {
	common.BaseModel
	Token    string `gorm:"column:token;unique,type:varchar(100);NOT NULL" json:"token"`
	Uid      int64  `gorm:"column:uid;type:bigint(20);NOT NULL" json:"uid"`
	ExpireAt int64  `gorm:"column:expire_at;type:bigint(20);NOT NULL" json:"expire_at"`
}

func (m *BeeToken) TableName() string {
	return "bee_token"
}
