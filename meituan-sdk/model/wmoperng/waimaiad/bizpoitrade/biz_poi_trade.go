package bizpoitrade

/**
* 经营分析-商家交易信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const biz_poi_trade_url = "/wmoper/ng/waimaiad/biz/poitrade"

type BizPoiTradeRequest struct {
    /**
    * 日期，yyyyMMdd格式
    */
    Dt string `json:"dt"`
}
type BizPoiTradeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type PoiTrade struct {
    OriginalPrice int64 `json:"originalPrice"`
    ActualPrice int64 `json:"actualPrice"`
    OrdNum int64 `json:"ordNum"`
    ReceiveOrdNum int64 `json:"receiveOrdNum"`
    Dt int64 `json:"dt"`
}
type Data struct {
    Result Result `json:"result"`
    PoiTrade PoiTrade `json:"poiTrade"`
}
type Result struct {
    Success bool `json:"success"`
    Code string `json:"code"`
    Message string `json:"message"`
}



func (req *BizPoiTradeRequest) DoInvoke(client mtclient.MeituanClient) (*BizPoiTradeResponse, error) {
    resp, err := client.InvokeApi(biz_poi_trade_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response BizPoiTradeResponse
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

func (response *BizPoiTradeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

