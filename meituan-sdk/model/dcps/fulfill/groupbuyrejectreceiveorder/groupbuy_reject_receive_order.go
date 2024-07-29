package groupbuyrejectreceiveorder

/**
* 套餐配送-拒绝接单
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const groupbuy_reject_receive_order_url = "/dcps/fulfill/reject/receive/order"

type GroupbuyRejectReceiveOrderRequest struct {
    BookingOrderId string `json:"bookingOrderId"`
    OpIp string `json:"opIp"`
    OpPlatform string `json:"opPlatform"`
    OpUuid string `json:"opUuid"`
}
type Data struct {
    BookingOrderId string `json:"bookingOrderId"`
    OpKey string `json:"opKey"`
}
type GroupbuyRejectReceiveOrderResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *GroupbuyRejectReceiveOrderRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupbuyRejectReceiveOrderResponse, error) {
    resp, err := client.InvokeApi(groupbuy_reject_receive_order_url, 46, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupbuyRejectReceiveOrderResponse
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

func (response *GroupbuyRejectReceiveOrderResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

