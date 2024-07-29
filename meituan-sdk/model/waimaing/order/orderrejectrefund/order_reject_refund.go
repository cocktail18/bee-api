package orderrejectrefund

/**
* 订单拒绝退款
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_reject_refund_url = "/waimai/order/rejectRefund"

type OrderRejectRefundRequest struct {
    /**
    *  原因 
    */
    Reason string `json:"reason"`
    /**
    *  订单号 
    */
    OrderId int64 `json:"orderId"`
    /**
    *  业务类型 
    */
    Source string `json:"source"`
    /**
    *  mop退款唯一标识 
    */
    RefundNum string `json:"refundNum"`
}
type OrderRejectRefundResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 操作结果
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *OrderRejectRefundRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderRejectRefundResponse, error) {
    resp, err := client.InvokeApi(order_reject_refund_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderRejectRefundResponse
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

func (response *OrderRejectRefundResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

