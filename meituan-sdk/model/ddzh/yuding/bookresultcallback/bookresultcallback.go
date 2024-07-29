package bookresultcallback

/**
* 预订结果回调
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const bookresultcallback_url = "/ddzh/yuding/bookresultcallback"

type BookresultcallbackResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data RbookResultCallBackRes `json:"data"`
    TraceId string `json:"traceId"`
}
type RbookResultCallBackRes struct {
    /**
    * 平台订单id
    */
    OrderId string `json:"orderId"`
}
type BookresultcallbackRequest struct {
    /**
    *  美团订单id 
    */
    OrderId string `json:"orderId"`
    /**
    *  预订结果，2-预订成功，3-预订失败 
    */
    BookStatus int64 `json:"bookStatus"`
    /**
    *  业务code，应遵从预订接口的错误码，成功返回200； 预订失败时用户侧会根据code展示对应文案，详细可参考参数说明 
    */
    Code int64 `json:"code"`
    /**
    *  三方的订单号 
    */
    AppOrderId string `json:"appOrderId"`
    /**
    *  生活服务行业新增字段，其他行业不传。交易类型，枚举值：1-预订，2-团单（团购后预约）。商家根据此属性字段识别交易类型，预订业务该字段非必传，团购后预约业务必传 
    */
    Type int64 `json:"type"`
    /**
    *  扩展信息，预订失败时无需传该参数 
    */
    ExtMap string `json:"extMap"`
    /**
    *  服务人员id，在美团侧通过接口生成的服务人员id 
    */
    ServiceTechId string `json:"serviceTechId"`
}



func (req *BookresultcallbackRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*BookresultcallbackResponse, error) {
    resp, err := client.InvokeApi(bookresultcallback_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response BookresultcallbackResponse
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

func (response *BookresultcallbackResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

