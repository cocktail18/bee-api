package updatestock

/**
* 更新三方库存
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const updatestock_url = "/ddzh/yuding/updatestock"

type UpdatestockRequest struct {
    /**
    *  三方房间ID 
    */
    ThirdPartyRoomId int64 `json:"thirdPartyRoomId"`
    /**
    *  房间名称 
    */
    RoomName string `json:"roomName"`
    /**
    *  时段库存信息 注：类型为对象类型，详见扩展参数desk_sold_time_periods 
    */
    DeskSoldTimePeriods []DeskSoldTimePeriodsSub `json:"deskSoldTimePeriods"`
}
type TimePeriodItemsSub struct {
    /**
    *  开始时间（从0:00开始至开始时间的分钟数，如8:00开始即为 480 【8*60】） 
    */
    BeginMinutes int32 `json:"beginMinutes"`
    /**
    *  结束时间（从0:00开始至结束的分钟数） 
    */
    EndMinutes int32 `json:"endMinutes"`
    /**
    *  开始日期（时间戳毫秒数） 
    */
    BeginTime int64 `json:"beginTime"`
    /**
    *  结束时间（从0:00开始至结束的分钟数） 
    */
    EndTime int64 `json:"endTime"`
}
type UpdatestockResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}
type DeskSoldTimePeriodsSub struct {
    /**
    *  业务类型，1-普通预订，2-包座，3-押金预订 
    */
    BizType string `json:"bizType"`
    /**
    *  占用时段信息 注：类型为对象类型，详见扩展参数time_period_items 
    */
    TimePeriodItems []TimePeriodItemsSub `json:"timePeriodItems"`
}



func (req *UpdatestockRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*UpdatestockResponse, error) {
    resp, err := client.InvokeApi(updatestock_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response UpdatestockResponse
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

func (response *UpdatestockResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

