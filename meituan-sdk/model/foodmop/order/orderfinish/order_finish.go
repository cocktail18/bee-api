package orderfinish

/**
* 品牌订单完成通知
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_finish_url = "/foodmop/order/diancannew/order/finish"

type OrderFinishRequest struct {
    /**
    *  美团平台订单id 
    */
    OrderId int64 `json:"orderId"`
    /**
    *  商户订单号 
    */
    VendorOrderId string `json:"vendorOrderId"`
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
    *  请求时间戳 
    */
    Timestamp int64 `json:"timestamp"`
    /**
    *  业务类型 10-MOP 
    */
    BizType int32 `json:"bizType"`
}
type OrderFinishResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}



func (req *OrderFinishRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderFinishResponse, error) {
    resp, err := client.InvokeApi(order_finish_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderFinishResponse
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

func (response *OrderFinishResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

