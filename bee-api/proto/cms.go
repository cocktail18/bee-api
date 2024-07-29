package proto

import "gitee.com/stuinfer/bee-api/model"

type CmsResp struct {
	ExtJson interface{}       `json:"extJson"`
	Info    *model.BeeCmsInfo `json:"info"`
}
