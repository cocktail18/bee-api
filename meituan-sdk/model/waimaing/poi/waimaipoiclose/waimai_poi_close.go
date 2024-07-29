package waimaipoiclose

/**
* 门店置休息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const waimai_poi_close_url = "/waimai/poi/close"

type WaimaiPoiCloseResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type WaimaiPoiCloseRequest struct {
}



func (req *WaimaiPoiCloseRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*WaimaiPoiCloseResponse, error) {
    resp, err := client.InvokeApi(waimai_poi_close_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response WaimaiPoiCloseResponse
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

func (response *WaimaiPoiCloseResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

