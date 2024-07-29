package waimaikaidianbizsettle

/**
* 获取美团外卖开店绿色通道链接地址
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const waimai_kaidian_bizsettle_url = "/waimai/kaidian/bizsettle"

type WaimaiKaidianBizsettleResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data WaimaiKaidianBizsettleData `json:"data"`
    TraceId string `json:"traceId"`
}
type WaimaiKaidianBizsettleRequest struct {
    DeveloperId int64 `json:"developerId"`
    /**
    *  服务商门店id 
    */
    EPoiId string `json:"ePoiId"`
}
type WaimaiKaidianBizsettleData struct {
    /**
    * 美团外卖开店绿色通道链接地址
    */
    Url string `json:"url"`
}



func (req *WaimaiKaidianBizsettleRequest) DoInvoke(client mtclient.MeituanClient) (*WaimaiKaidianBizsettleResponse, error) {
    resp, err := client.InvokeApi(waimai_kaidian_bizsettle_url, 2, "", req)

    if err != nil {
        return nil, err
    }
    var response WaimaiKaidianBizsettleResponse
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

func (response *WaimaiKaidianBizsettleResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

