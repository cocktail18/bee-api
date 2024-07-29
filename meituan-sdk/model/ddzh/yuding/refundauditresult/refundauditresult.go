package refundauditresult

/**
* 取消预订审核结果回调
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const refundauditresult_url = "/ddzh/yuding/refundauditresult"

type RefundauditresultResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type RefundauditresultRequest struct {
    /**
    *  审核结果，2-审核通过，3-审核未通过 
    */
    AuditResult int32 `json:"auditResult"`
    /**
    *  开放平台上的订单ID 
    */
    OrderId string `json:"orderId"`
    /**
    *  生活服务行业新增字段，其他行业不传。 交易类型，枚举值：1-预订，2-团单（团购后预约） 商家根据此属性字段识别交易类型，预订业务该字段非必传，团购后预约业务必传 
    */
    Type int32 `json:"type"`
}
type Data struct {
    OrderId string `json:"orderId"`
}



func (req *RefundauditresultRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*RefundauditresultResponse, error) {
    resp, err := client.InvokeApi(refundauditresult_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response RefundauditresultResponse
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

func (response *RefundauditresultResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

