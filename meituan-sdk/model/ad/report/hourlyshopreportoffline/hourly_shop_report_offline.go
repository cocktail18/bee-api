package hourlyshopreportoffline

/**
* 门店分时报告
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const hourly_shop_report_offline_url = "/ad/report/getHourlyDataByShopOffline"

type HourlyShopReportOfflineRequest struct {
    /**
    * 门店id列表,单次最多支持50个
    */
    ShopIds []int64 `json:"shopIds"`
    /**
    * 点评门店id就还是美团门店id，默认点评门店id
    */
    ShopType int32 `json:"shopType"`
    /**
    * 查询开始时间
    */
    BeginTime string `json:"beginTime"`
    /**
    * 查询截止时间
    */
    EndTime string `json:"endTime"`
}
type HourlyShopReportOfflineResponse struct {
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
    * 天数
    */
    Date string `json:"date"`
    /**
    * 小时
    */
    Hour int32 `json:"hour"`
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



func (req *HourlyShopReportOfflineRequest) DoInvoke(client mtclient.MeituanClient) (*HourlyShopReportOfflineResponse, error) {
    resp, err := client.InvokeApi(hourly_shop_report_offline_url, 22, "", req)

    if err != nil {
        return nil, err
    }
    var response HourlyShopReportOfflineResponse
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

func (response *HourlyShopReportOfflineResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

