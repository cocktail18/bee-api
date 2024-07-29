package querylivegoodsinfo

/**
* 查询直播间商品
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_live_goods_info_url = "/mlive/goods/query"

type QueryLiveGoodsInfoRequest struct {
    /**
    *  直播场次Id 
    */
    LiveId int64 `json:"liveId"`
}
type QueryLiveGoodsInfoResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []LiveRoomGoodsDTO `json:"data"`
    TraceId string `json:"traceId"`
}
type LiveRoomGoodsDTO struct {
    /**
    * 直播id
    */
    LiveId int64 `json:"liveId"`
    /**
    * 商品类型
    */
    GoodsType int32 `json:"goodsType"`
    /**
    * 商品id
    */
    GoodsId string `json:"goodsId"`
    /**
    * 商品标题
    */
    GoodsTitle string `json:"goodsTitle"`
    /**
    * 序号（直播中当前的序号）
    */
    Rank int32 `json:"rank"`
    /**
    * 是否置顶
    */
    Top bool `json:"top"`
}



func (req *QueryLiveGoodsInfoRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryLiveGoodsInfoResponse, error) {
    resp, err := client.InvokeApi(query_live_goods_info_url, 50, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryLiveGoodsInfoResponse
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

func (response *QueryLiveGoodsInfoResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

