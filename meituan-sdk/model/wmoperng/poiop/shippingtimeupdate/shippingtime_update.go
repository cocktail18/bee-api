package shippingtimeupdate

/**
* 更新门店营业时间
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const shippingtime_update_url = "/wmoper/ng/poiop/shippingtime/update"

type ShippingtimeUpdateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type ShippingtimeUpdateRequest struct {
    /**
    *  门店营业时间 （注意格式，且保证不同时间段之间不存在交集） 
    */
    ShippingTime string `json:"shipping_time"`
}



func (req *ShippingtimeUpdateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ShippingtimeUpdateResponse, error) {
    resp, err := client.InvokeApi(shippingtime_update_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ShippingtimeUpdateResponse
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

func (response *ShippingtimeUpdateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

