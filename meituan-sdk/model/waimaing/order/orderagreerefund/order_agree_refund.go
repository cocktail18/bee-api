package orderagreerefund

/**
* 订单同意退款
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_agree_refund_url = "/waimai/order/agreeRefund"

type OrderAgreeRefundResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 操作结果
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type OrderAgreeRefundRequest struct {
    /**
    *  原因 
    */
    Reason string `json:"reason"`
    /**
    *  订单号 
    */
    OrderId string `json:"orderId"`
    /**
    *  业务类型 
    */
    Source string `json:"source"`
    /**
    *  mop退款唯一标识 
    */
    RefundNum string `json:"refundNum"`
}



func (req *OrderAgreeRefundRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderAgreeRefundResponse, error) {
    resp, err := client.InvokeApi(order_agree_refund_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderAgreeRefundResponse
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

func (response *OrderAgreeRefundResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

