package cancelorder

/**
* 取消订单
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const cancel_order_url = "/peisong/order/cancel"

type CancelOrderData struct {
    /**
    * 美团配送内部订单id
    */
    MtPeisongId string `json:"mt_peisong_id"`
    /**
    * 配送活动标识
    */
    DeliveryId int64 `json:"delivery_id"`
    /**
    * 外部订单id
    */
    OrderId string `json:"order_id"`
}
type CancelOrderRequest struct {
    /**
    * 配送活动标识
    */
    DeliveryId int64 `json:"delivery_id"`
    /**
    * 美团配送内部订单id，最长不超过32个字符
    */
    MtPeisongId string `json:"mt_peisong_id"`
    /**
    * 取消原因类别，默认为接入方原因 详情请参考 美团配送开放平台接口文档--门户页面-4.3.1，客户取消订单原因
    */
    CancelReasonId int32 `json:"cancel_reason_id"`
    /**
    * 详细取消原因，最长不超过256个字符
    */
    CancelReason string `json:"cancel_reason"`
}
type CancelOrderResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data CancelOrderData `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *CancelOrderRequest) DoInvoke(client mtclient.MeituanClient) (*CancelOrderResponse, error) {
    resp, err := client.InvokeApi(cancel_order_url, 19, "", req)

    if err != nil {
        return nil, err
    }
    var response CancelOrderResponse
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

func (response *CancelOrderResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

