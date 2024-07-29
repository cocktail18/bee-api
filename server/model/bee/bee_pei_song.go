// 自动生成模板BeePeiSong
package bee

import (
	"time"
)

// 配送信息 结构体  BeePeiSong
type BeePeiSong struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`  //id字段 
    UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`  //商店用户ID 
    IsDeleted  *bool `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`  //已删除 
    DateAdd  *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`  //创建时间 
    DateUpdate  *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`  //更新时间 
    DateDelete  *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`  //删除时间 
    Distance  *float64 `json:"distance" form:"distance" gorm:"column:distance;comment:距离，单位km;size:10;"`  //距离，单位km 
    Fwf1Min  *float64 `json:"fwf1Min" form:"fwf1Min" gorm:"column:fwf1_min;comment:最低收费金额;size:10;"`  //最低收费金额 
    Fwf1Name  *float64 `json:"fwf1Name" form:"fwf1Name" gorm:"column:fwf1_name;comment:收费项目名称;size:10;"`  //收费项目名称 
    Fwf1Number  *float64 `json:"fwf1Number" form:"fwf1Number" gorm:"column:fwf1_number;comment:比例或者金额;size:10;"`  //比例或者金额 
    Fwf1Type  *int `json:"fwf1Type" form:"fwf1Type" gorm:"column:fwf1_type;comment:类型;size:10;"`  //类型 
    ZtDiscounts  *float64 `json:"ztDiscounts" form:"ztDiscounts" gorm:"column:zt_discounts;comment:自提补贴;size:10;"`  //自提补贴 
}


// TableName 配送信息 BeePeiSong自定义表名 bee_pei_song
func (BeePeiSong) TableName() string {
    return "bee_pei_song"
}

