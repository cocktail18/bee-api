package model

import "gitee.com/stuinfer/bee-api/common"

type BeeToken struct {
	common.BaseModel
	SessionKey string `gorm:"column:session_key;type:varchar(100);NOT NULL;default:''" json:"session_key"`
	Token      string `gorm:"column:token;unique,type:varchar(100);NOT NULL" json:"token"`
	Uid        int64  `gorm:"column:uid;type:bigint(20);NOT NULL" json:"uid"`
	ExpireAt   int64  `gorm:"column:expire_at;type:bigint(20);NOT NULL" json:"expire_at"`
}

func (m *BeeToken) TableName() string {
	return "bee_token"
}
