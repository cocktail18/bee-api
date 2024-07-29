package poiadeffectdata

/**
* 营业分析-查询流量曝光与入店数据
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const poi_ad_effectdata_url = "/wmoper/ng/waimaiad/poi/effectdata"

type Item struct {
    Dt int64 `json:"dt"`
    Data []Prop `json:"data"`
}
type Prop struct {
    FlowProp int32 `json:"flowProp"`
    Pos string `json:"pos"`
    Show string `json:"show"`
    Click string `json:"click"`
}
type PoiAdEffectdataRequest struct {
    /**
    * 流量属性-1 全部 1自然流量 2付费流量 3全部流量
    */
    FlowProp int32 `json:"flowProp"`
    /**
    * 最新N日内数据，[1-31]
    */
    Days int32 `json:"days"`
}
type EffectdataResponse struct {
    Result Result `json:"result"`
    Data []Item `json:"data"`
}
type PoiAdEffectdataResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data EffectdataResponse `json:"data"`
    TraceId string `json:"traceId"`
}
type Result struct {
    Success bool `json:"success"`
    Code string `json:"code"`
    Message string `json:"message"`
}



func (req *PoiAdEffectdataRequest) DoInvoke(client mtclient.MeituanClient) (*PoiAdEffectdataResponse, error) {
    resp, err := client.InvokeApi(poi_ad_effectdata_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response PoiAdEffectdataResponse
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

func (response *PoiAdEffectdataResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

