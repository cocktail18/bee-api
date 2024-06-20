package model

import "gitee.com/stuinfer/bee-api/common"

type BeeUploadFile struct {
	common.BaseModel
	Domain   string `gorm:"column:domain;type:varchar(255);comment:domain" json:"domain"`
	Filename string `gorm:"column:filename;type:varchar(255);comment:文件名" json:"filename"`
	Dst      string `gorm:"column:dst;type:varchar(255);comment:保存路径" json:"dst"`
	ExpireAt int64  `gorm:"column:expire_at;type:bigint(20);comment:过期时间" json:"expire_at"`
}

func (b *BeeUploadFile) TableName() string {
	return "bee_upload_file"
}
