package shippingdelete

/**
* 删除门店配送范围（自配）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const shipping_delete_url = "/wmoper/ng/shippingop/delete"

type ShippingDeleteResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type ShippingDeleteRequest struct {
    /**
    *  服务商方提供的配送范围id 
    */
    AppShippingCode string `json:"app_shipping_code"`
}



func (req *ShippingDeleteRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ShippingDeleteResponse, error) {
    resp, err := client.InvokeApi(shipping_delete_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ShippingDeleteResponse
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

func (response *ShippingDeleteResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

