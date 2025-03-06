package dto

import "gitee.com/stuinfer/bee-api/model"

type BeeOrderDto struct {
	model.BeeOrder
	PageNum  int `json:"pageNum" form:"pageNum"`
	PageSize int `json:"pageSize" form:"pageSize"`
}
