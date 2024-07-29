package unsubscribeimbyepoi

/**
* 根据门店解除订阅外卖IM消息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const unsubscribe_im_by_epoi_url = "/waimai/ng/im/unsubscribeImByEpoi"

type UnsubscribeImByEpoiResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}
type UnsubscribeImByEpoiRequest struct {
}



func (req *UnsubscribeImByEpoiRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*UnsubscribeImByEpoiResponse, error) {
    resp, err := client.InvokeApi(unsubscribe_im_by_epoi_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response UnsubscribeImByEpoiResponse
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

func (response *UnsubscribeImByEpoiResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

