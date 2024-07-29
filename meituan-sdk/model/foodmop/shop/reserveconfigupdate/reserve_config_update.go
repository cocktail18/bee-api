package reserveconfigupdate

/**
* 门店“线上点”预约配置（非必接）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const reserve_config_update_url = "/foodmop/shop/reserve/config/update"

type ReserveConfigUpdateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * Key：vendorShopId（品牌门店Id） Value：OperateResulst（操作结果）
    */
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}
type ReserveConfigUpdateRequest struct {
    /**
    *  厂商配置列表，批量限制20 
    */
    VendorReserveConfigDTOS []VendorReserveConfigRequest `json:"vendorReserveConfigDTOS"`
}
type ResultData struct {
}
type MopReserveHour struct {
}
type VendorReserveConfigRequest struct {
    /**
    *  厂商门店ID 
    */
    VendorShopId string `json:"vendorShopId"`
    /**
    *  是否支持预约： TRUE- 支持预约 FALSE- 不支持预约 
    */
    MopReserveOpen bool `json:"mopReserveOpen"`
    /**
    *  mopReserveOpen为true时必填 mopReserveOpen为true时必填 线上点可预约时间 { "1"://周一,DayofWeek [{ "startTime":"08:00", //Time "endTime":"12:00" }], "2"://周二 [{ "startTime":"08:00", "endTime":"18:00" }] } 
    */
    MopReserveHour MopReserveHour `json:"mopReserveHour"`
    /**
    *  mopReserveOpen为true时必填 预约周期（单位分钟），用户预约选择时间的间隔 eg：配置5min，则每小时的0,5,10, ... eg：配置7min，则每小时的0,7,14,... 
    */
    MopReservePeriod int32 `json:"mopReservePeriod"`
    /**
    *  mopReserveOpen为true时必填 距离现在的最早可预约时间分钟数（单位分钟） 
    */
    EarliesReserveTimeFromNowByMinutes int32 `json:"earliesReserveTimeFromNowByMinutes"`
    /**
    *  距离现在的最晚可预约时间分钟数（单位分钟） 
    */
    LatestReserveTimeFromNowByMinutes int32 `json:"latestReserveTimeFromNowByMinutes"`
}



func (req *ReserveConfigUpdateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ReserveConfigUpdateResponse, error) {
    resp, err := client.InvokeApi(reserve_config_update_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ReserveConfigUpdateResponse
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

func (response *ReserveConfigUpdateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

