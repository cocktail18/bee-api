package waimaipoiinfo

/**
* 获取商家基本信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const waimai_poi_info_url = "/wmoper/ng/waimaiad/common/wmpoiinfo"

type WaimaiPoiInfoResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type WaimaiPoiInfoRequest struct {
}
type Data struct {
    Result Result `json:"result"`
    WmPoiInfo WmPoiInfo `json:"wmPoiInfo"`
}
type WmPoiInfo struct {
    FirstCityName string `json:"firstCityName"`
    SecondCityName string `json:"secondCityName"`
    ThirdCityName string `json:"thirdCityName"`
    FirstTagName string `json:"firstTagName"`
    SecondTagName string `json:"secondTagName"`
    ThirdTagName string `json:"thirdTagName"`
}
type Result struct {
    Success bool `json:"success"`
    Code string `json:"code"`
    Message string `json:"message"`
}



func (req *WaimaiPoiInfoRequest) DoInvoke(client mtclient.MeituanClient) (*WaimaiPoiInfoResponse, error) {
    resp, err := client.InvokeApi(waimai_poi_info_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response WaimaiPoiInfoResponse
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

func (response *WaimaiPoiInfoResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

