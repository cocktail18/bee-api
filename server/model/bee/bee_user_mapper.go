// 自动生成模板BeeUserMapper
package bee

import (
	"time"
)

// beeUserMapper表 结构体  BeeUserMapper
type BeeUserMapper struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`  //id字段 
    UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`  //商店用户ID 
    IsDeleted  *bool `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`  //已删除 
    DateAdd  *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`  //创建时间 
    DateUpdate  *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`  //更新时间 
    DateDelete  *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`  //删除时间 
    Uid  *int `json:"uid" form:"uid" gorm:"column:uid;comment:;size:19;" binding:"required"`  //uid字段 
    Source  *int `json:"source" form:"source" gorm:"column:source;comment:;size:19;" binding:"required"`  //source字段 
    OpenId  string `json:"openId" form:"openId" gorm:"column:open_id;comment:;size:200;" binding:"required"`  //openId字段 
    UnionId  string `json:"unionId" form:"unionId" gorm:"column:union_id;comment:;size:200;"`  //unionId字段 
}


// TableName beeUserMapper表 BeeUserMapper自定义表名 bee_user_mapper
func (BeeUserMapper) TableName() string {
    return "bee_user_mapper"
}

