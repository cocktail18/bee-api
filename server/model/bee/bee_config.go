// 自动生成模板BeeConfig
package bee

import (
	"time"
)

// 公用配置表 结构体  BeeConfig
type BeeConfig struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`  //id字段 
    UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`  //商店用户ID 
    IsDeleted  *bool `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`  //已删除 
    DateAdd  *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`  //创建时间 
    DateUpdate  *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`  //更新时间 
    DateDelete  *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`  //删除时间 
    Key  string `json:"key" form:"key" gorm:"column:key;comment:;size:100;"`  //key字段 
    Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:;size:100;"`  //remark字段 
    Value  string `json:"value" form:"value" gorm:"column:value;comment:;size:1000;"`  //value字段 
}


// TableName 公用配置表 BeeConfig自定义表名 bee_config
func (BeeConfig) TableName() string {
    return "bee_config"
}

