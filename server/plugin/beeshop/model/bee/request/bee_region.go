package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BeeRegionSearch struct {
	IsDeleted       *bool      `json:"isDeleted" form:"isDeleted" `
	StartDateAdd    *time.Time `json:"startDateAdd" form:"startDateAdd"`
	EndDateAdd      *time.Time `json:"endDateAdd" form:"endDateAdd"`
	StartDateUpdate *time.Time `json:"startDateUpdate" form:"startDateUpdate"`
	EndDateUpdate   *time.Time `json:"endDateUpdate" form:"endDateUpdate"`
	FirstLetter     string     `json:"firstLetter" form:"firstLetter" `
	Jianpin         string     `json:"jianpin" form:"jianpin" `
	Level           *int       `json:"level" form:"level" `
	LevelList       []int      `json:"levelList" form:"levelList" `
	Name            string     `json:"name" form:"name" `
	NameEn          string     `json:"nameEn" form:"nameEn" `
	Pinyin          string     `json:"pinyin" form:"pinyin" `
	Pid             string     `json:"pid" form:"pid" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
