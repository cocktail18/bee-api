package cpmrtdata

/**
* getCpmRtData
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const cpm_rtdata_url = "/ad/report/getCpmRtData"

type CpmRtdataResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 对象列表
    */
    Data []LaunchInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type CpmRtdataRequest struct {
    /**
    * 对应查询元素的id列表,单次最多支持50个
    */
    Ids []int64 `json:"ids"`
    /**
    * 0账户，1计划，2品牌，3推广 4 创意
    */
    Dimension int32 `json:"dimension"`
    /**
    * 子账户
    */
    AccountId int32 `json:"accountId"`
}
type LaunchInfo struct {
    /**
    * 日期
    */
    Date string `json:"date"`
    /**
    * 查询元素id
    */
    Id int32 `json:"id"`
    /**
    * 查询元素名称
    */
    Name string `json:"name"`
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



func (req *CpmRtdataRequest) DoInvoke(client mtclient.MeituanClient) (*CpmRtdataResponse, error) {
    resp, err := client.InvokeApi(cpm_rtdata_url, 22, "", req)

    if err != nil {
        return nil, err
    }
    var response CpmRtdataResponse
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

func (response *CpmRtdataResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

