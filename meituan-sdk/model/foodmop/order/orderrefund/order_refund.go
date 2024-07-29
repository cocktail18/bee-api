package orderrefund

/**
* 商家发起退款
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_refund_url = "/foodmop/order/refund"

type OrderRefundRequest struct {
    /**
    *  美团平台订单id 
    */
    OrderId int64 `json:"orderId"`
    /**
    *  商户订单号 
    */
    VendorOrderId string `json:"vendorOrderId"`
    /**
    *  请求时间戳 
    */
    Timestamp int64 `json:"timestamp"`
    /**
    *  商家退款原因： MERCHANT_REFUND(1, "商家原因退款"), OTHER_REASON_REFUND(2, "其他原因退款"); 
    */
    VendorDirectRefundType int32 `json:"vendorDirectRefundType"`
    /**
    *  是否已取餐 
    */
    Consumed bool `json:"consumed"`
    /**
    *  操作者名称 
    */
    OperatorName string `json:"operatorName"`
    /**
    *  操作者Id 
    */
    OperatorId string `json:"operatorId"`
    /**
    *  申请退款原因 
    */
    RefundReason string `json:"refundReason"`
    /**
    *  处理时间 
    */
    HandleTime int64 `json:"handleTime"`
    /**
    *  业务类型 10-MOP 
    */
    BizType int32 `json:"bizType"`
}
type OrderRefundResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}



func (req *OrderRefundRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderRefundResponse, error) {
    resp, err := client.InvokeApi(order_refund_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderRefundResponse
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

func (response *OrderRefundResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

