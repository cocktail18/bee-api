package peeradbidprice

/**
* 竞价推广-附近商家平均出价
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const peer_ad_bidprice_url = "/wmoper/ng/waimaiad/peer/bidprice"

type PeerAdBidpriceResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type PeerAdBidpriceRequest struct {
}
type Data struct {
    Result Result `json:"result"`
    AveragePrice string `json:"averagePrice"`
}
type Result struct {
    Success bool `json:"success"`
    Code string `json:"code"`
    Message string `json:"message"`
}



func (req *PeerAdBidpriceRequest) DoInvoke(client mtclient.MeituanClient) (*PeerAdBidpriceResponse, error) {
    resp, err := client.InvokeApi(peer_ad_bidprice_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response PeerAdBidpriceResponse
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

func (response *PeerAdBidpriceResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

