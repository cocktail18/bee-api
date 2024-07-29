package adappbuyinfo

/**
* 获取商家购买的应用信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const ad_app_buy_info_url = "/wmoper/ng/waimaiad/buyinfo"

type AdAppBuyInfoRequest struct {
}
type Data struct {
    Result Result `json:"result"`
    Valid int64 `json:"valid"`
    Version string `json:"version"`
    Days int64 `json:"days"`
    StartTime int64 `json:"startTime"`
    EndTime int64 `json:"endTime"`
}
type AdAppBuyInfoResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Result struct {
    Success bool `json:"success"`
    Code string `json:"code"`
    Message string `json:"message"`
}



func (req *AdAppBuyInfoRequest) DoInvoke(client mtclient.MeituanClient) (*AdAppBuyInfoResponse, error) {
    resp, err := client.InvokeApi(ad_app_buy_info_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response AdAppBuyInfoResponse
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

func (response *AdAppBuyInfoResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

