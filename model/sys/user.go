package sys

import "gitee.com/stuinfer/bee-api/common"

type SysUserModel struct {
	common.BaseModel
	Domain   string `gorm:"column:domain;type:varchar(100);unique;not null" json:"domain"`
	Username string `gorm:"column:username;type:varchar(100);unique;not null" json:"username"`
	Password string `gorm:"column:password;type:varchar(100);not null" json:"password"`
	Salt     string `gorm:"column:salt;type:varchar(100);not null" json:"salt"`
	Phone    string `gorm:"column:phone;type:varchar(100);not null" json:"phone"`
}
