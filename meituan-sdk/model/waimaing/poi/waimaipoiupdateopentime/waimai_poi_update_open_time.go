package waimaipoiupdateopentime

/**
* 修改门店营业时间
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const waimai_poi_update_open_time_url = "/waimai/poi/updateOpenTime"

type WaimaiPoiUpdateOpenTimeRequest struct {
    /**
    *  营业时间时间段，且支持多个时间段。 如果一天多个时段，可以用逗号隔开， 比如：01:00-02:00,03:00-04:00,06:00-08:00， 时间段的数量不限 不能修改具体哪天是否营业 
    */
    OpenTime string `json:"openTime"`
}
type WaimaiPoiUpdateOpenTimeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 操作结果
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *WaimaiPoiUpdateOpenTimeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*WaimaiPoiUpdateOpenTimeResponse, error) {
    resp, err := client.InvokeApi(waimai_poi_update_open_time_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response WaimaiPoiUpdateOpenTimeResponse
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

func (response *WaimaiPoiUpdateOpenTimeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

