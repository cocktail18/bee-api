package merchantcancelorder

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

const merchantcancelorder_url = "/ddzh/yuding/merchantcancelorder"

type MerchantcancelorderRequest struct {
    /**
    *  订单id 
    */
    OrderId string `json:"orderId"`
    /**
    *  取消原因 
    */
    CancelReason string `json:"cancelReason"`
    /**
    *  生活服务行业新增字段，其他行业不传。 交易类型，枚举值： 1-预订（旧字段，后面会废弃）， 2-团单（团购后预约）， 3-预订（新增字段，支持核销后取消订单）， 注：商家根据此属性字段识别交易类型，团购后预约业务必传，预订业务传3，不传参时默认参数为1。 
    */
    Type int32 `json:"type"`
}
type MerchantcancelorderResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Data struct {
    Result string `json:"result"`
}



func (req *MerchantcancelorderRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*MerchantcancelorderResponse, error) {
    resp, err := client.InvokeApi(merchantcancelorder_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response MerchantcancelorderResponse
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

func (response *MerchantcancelorderResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

