package adstatusinfo

/**
* 获取广告计划状态，包括是否开启、预算、出价
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const ad_status_info_url = "/wmoper/ng/waimaiad/status"

type AdStatusInfoResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Data struct {
    Result Result `json:"result"`
    Channel int64 `json:"channel"`
    Status int64 `json:"status"`
    Budget int64 `json:"budget"`
    Bid int64 `json:"bid"`
    BudgetUsed string `json:"budgetUsed"`
}
type AdStatusInfoRequest struct {
    /**
    * 外卖广告的产品类型，1表示点金推广
    */
    Channel int32 `json:"channel"`
}
type Result struct {
    Success bool `json:"success"`
    Code string `json:"code"`
    Message string `json:"message"`
}



func (req *AdStatusInfoRequest) DoInvoke(client mtclient.MeituanClient) (*AdStatusInfoResponse, error) {
    resp, err := client.InvokeApi(ad_status_info_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response AdStatusInfoResponse
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

func (response *AdStatusInfoResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

