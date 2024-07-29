package groupdeliverypoideliveryrulecreate

/**
* 门店配送规则创建
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const group_delivery_poi_delivery_rule_create_url = "/dcps/fulfill/poi/deliveryRule/create"

type DeliveryFeeDiscount struct {
    DiscountSwitch bool `json:"discountSwitch"`
    DiscountType int32 `json:"discountType"`
    TieredDiscounts []TieredDiscount `json:"tieredDiscounts"`
    UnifiedDiscount string `json:"unifiedDiscount"`
}
type DeliveryHours struct {
    BeginTime string `json:"beginTime"`
    EndTime string `json:"endTime"`
}
type DeliveryReservation struct {
    AvailableDays []int32 `json:"availableDays"`
    BeyondDeliveryTimeOrder bool `json:"beyondDeliveryTimeOrder"`
    ReservationRemind int32 `json:"reservationRemind"`
    ReservationSwitch bool `json:"reservationSwitch"`
}
type GroupDeliveryPoiDeliveryRuleCreateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 操作结果
    */
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}
type TieredDiscount struct {
    Discount string `json:"discount"`
    EndDistance string `json:"endDistance"`
    StartDistance string `json:"startDistance"`
    Tiered int32 `json:"tiered"`
}
type GroupDeliveryPoiDeliveryRuleCreateRequest struct {
    ContactPhone string `json:"contactPhone"`
    Deliverable int32 `json:"deliverable"`
    DeliveryFeeDiscount DeliveryFeeDiscount `json:"deliveryFeeDiscount"`
    DeliveryHours []DeliveryHours `json:"deliveryHours"`
    DeliveryReservation DeliveryReservation `json:"deliveryReservation"`
}



func (req *GroupDeliveryPoiDeliveryRuleCreateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupDeliveryPoiDeliveryRuleCreateResponse, error) {
    resp, err := client.InvokeApi(group_delivery_poi_delivery_rule_create_url, 46, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupDeliveryPoiDeliveryRuleCreateResponse
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

func (response *GroupDeliveryPoiDeliveryRuleCreateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

