package orderfullfillnotice

/**
* 品牌订单状态变更通知
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_full_fill_notice_url = "/foodmop/order/diancannew/order/fullFillNotice"

type OrderFullFillNoticeRequest struct {
    /**
    *  美团平台订单id 
    */
    OrderId int64 `json:"orderId"`
    /**
    *  请求时间戳 
    */
    Timestamp int64 `json:"timestamp"`
    /**
    *  商户订单号 
    */
    VendorOrderId string `json:"vendorOrderId"`
    /**
    *  订单履约动作类型： START_MAKING(1, "预约单开始制作（进入到制作中）"), 实时单接单后就直接进入制作中 MAKE_COMPLETE(2, "制作完成（进入到待取餐）"); 
    */
    OrderFulFillActionType int32 `json:"orderFulFillActionType"`
    /**
    *  操作者Id 
    */
    OperatorId string `json:"operatorId"`
    /**
    *  操作者名称 
    */
    OperatorName string `json:"operatorName"`
    /**
    *  处理时间 
    */
    HandleTime int64 `json:"handleTime"`
    /**
    *  业务类型 10-MOP 
    */
    BizType int32 `json:"bizType"`
}
type OrderFullFillNoticeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}



func (req *OrderFullFillNoticeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderFullFillNoticeResponse, error) {
    resp, err := client.InvokeApi(order_full_fill_notice_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderFullFillNoticeResponse
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

func (response *OrderFullFillNoticeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

