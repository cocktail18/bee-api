// 自动生成模板BeeRegion
package bee

import (
	"time"
)

// 地址库 结构体  BeeRegion
type BeeRegion struct {
	Id          *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`                  //id字段
	UserId      *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`          //商店用户ID
	IsDeleted   *bool      `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`            //已删除
	DateAdd     *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`                 //创建时间
	DateUpdate  *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`        //更新时间
	DateDelete  *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`        //删除时间
	FirstLetter string     `json:"firstLetter" form:"firstLetter" gorm:"column:first_letter;comment:;size:10;"` //firstLetter字段
	Jianpin     string     `json:"jianpin" form:"jianpin" gorm:"column:jianpin;comment:;size:100;"`             //jianpin字段
	Level       *int       `json:"level" form:"level" gorm:"column:level;comment:;size:10;"`                    //level字段
	Name        string     `json:"name" form:"name" gorm:"column:name;comment:;size:100;"`                      //name字段
	NameEn      string     `json:"nameEn" form:"nameEn" gorm:"column:name_en;comment:;size:200;"`               //nameEn字段
	Pinyin      string     `json:"pinyin" form:"pinyin" gorm:"column:pinyin;comment:;size:100;"`                //pinyin字段
	Pid         string     `json:"pid" form:"pid" gorm:"column:pid;comment:;size:100;"`                         //pid字段
	RegionId    *int       `json:"regionId" form:"regionId" gorm:"column:region_id;comment:;size:19;"`          //regionId字段

	Childs []*BeeRegion `json:"childs" form:"childs" gorm:"-"`
}

// TableName 地址库 BeeRegion自定义表名 bee_region
func (BeeRegion) TableName() string {
	return "bee_region"
}
