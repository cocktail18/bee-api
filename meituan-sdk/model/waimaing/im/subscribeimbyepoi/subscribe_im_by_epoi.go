package subscribeimbyepoi

/**
* 根据门店订阅外卖IM消息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const subscribe_im_by_epoi_url = "/waimai/ng/im/subscribeImByEpoi"

type SubscribeImByEpoiRequest struct {
}
type SubscribeImByEpoiResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}



func (req *SubscribeImByEpoiRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*SubscribeImByEpoiResponse, error) {
    resp, err := client.InvokeApi(subscribe_im_by_epoi_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response SubscribeImByEpoiResponse
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

func (response *SubscribeImByEpoiResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

