package peeradeffectdata

/**
* 竞价推广-商圈效果数据-曝光与点击
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const peer_ad_effectdata_url = "/wmoper/ng/waimaiad/peer/effectdata"

type Data struct {
    Dt int64 `json:"dt"`
    Show string `json:"show"`
    Click string `json:"click"`
}
type PeerAdEffectdataResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []Data `json:"data"`
    TraceId string `json:"traceId"`
}
type PeerAdEffectdataRequest struct {
    /**
    * 最新N日内数据，[1-31]
    */
    Days int32 `json:"days"`
}



func (req *PeerAdEffectdataRequest) DoInvoke(client mtclient.MeituanClient) (*PeerAdEffectdataResponse, error) {
    resp, err := client.InvokeApi(peer_ad_effectdata_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response PeerAdEffectdataResponse
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

func (response *PeerAdEffectdataResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

