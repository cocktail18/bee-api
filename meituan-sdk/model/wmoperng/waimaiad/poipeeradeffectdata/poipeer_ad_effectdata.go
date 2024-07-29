package poipeeradeffectdata

/**
* 竞价推广-商圈与门店效果数据-新客进店
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const poipeer_ad_effectdata_url = "/wmoper/ng/waimaiad/poipeer/effectdata"

type NewUserClickRate struct {
    Dt int64 `json:"dt"`
    Poi string `json:"poi"`
    Peer string `json:"peer"`
}
type PoipeerAdEffectdataResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type PoipeerAdEffectdataRequest struct {
    /**
    * 最新N日内数据，[1-31]
    */
    Days int32 `json:"days"`
}
type Data struct {
    Result Result `json:"result"`
    NewUserClickRate []NewUserClickRate `json:"newUserClickRate"`
}
type Result struct {
    Success bool `json:"success"`
    Code string `json:"code"`
    Message string `json:"message"`
}



func (req *PoipeerAdEffectdataRequest) DoInvoke(client mtclient.MeituanClient) (*PoipeerAdEffectdataResponse, error) {
    resp, err := client.InvokeApi(poipeer_ad_effectdata_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response PoipeerAdEffectdataResponse
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

func (response *PoipeerAdEffectdataResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

