package cpcrtdatabylaunch

/**
* cpc推广实时天数据
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const cpc_rtdata_bylaunch_url = "/ad/report/getCpcRtDataByLaunch"

type CpcRtdataBylaunchRequest struct {
    /**
    * 推广id列表,单次最多支持50个
    */
    LaunchIds []int64 `json:"launchIds"`
    /**
    * 平台，默认全部
    */
    Platform int32 `json:"platform"`
    /**
    * 子账号id
    */
    AccountId int32 `json:"accountId"`
}
type CpcRtdataBylaunchResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 对象列表
    */
    Data []LaunchInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type LaunchInfo struct {
    /**
    * 日期
    */
    Date string `json:"date"`
    /**
    * 推广id
    */
    LaunchId int32 `json:"launch_id"`
    /**
    * 推广名称
    */
    LaunchName string `json:"launch_name"`
    /**
    * 门店名称
    */
    ShopName string `json:"shop_name"`
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
    Click int32 `json:"click"`
}



func (req *CpcRtdataBylaunchRequest) DoInvoke(client mtclient.MeituanClient) (*CpcRtdataBylaunchResponse, error) {
    resp, err := client.InvokeApi(cpc_rtdata_bylaunch_url, 22, "", req)

    if err != nil {
        return nil, err
    }
    var response CpcRtdataBylaunchResponse
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

func (response *CpcRtdataBylaunchResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

