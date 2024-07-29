package servestatusupdate

/**
* “线上点”门店营业状态批量变更（必接）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const serve_status_update_url = "/foodmop/shop/serve/status/update"

type ServeStatusUpdateRequest struct {
    /**
    *  厂商配置列表，单次请求批量限制20 
    */
    VendorServiceConfigDTOS []VendorServiceConfigRequest `json:"vendorServiceConfigDTOS"`
}
type VendorServiceConfigRequest struct {
    /**
    *  厂商门店ID 
    */
    VendorShopId string `json:"vendorShopId"`
    /**
    *  线上点营业时间，需要提前约定好同步模式，支持天模式/周模式： 天模式（key值为具体时间，仅支持更新未来7天内的营业时间） 未来7天的营业时间 { "2021-12-01":// [{ "startTime":"08:00", //Time "endTime":"12:00" }], "2021-12-02":[{ "startTime":"08:00", "endTime":"18:00" }] } 周模式（key值为星期几） { "1":// [{ "startTime":"08:00", //Time "endTime":"12:00" }], "2":[{ "startTime":"08:00", "endTime":"18:00" }] ... } 
    */
    MopBusinessHour MopBusinessHour `json:"mopBusinessHour"`
    /**
    *  线上点服务开关，控制线上点菜单是否能用 10: 营业 20：歇业 
    */
    VendorServiceOpen int32 `json:"vendorServiceOpen"`
}
type ResultData struct {
}
type ServeStatusUpdateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * Key：vendorShopId（品牌门店id) Value：OperateResulst（操作结果）
    */
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}
type MopBusinessHour struct {
}



func (req *ServeStatusUpdateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ServeStatusUpdateResponse, error) {
    resp, err := client.InvokeApi(serve_status_update_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ServeStatusUpdateResponse
    err = json.Unmarshal(resp, &response)
    if err != nil {
        return nil, err
    }

    if response.IsSuccess() {
        return &response, nil
    } else {
        return &response, &api_error.ApiError{Code: response.Code, Msg: response.Msg, TraceId: response.TraceId}
    }
}

func (response *ServeStatusUpdateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

