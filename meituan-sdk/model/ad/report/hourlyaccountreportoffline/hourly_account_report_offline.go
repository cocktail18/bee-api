package hourlyaccountreportoffline

/**
* 账户分时报告
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const hourly_account_report_offline_url = "/ad/report/getHourlyDataByAccountOffline"

type HourlyAccountReportOfflineResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 对象列表
    */
    Data []Data `json:"data"`
    TraceId string `json:"traceId"`
}
type HourlyAccountReportOfflineRequest struct {
    /**
    * 查询范围开始时间
    */
    BeginTime string `json:"beginTime"`
    /**
    * 查询范围截止时间
    */
    EndTime string `json:"endTime"`
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
    Click int32 `json:"click"`
}



func (req *HourlyAccountReportOfflineRequest) DoInvoke(client mtclient.MeituanClient) (*HourlyAccountReportOfflineResponse, error) {
    resp, err := client.InvokeApi(hourly_account_report_offline_url, 22, "", req)

    if err != nil {
        return nil, err
    }
    var response HourlyAccountReportOfflineResponse
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

func (response *HourlyAccountReportOfflineResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

