package queryadaccountbalance

/**
* 查询账号余额
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_adaccount_balance_url = "/ad/launch/queryAdAccountBalance"

type ProductAccountInfo struct {
}
type QueryAdaccountBalanceRequest struct {
    SonAdaccount int32 `json:"sonAdaccount"`
}
type AdAccountBalance struct {
    /**
    * 账号总余额
    */
    Balance string `json:"balance"`
    /**
    * 各个资金池明细
    */
    ProductAccountInfoList []ProductAccountInfo `json:"productAccountInfoList"`
}
type QueryAdaccountBalanceResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * validShopIds合法的门店列表，invalidShopIds无法投放的门店
    */
    Data AdAccountBalance `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *QueryAdaccountBalanceRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryAdaccountBalanceResponse, error) {
    resp, err := client.InvokeApi(query_adaccount_balance_url, 22, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryAdaccountBalanceResponse
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

func (response *QueryAdaccountBalanceResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

