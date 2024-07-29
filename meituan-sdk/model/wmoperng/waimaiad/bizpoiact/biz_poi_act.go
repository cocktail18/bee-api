package bizpoiact

/**
* 经营分析-商家活动信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const biz_poi_act_url = "/wmoper/ng/waimaiad/biz/poiact"

type PoiActList struct {
    WmActType int64 `json:"wmActType"`
    Detail string `json:"detail"`
    Priority int64 `json:"priority"`
    StartTime int64 `json:"startTime"`
    EndTime int64 `json:"endTime"`
    Dt int64 `json:"dt"`
}
type Data struct {
    Result Result `json:"result"`
    PoiActList []PoiActList `json:"poiActList"`
}
type BizPoiActRequest struct {
    /**
    * 日期，yyyyMMdd格式
    */
    Dt string `json:"dt"`
}
type BizPoiActResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Result struct {
    Success bool `json:"success"`
    Code string `json:"code"`
    Message string `json:"message"`
}



func (req *BizPoiActRequest) DoInvoke(client mtclient.MeituanClient) (*BizPoiActResponse, error) {
    resp, err := client.InvokeApi(biz_poi_act_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response BizPoiActResponse
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

func (response *BizPoiActResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

