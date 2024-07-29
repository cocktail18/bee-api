package hourlycpcrtdatabyaccount

/**
* cpc账户实时分时数据
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const hourly_cpc_rt_data_byaccount_url = "/ad/report/getHourlyCpcRtDataByAccount"

type HourlyCpcRtDataByaccountResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []Data `json:"data"`
    TraceId string `json:"traceId"`
}
type HourlyCpcRtDataByaccountRequest struct {
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
    * 小时
    */
    Hour int32 `json:"hour"`
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



func (req *HourlyCpcRtDataByaccountRequest) DoInvoke(client mtclient.MeituanClient) (*HourlyCpcRtDataByaccountResponse, error) {
    resp, err := client.InvokeApi(hourly_cpc_rt_data_byaccount_url, 22, "", req)

    if err != nil {
        return nil, err
    }
    var response HourlyCpcRtDataByaccountResponse
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

func (response *HourlyCpcRtDataByaccountResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

