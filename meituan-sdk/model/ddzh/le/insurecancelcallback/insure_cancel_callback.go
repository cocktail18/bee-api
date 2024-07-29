package insurecancelcallback

/**
* 退保回调接口
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const insure_cancel_callback_url = "/ddzh/le/insurance/insure/cancel/callback"

type InsureCancelCallbackRequest struct {
    Biz string `json:"biz"`
}
type InsureCancelCallbackResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *InsureCancelCallbackRequest) DoInvoke(client mtclient.MeituanClient) (*InsureCancelCallbackResponse, error) {
    resp, err := client.InvokeApi(insure_cancel_callback_url, 58, "", req)

    if err != nil {
        return nil, err
    }
    var response InsureCancelCallbackResponse
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

func (response *InsureCancelCallbackResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

