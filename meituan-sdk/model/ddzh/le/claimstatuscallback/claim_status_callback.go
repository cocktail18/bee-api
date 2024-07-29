package claimstatuscallback

/**
* 理赔状态回调接口
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const claim_status_callback_url = "/ddzh/le/insurance/claim/status/callback"

type ClaimStatusCallbackResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type ClaimStatusCallbackRequest struct {
    Biz string `json:"biz"`
}



func (req *ClaimStatusCallbackRequest) DoInvoke(client mtclient.MeituanClient) (*ClaimStatusCallbackResponse, error) {
    resp, err := client.InvokeApi(claim_status_callback_url, 58, "", req)

    if err != nil {
        return nil, err
    }
    var response ClaimStatusCallbackResponse
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

func (response *ClaimStatusCallbackResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

