package proto

import "gitee.com/stuinfer/bee-api/enum"

type Response struct {
	Code enum.ResCode `json:"code"`
	Data interface{}  `json:"data"`
	Msg  string       `json:"msg"`
}
