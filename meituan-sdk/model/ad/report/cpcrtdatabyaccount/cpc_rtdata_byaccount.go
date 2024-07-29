package cpcrtdatabyaccount

/**
* cpc账户实时天数据
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const cpc_rtdata_byaccount_url = "/ad/report/getCpcRtDataByAccount"

type CpcRtdataByaccountResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []Data `json:"data"`
    TraceId string `json:"traceId"`
}
type CpcRtdataByaccountRequest struct {
    /**
    * 子账号id
    */
    AccountId int32 `json:"accountId"`
}
type Data struct {
    /**
    * 天
    */
    Date string `json:"date"`
    /**
    * 花费
    */
    Charge float64 `json:"charge"`
    /**
    * 曝光
    */
    Imp int32 `json:"imp"`
    /**
    * 点击
    */
    Click string `json:"click"`
}



func (req *CpcRtdataByaccountRequest) DoInvoke(client mtclient.MeituanClient) (*CpcRtdataByaccountResponse, error) {
    resp, err := client.InvokeApi(cpc_rtdata_byaccount_url, 22, "", req)

    if err != nil {
        return nil, err
    }
    var response CpcRtdataByaccountResponse
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

func (response *CpcRtdataByaccountResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

