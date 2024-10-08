package queryzbcanceldeliveryreason

/**
* 获取订单可以取消跑腿的原因
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_zb_cancel_delivery_reason_url = "/wmoper/ng/order/queryZbCancelDeliveryReason"

type QueryZbCancelDeliveryReasonResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Reason `json:"data"`
    TraceId string `json:"traceId"`
}
type QueryZbCancelDeliveryReasonRequest struct {
    /**
    *  外卖订单id 
    */
    OrderId int64 `json:"orderId"`
}
type Reasonlist struct {
    /**
    * 原因代码
    */
    Code int32 `json:"code"`
    /**
    * 具体原因
    */
    Content string `json:"content"`
    /**
    * 预取消code
    */
    PreCancelCode string `json:"preCancelCode"`
    /**
    * 预取消描述
    */
    PreCancelMsg string `json:"preCancelMsg"`
}
type Reason struct {
    /**
    * 状态码
    */
    Code string `json:"code"`
    /**
    * 消息
    */
    Msg string `json:"msg"`
    /**
    * 返回原因标题
    */
    Title string `json:"title"`
    /**
    * 配送状态
    */
    DeliveryStatus int32 `json:"deliveryStatus"`
    ReasonList []Reasonlist `json:"reasonList"`
}



func (req *QueryZbCancelDeliveryReasonRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryZbCancelDeliveryReasonResponse, error) {
    resp, err := client.InvokeApi(query_zb_cancel_delivery_reason_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryZbCancelDeliveryReasonResponse
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

func (response *QueryZbCancelDeliveryReasonResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

