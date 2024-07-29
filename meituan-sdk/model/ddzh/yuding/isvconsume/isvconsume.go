package isvconsume

/**
* 用户到店核销
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const isvconsume_url = "/ddzh/yuding/isvconsume"

type Data struct {
    /**
    * 平台订单id
    */
    OrderId string `json:"orderId"`
}
type IsvconsumeRequest struct {
    /**
    *  技术合作中心订单id 
    */
    OrderId string `json:"orderId"`
}
type IsvconsumeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *IsvconsumeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*IsvconsumeResponse, error) {
    resp, err := client.InvokeApi(isvconsume_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response IsvconsumeResponse
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

func (response *IsvconsumeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

