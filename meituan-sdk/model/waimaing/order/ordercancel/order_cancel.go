package ordercancel

/**
* 商家取消订单
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_cancel_url = "/waimai/order/cancel"

type OrderCancelResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 操作结果
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type OrderCancelRequest struct {
    /**
    *  订单号 
    */
    OrderId int64 `json:"orderId"`
    /**
    *  取消原因 
    */
    Reason string `json:"reason"`
    /**
    *  取消原因码 参考备注信息 
    */
    ReasonCode string `json:"reasonCode"`
    /**
    *  到餐线上核销填：mop，到店在线核销业务必传 
    */
    Source string `json:"source"`
}



func (req *OrderCancelRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderCancelResponse, error) {
    resp, err := client.InvokeApi(order_cancel_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderCancelResponse
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

func (response *OrderCancelResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

