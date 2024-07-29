// 自动生成模板BeeUploadFile
package bee

import (
	"time"
)

// 用户上传文件 结构体  BeeUploadFile
type BeeUploadFile struct {
	Id         *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`             //id字段
	UserId     *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`     //商店用户ID
	IsDeleted  *bool      `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`       //已删除
	DateAdd    *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`            //创建时间
	DateUpdate *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`   //更新时间
	DateDelete *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`   //删除时间
	Domain     string     `json:"domain" form:"domain" gorm:"column:domain;comment:domain;size:255;"`     //domain
	Filename   string     `json:"filename" form:"filename" gorm:"column:filename;comment:文件名;size:255;"`  //文件名
	Dst        string     `json:"dst" form:"dst" gorm:"column:dst;comment:保存路径;size:255;"`                //保存路径
	ExpireAt   *int       `json:"expireAt" form:"expireAt" gorm:"column:expire_at;comment:过期时间;size:19;"` //过期时间
}

// TableName 用户上传文件 BeeUploadFile自定义表名 bee_upload_file
func (BeeUploadFile) TableName() string {
	return "bee_upload_file"
}
