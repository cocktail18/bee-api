package groupbuyselfdelivery

/**
* 套餐配送-转自配
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const groupbuy_self_delivery_url = "/dcps/fulfill/self/delivery"

type GroupbuySelfDeliveryRequest struct {
    BookingOrderId string `json:"bookingOrderId"`
    OpIp string `json:"opIp"`
    OpPlatform string `json:"opPlatform"`
    OpUuid string `json:"opUuid"`
}
type GroupbuySelfDeliveryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Data struct {
    BookingOrderId string `json:"bookingOrderId"`
    OpKey string `json:"opKey"`
}



func (req *GroupbuySelfDeliveryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupbuySelfDeliveryResponse, error) {
    resp, err := client.InvokeApi(groupbuy_self_delivery_url, 46, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupbuySelfDeliveryResponse
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

func (response *GroupbuySelfDeliveryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

