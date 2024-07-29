// 自动生成模板BeeComment
package bee

import (
	"time"
)

// 评论表 结构体  BeeComment
type BeeComment struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`  //id字段 
    UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`  //商店用户ID 
    IsDeleted  *bool `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`  //已删除 
    DateAdd  *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`  //创建时间 
    DateUpdate  *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`  //更新时间 
    DateDelete  *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`  //删除时间 
    Uid  *int `json:"uid" form:"uid" gorm:"column:uid;comment:uid;size:19;"`  //用户id 
    RefId  *int `json:"refId" form:"refId" gorm:"column:ref_id;comment:ref_id;size:19;"`  //关联id 
    Pid  *int `json:"pid" form:"pid" gorm:"column:pid;comment:pid;size:19;"`  //父评论id 
    ShopId  *int `json:"shopId" form:"shopId" gorm:"column:shop_id;comment:shop_id;size:19;"`  //商店id 
    Type  *int `json:"type" form:"type" gorm:"column:type;comment:type;size:10;"`  //类型 
    Pics  string `json:"pics" form:"pics" gorm:"column:pics;comment:pics;size:2000;"`  //图片 
    File  string `json:"file" form:"file" gorm:"column:file;comment:file;size:255;"`  //文件 
    Mobile  string `json:"mobile" form:"mobile" gorm:"column:mobile;comment:mobile;size:100;"`  //手机号 
    Content  string `json:"content" form:"content" gorm:"column:content;comment:content;size:2000;"`  //内容 
    ExtJsonStr  string `json:"extJsonStr" form:"extJsonStr" gorm:"column:ext_json_str;comment:ext_json_str;size:2000;"`  //额外json信息 
    Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`  //状态 
}


// TableName 评论表 BeeComment自定义表名 bee_comment
func (BeeComment) TableName() string {
    return "bee_comment"
}

