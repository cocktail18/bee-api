package wmoperqueryboxpricetypeanddetail

/**
* 查询门店打包费
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const wmoper_query_box_price_type_and_detail_url = "/wmoper/ng/food/queryBoxPriceTypeAndDetail"

type WmoperQueryBoxPriceTypeAndDetailResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data BoxPrice `json:"data"`
    TraceId string `json:"traceId"`
}
type BoxPriceRulesVo struct {
    /**
    * 阶梯最高价
    */
    EndPrice float32 `json:"endPrice"`
    /**
    * 该阶梯内的打包费
    */
    LadderPrice float32 `json:"ladderPrice"`
    /**
    * 阶梯起始价
    */
    StartPrice float32 `json:"startPrice"`
}
type BoxPriceDetail struct {
    /**
    * 打包费收取方式 1为按阶梯价格收费, 2为按一口价收费
    */
    BoxPriceRuleType int32 `json:"boxPriceRuleType"`
    /**
    * 一口价价格
    */
    FixedPrice float32 `json:"fixedPrice"`
    /**
    * 阶梯规则, 要求规则数量最少两条最多十条
    */
    BoxPriceRulesVoList []BoxPriceRulesVo `json:"boxPriceRulesVoList"`
}
type BoxPrice struct {
    BoxPriceDetail BoxPriceDetail `json:"boxPriceDetail"`
    /**
    * 打包费类型, 0为按商品收费, 1 为按订单收费, 2为按口袋收费
    */
    BoxPriceType int32 `json:"boxPriceType"`
}
type WmoperQueryBoxPriceTypeAndDetailRequest struct {
}



func (req *WmoperQueryBoxPriceTypeAndDetailRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*WmoperQueryBoxPriceTypeAndDetailResponse, error) {
    resp, err := client.InvokeApi(wmoper_query_box_price_type_and_detail_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response WmoperQueryBoxPriceTypeAndDetailResponse
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

func (response *WmoperQueryBoxPriceTypeAndDetailResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

