package proto

import "gitee.com/stuinfer/bee-api/model"

type GetLevelListResp struct {
	PageRespCommon
	Result []*model.BeeLevel `json:"result"`
}
