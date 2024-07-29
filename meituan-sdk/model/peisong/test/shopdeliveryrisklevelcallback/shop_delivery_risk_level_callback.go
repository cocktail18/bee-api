package shopdeliveryrisklevelcallback

/**
* 模拟门店配送风险等级变更
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const shop_delivery_risk_level_callback_url = "/peisong/test/shop/deliveryRiskLevel/callback"

type ShopDeliveryRiskLevelCallbackRequest struct {
    /**
    * 取货门店id，即合作方向美团提供的门店id 注：测试门店的shop_id固定为 test_0001 ，仅用于对接时联调测试。
    */
    ShopId string `json:"shop_id"`
}
type ShopDeliveryRiskLevelCallbackResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}



func (req *ShopDeliveryRiskLevelCallbackRequest) DoInvoke(client mtclient.MeituanClient) (*ShopDeliveryRiskLevelCallbackResponse, error) {
    resp, err := client.InvokeApi(shop_delivery_risk_level_callback_url, 19, "", req)

    if err != nil {
        return nil, err
    }
    var response ShopDeliveryRiskLevelCallbackResponse
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

func (response *ShopDeliveryRiskLevelCallbackResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

