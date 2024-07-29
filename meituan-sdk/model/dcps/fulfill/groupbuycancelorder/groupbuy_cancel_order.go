package groupbuycancelorder

/**
* 套餐配送-取消订单
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const groupbuy_cancel_order_url = "/dcps/fulfill/cancel/order"

type GroupbuyCancelOrderRequest struct {
    BookingOrderId string `json:"bookingOrderId"`
    OpIp string `json:"opIp"`
    OpPlatform string `json:"opPlatform"`
    OpUuid string `json:"opUuid"`
    CancelReason string `json:"cancelReason"`
}
type Data struct {
    BookingOrderId string `json:"bookingOrderId"`
    OpKey string `json:"opKey"`
}
type GroupbuyCancelOrderResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *GroupbuyCancelOrderRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupbuyCancelOrderResponse, error) {
    resp, err := client.InvokeApi(groupbuy_cancel_order_url, 46, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupbuyCancelOrderResponse
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

func (response *GroupbuyCancelOrderResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

