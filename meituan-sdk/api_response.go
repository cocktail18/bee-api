package gosdk

import "meituan/gosdk/constants"

type ApiResponse struct {
	Code    string
	Msg     string
	Data    interface{}
	TraceId string
}

func (response *ApiResponse) IsSuccess() bool {
	return response.Code == constants.SuccessCode
}
