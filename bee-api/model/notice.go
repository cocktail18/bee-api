package model

import "gitee.com/stuinfer/bee-api/common"

// 店铺公告
type BeeNotice struct {
	common.BaseModel
	Content   string `gorm:"column:content;type:text;comment:内容" json:"content"`
	IsRemind  bool   `gorm:"column:is_remind;type:tinyint(1);comment:是否提醒" json:"isRemind"`
	IsShow    bool   `gorm:"column:is_show;type:tinyint(1);comment:是否展示" json:"isShow"`
	RemindUid int64  `gorm:"column:remind_uid;type:bigint(11);comment:提醒用户id" json:"remindUid"`
	Title     string `gorm:"column:title;type:varchar(100);comment:标题" json:"title"`
}

func (m *BeeNotice) TableName() string {
	return "bee_notice"
}
