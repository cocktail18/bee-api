package changeresultcallback

/**
* 改约结果回调
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const changeresultcallback_url = "/ddzh/yuding/changeresultcallback"

type ChangeresultcallbackRequest struct {
    /**
    *  订单id 
    */
    ReserveOrderId int64 `json:"reserveOrderId"`
    /**
    *  改约结果，2-成功，3-失败 
    */
    ChangeOrderResult int32 `json:"changeOrderResult"`
    /**
    *  业务code，应遵从接口的错误码，成功返回200 
    */
    Code int32 `json:"code"`
}
type ChangeresultcallbackResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Data struct {
    /**
    * true-成功，false-失败
    */
    Result bool `json:"result"`
}



func (req *ChangeresultcallbackRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ChangeresultcallbackResponse, error) {
    resp, err := client.InvokeApi(changeresultcallback_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ChangeresultcallbackResponse
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

func (response *ChangeresultcallbackResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

