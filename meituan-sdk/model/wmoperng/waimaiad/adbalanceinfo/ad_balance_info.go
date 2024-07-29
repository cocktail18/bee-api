package adbalanceinfo

/**
* 获取商家广告余额
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const ad_balance_info_url = "/wmoper/ng/waimaiad/balance"

type AdBalanceInfoResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type AdBalanceInfoRequest struct {
}
type Data struct {
    Result Result `json:"result"`
    PrimaryBalance string `json:"primaryBalance"`
    CommonBalance string `json:"commonBalance"`
    Debit string `json:"debit"`
}
type Result struct {
    Success bool `json:"success"`
    Code string `json:"code"`
    Message string `json:"message"`
}



func (req *AdBalanceInfoRequest) DoInvoke(client mtclient.MeituanClient) (*AdBalanceInfoResponse, error) {
    resp, err := client.InvokeApi(ad_balance_info_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response AdBalanceInfoResponse
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

func (response *AdBalanceInfoResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

