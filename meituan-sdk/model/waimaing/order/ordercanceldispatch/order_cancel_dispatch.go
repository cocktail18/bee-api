package ordercanceldispatch

/**
* 取消美团配送（除自配送场景）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_cancel_dispatch_url = "/waimai/order/cancelDispatch"

type OrderCancelDispatchResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 操作结果
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type OrderCancelDispatchRequest struct {
    /**
    *  订单号 
    */
    OrderId int64 `json:"orderId"`
}



func (req *OrderCancelDispatchRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderCancelDispatchResponse, error) {
    resp, err := client.InvokeApi(order_cancel_dispatch_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderCancelDispatchResponse
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

func (response *OrderCancelDispatchResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

