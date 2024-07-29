package shopareacallback

/**
* 模拟门店配送范围变更
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const shop_area_callback_url = "/peisong/test/shop/area/callback"

type ShopAreaCallbackRequest struct {
    /**
    * 取货门店id，即合作方向美团提供的门店id 注：测试门店的shop_id固定为 test_0001 ，仅用于对接时联调测试。
    */
    ShopId string `json:"shop_id"`
    /**
    * 配送服务代码，详情见合同 飞速达:4002 快速达:4011 及时达:4012 集中送:4013
    */
    DeliveryServiceCode int32 `json:"delivery_service_code"`
}
type ShopAreaCallbackResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}



func (req *ShopAreaCallbackRequest) DoInvoke(client mtclient.MeituanClient) (*ShopAreaCallbackResponse, error) {
    resp, err := client.InvokeApi(shop_area_callback_url, 19, "", req)

    if err != nil {
        return nil, err
    }
    var response ShopAreaCallbackResponse
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

func (response *ShopAreaCallbackResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

