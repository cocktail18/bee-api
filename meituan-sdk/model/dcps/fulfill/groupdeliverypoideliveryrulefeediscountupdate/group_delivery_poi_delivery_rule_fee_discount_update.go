package groupdeliverypoideliveryrulefeediscountupdate

/**
* 门店配送规则更新配送费折扣信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const group_delivery_poi_delivery_rule_fee_discount_update_url = "/dcps/fulfill/poi/deliveryRule/feeDiscount/update"

type GroupDeliveryPoiDeliveryRuleFeeDiscountUpdateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}
type TieredDiscount struct {
    Discount string `json:"discount"`
    EndDistance string `json:"endDistance"`
    StartDistance string `json:"startDistance"`
    Tiered int32 `json:"tiered"`
}
type GroupDeliveryPoiDeliveryRuleFeeDiscountUpdateRequest struct {
    DiscountSwitch bool `json:"discountSwitch"`
    DiscountType int32 `json:"discountType"`
    TieredDiscounts []TieredDiscount `json:"tieredDiscounts"`
    UnifiedDiscount string `json:"unifiedDiscount"`
}



func (req *GroupDeliveryPoiDeliveryRuleFeeDiscountUpdateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupDeliveryPoiDeliveryRuleFeeDiscountUpdateResponse, error) {
    resp, err := client.InvokeApi(group_delivery_poi_delivery_rule_fee_discount_update_url, 46, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupDeliveryPoiDeliveryRuleFeeDiscountUpdateResponse
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

func (response *GroupDeliveryPoiDeliveryRuleFeeDiscountUpdateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

