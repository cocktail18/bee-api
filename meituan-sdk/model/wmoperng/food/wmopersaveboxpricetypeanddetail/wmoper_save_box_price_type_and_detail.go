package wmopersaveboxpricetypeanddetail

/**
* 设置门店打包费
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const wmoper_save_box_price_type_and_detail_url = "/wmoper/ng/food/saveBoxPriceTypeAndDetail"

type BoxPriceRulesVo struct {
    /**
    *  阶梯最高价（最后一个阶梯的最高价用-1标识正无穷） 
    */
    EndPrice float32 `json:"endPrice"`
    /**
    *  该阶梯内的打包费 
    */
    LadderPrice float32 `json:"ladderPrice"`
    /**
    *  阶梯起始价 (1.后一个阶梯的起始价需要和前一个阶梯的最高价相同; 2.第一个阶梯起始价必须为0) 
    */
    StartPrice float32 `json:"startPrice"`
}
type WmoperSaveBoxPriceTypeAndDetailResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type BoxPriceDetail struct {
    /**
    *  打包费收取方式 1为按阶梯价格收费, 2为按一口价收费 
    */
    BoxPriceRuleType int32 `json:"boxPriceRuleType"`
    /**
    *  一口价价格(boxPriceRulesVoList不为空时，该字段不传或传值为0) 
    */
    FixedPrice float32 `json:"fixedPrice"`
    /**
    *  阶梯规则 (1. 要求规则数量最少两条最多十条; 2. fixedPrice&gt;0时，该字段不传或为空) 
    */
    BoxPriceRulesVoList []BoxPriceRulesVo `json:"boxPriceRulesVoList"`
}
type WmoperSaveBoxPriceTypeAndDetailRequest struct {
    /**
    *  打包费信息(boxPriceType为0时该字段可不传) 
    */
    BoxPriceDetail BoxPriceDetail `json:"boxPriceDetail"`
    /**
    *  打包费类型, 0为按商品收费, 1 为按订单收费, 2为按口袋收费 
    */
    BoxPriceType int32 `json:"boxPriceType"`
}



func (req *WmoperSaveBoxPriceTypeAndDetailRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*WmoperSaveBoxPriceTypeAndDetailResponse, error) {
    resp, err := client.InvokeApi(wmoper_save_box_price_type_and_detail_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response WmoperSaveBoxPriceTypeAndDetailResponse
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

func (response *WmoperSaveBoxPriceTypeAndDetailResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

