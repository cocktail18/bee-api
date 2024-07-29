package livegoodsontop

/**
* 直播间商品置顶
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const live_goods_on_top_url = "/mlive/goods/onTop"

type LiveGoodsOnTopResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 布尔值，code=0时，返回true
    */
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}
type LiveGoodsOnTopRequest struct {
    /**
    *  业务商品id 
    */
    OutGoodsId string `json:"outGoodsId"`
    /**
    *  商品id，与outGoodsId传任意一个，goodsId优先 
    */
    GoodsId string `json:"goodsId"`
    /**
    *  数字人形象id 
    */
    DigitalImageId string `json:"digitalImageId"`
    /**
    *  直播场次id 
    */
    LiveId int64 `json:"liveId"`
    /**
    *  商品类型 
    */
    GoodsType int64 `json:"goodsType"`
}



func (req *LiveGoodsOnTopRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*LiveGoodsOnTopResponse, error) {
    resp, err := client.InvokeApi(live_goods_on_top_url, 50, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response LiveGoodsOnTopResponse
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

func (response *LiveGoodsOnTopResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

