package dada_sdk

import "encoding/json"

type Response struct {
	Status    string          `json:"status"`    // 响应状态，成功为"success"，失败为"fail"
	Code      int             `json:"code"`      // 响应返回吗，参考接口返回码
	Msg       string          `json:"msg"`       //	响应描述
	Result    json.RawMessage `json:"result"`    //	响应结果，JSON对象，详见具体的接口描述
	ErrorCode int             `json:"errorCode"` //	错误编码，与code一致
}
