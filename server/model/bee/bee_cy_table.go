// 自动生成模板BeeCyTable
package bee

import (
	"time"
)

// 桌号信息 结构体  BeeCyTable
type BeeCyTable struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`  //id字段 
    UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`  //商店用户ID 
    IsDeleted  *bool `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`  //已删除 
    DateAdd  *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`  //创建时间 
    DateUpdate  *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`  //更新时间 
    DateDelete  *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`  //删除时间 
    ShopId  *int `json:"shopId" form:"shopId" gorm:"column:shop_id;comment:;size:19;"`  //shopId字段 
    TableNum  string `json:"tableNum" form:"tableNum" gorm:"column:table_num;comment:;size:255;"`  //tableNum字段 
    Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:;size:255;"`  //remark字段 
    Uid  *int `json:"uid" form:"uid" gorm:"column:uid;comment:;size:19;"`  //uid字段 
}


// TableName 桌号信息 BeeCyTable自定义表名 bee_cy_table
func (BeeCyTable) TableName() string {
    return "bee_cy_table"
}

