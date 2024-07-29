package groupdeliverypoidealquery

/**
* 2.1.8.查询门店团购配送套餐列表
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const group_delivery_poi_deal_query_url = "/dcps/fulfill/poi/deal/query"

type Data struct {
    /**
    * 团单id
    */
    DealId int32 `json:"dealId"`
    /**
    * 市场价格
    */
    DealValue float64 `json:"dealValue"`
    /**
    * 团单名称
    */
    DealTitle string `json:"dealTitle"`
    /**
    * 团单开始售卖时间戳，单位秒
    */
    BeginTime int64 `json:"beginTime"`
    /**
    * 售卖价格
    */
    DealPrice float64 `json:"dealPrice"`
    /**
    * 团单状态
    */
    DealStatus int32 `json:"dealStatus"`
    /**
    * 团单结束售卖时间戳，单位秒
    */
    EndTime int64 `json:"endTime"`
    /**
    * json字符串数组
    */
    DealMenu string `json:"dealMenu"`
    /**
    * 是否是代金券，true为代金券，false为套餐券
    */
    IsVoucher bool `json:"isVoucher"`
    /**
    * 第三方Id(商家商品ID)
    */
    ProductThirdPartyId string `json:"productThirdPartyId"`
}
type GroupDeliveryPoiDealQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []Data `json:"data"`
    TraceId string `json:"traceId"`
}
type GroupDeliveryPoiDealQueryRequest struct {
    /**
    *  团单状态，默认查询1/正在售卖状态团单 -1 / 全部状态（含以下所有） 0 / 结束售卖 1 / 正在售卖 2 / 隐藏单，前台通过POI列表及deal列表都不可见，但是通过收藏或url可以直接访问并购买 4 / 未开始售卖 5 / 不可购买 
    */
    DealStatus int32 `json:"dealStatus"`
    /**
    *  一次查询不超过50个，偏移量（&lt;=50） 
    */
    Limit int32 `json:"limit"`
    /**
    *  起始点 
    */
    Offset int32 `json:"offset"`
}



func (req *GroupDeliveryPoiDealQueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupDeliveryPoiDealQueryResponse, error) {
    resp, err := client.InvokeApi(group_delivery_poi_deal_query_url, 46, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupDeliveryPoiDealQueryResponse
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

func (response *GroupDeliveryPoiDealQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

