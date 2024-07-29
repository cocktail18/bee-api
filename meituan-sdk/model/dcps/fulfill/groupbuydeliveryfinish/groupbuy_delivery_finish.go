package groupbuydeliveryfinish

/**
* 套餐配送-自配完成
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const groupbuy_delivery_finish_url = "/dcps/fulfill/delivery/finish"

type GroupbuyDeliveryFinishResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Data struct {
    BookingOrderId string `json:"bookingOrderId"`
    OpKey string `json:"opKey"`
}
type GroupbuyDeliveryFinishRequest struct {
    BookingOrderId string `json:"bookingOrderId"`
    OpIp string `json:"opIp"`
    OpPlatform string `json:"opPlatform"`
    OpUuid string `json:"opUuid"`
}



func (req *GroupbuyDeliveryFinishRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupbuyDeliveryFinishResponse, error) {
    resp, err := client.InvokeApi(groupbuy_delivery_finish_url, 46, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupbuyDeliveryFinishResponse
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

func (response *GroupbuyDeliveryFinishResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

