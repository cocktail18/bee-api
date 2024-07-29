// 自动生成模板BeeNotice
package bee

import (
	"time"
)

// 系统公告 结构体  BeeNotice
type BeeNotice struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`  //id字段 
    UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`  //商店用户ID 
    Title  string `json:"title" form:"title" gorm:"column:title;comment:标题;size:100;"`  //标题 
    Content  string `json:"content" form:"content" gorm:"column:content;comment:内容;type:text;"`  //内容 
    IsRemind  *bool `json:"isRemind" form:"isRemind" gorm:"column:is_remind;comment:是否提醒;"`  //是否提醒 
    IsShow  *bool `json:"isShow" form:"isShow" gorm:"column:is_show;comment:是否展示;"`  //是否展示 
    RemindUid  *int `json:"remindUid" form:"remindUid" gorm:"column:remind_uid;comment:提醒用户id;size:19;"`  //提醒用户id 
    DateAdd  *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`  //创建时间 
    DateUpdate  *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`  //更新时间 
    DateDelete  *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`  //删除时间 
    IsDeleted  *bool `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`  //已删除 
}


// TableName 系统公告 BeeNotice自定义表名 bee_notice
func (BeeNotice) TableName() string {
    return "bee_notice"
}

