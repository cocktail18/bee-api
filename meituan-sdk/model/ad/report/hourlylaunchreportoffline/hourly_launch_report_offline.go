package hourlylaunchreportoffline

/**
* 推广分时报告
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const hourly_launch_report_offline_url = "/ad/report/getHourlyDataByLaunchOffline"

type HourlyLaunchReportOfflineResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 对象列表
    */
    Data []LaunchInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type HourlyLaunchReportOfflineRequest struct {
    /**
    * 推广id列表,单次最多支持50个
    */
    LaunchIds []int64 `json:"launchIds"`
    /**
    * 平台
    */
    Platform int32 `json:"platform"`
    /**
    * 开始时间
    */
    BeginTime string `json:"beginTime"`
    /**
    * 结束时间
    */
    EndTime string `json:"endTime"`
}
type LaunchInfo struct {
    /**
    * 日期
    */
    Date string `json:"date"`
    /**
    * 小时
    */
    Hour int32 `json:"hour"`
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



func (req *HourlyLaunchReportOfflineRequest) DoInvoke(client mtclient.MeituanClient) (*HourlyLaunchReportOfflineResponse, error) {
    resp, err := client.InvokeApi(hourly_launch_report_offline_url, 22, "", req)

    if err != nil {
        return nil, err
    }
    var response HourlyLaunchReportOfflineResponse
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

func (response *HourlyLaunchReportOfflineResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

