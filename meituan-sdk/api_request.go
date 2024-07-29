package gosdk

import (
	"encoding/json"
	"fmt"
	"meituan/gosdk/api_error"
	"meituan/gosdk/constants"
	"meituan/gosdk/log"
	"meituan/gosdk/mtclient"
)

type ApiRequest struct {
	ApiPath    string            //api路径，不包含域名
	BusinessId int               //api对应的业务id，请参考API文档的说明赋值
	BizParams  map[string]string //业务参数，key放字段名，value放字段值
}

// 创建一个新的ApiRequest
func NewApiRequest(apiPath string, businessId int, bizParams map[string]string) ApiRequest {
	return ApiRequest{ApiPath: apiPath, BusinessId: businessId, BizParams: bizParams}
}

// DoInvoke 方法用于向美团服务端发起请求。对于无需业务授权的接口，appAuthToken可以传空
func (req ApiRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ApiResponse, error) {
	resp, err := client.InvokeApi(req.ApiPath, req.BusinessId, appAuthToken, req.BizParams)
	log.Debug(fmt.Sprintf("InvokeApi response=%s, err=%s", resp, err))
	if err != nil {
		return nil, err
	}
	var response ApiResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}
	if response.Code == constants.SuccessCode {
		return &response, nil
	} else {
		return &response, &api_error.ApiError{Code: response.Code, Msg: response.Msg, TraceId: response.TraceId}
	}
}
