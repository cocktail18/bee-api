package groupbuyorderdetail

/**
* 套餐配送-订单详情
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const groupbuy_order_detail_url = "/dcps/fulfill/order/detail"

type OrderDetail struct {
    Status int32 `json:"status"`
    Quantity int32 `json:"quantity"`
}
type GroupbuyOrderDetailResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data OrderDetail `json:"data"`
    TraceId string `json:"traceId"`
}
type GroupbuyOrderDetailRequest struct {
    BookingOrderId string `json:"bookingOrderId"`
    OpIp string `json:"opIp"`
    OpPlatform string `json:"opPlatform"`
    OpUuid string `json:"opUuid"`
}



func (req *GroupbuyOrderDetailRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupbuyOrderDetailResponse, error) {
    resp, err := client.InvokeApi(groupbuy_order_detail_url, 46, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupbuyOrderDetailResponse
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

func (response *GroupbuyOrderDetailResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

