package groupdeliverypoideliveryrulequery

/**
* 门店配送规则查询
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const group_delivery_poi_delivery_rule_query_url = "/dcps/fulfill/poi/deliveryRule/query"

type DeliveryFeeDiscount struct {
    DiscountSwitch bool `json:"discountSwitch"`
    DiscountType int32 `json:"discountType"`
    UnifiedDiscount string `json:"unifiedDiscount"`
    TieredDiscounts []TieredDiscount `json:"tieredDiscounts"`
}
type GroupDeliveryPoiDeliveryRuleQueryRequest struct {
}
type DeliveryHours struct {
    BeginTime string `json:"beginTime"`
    EndTime string `json:"endTime"`
}
type DeliveryReservation struct {
    ReservationSwitch bool `json:"reservationSwitch"`
    ReservationRemind int32 `json:"reservationRemind"`
    BeyondDeliveryTimeOrder bool `json:"beyondDeliveryTimeOrder"`
    AvailableDays []int32 `json:"availableDays"`
}
type DeliveryRanges struct {
    X string `json:"x"`
    Y string `json:"y"`
}
type Data struct {
    Deliverable int32 `json:"deliverable"`
    Status int32 `json:"status"`
    ServiceDeliveryRangeDetails []ServiceDeliveryRangeDetails `json:"serviceDeliveryRangeDetails"`
    DeliveryHours []DeliveryHours `json:"deliveryHours"`
    DeliveryFeeDiscount DeliveryFeeDiscount `json:"deliveryFeeDiscount"`
    DeliveryReservation DeliveryReservation `json:"deliveryReservation"`
}
type TieredDiscount struct {
    Tiered int32 `json:"tiered"`
    StartDistance string `json:"startDistance"`
    EndDistance string `json:"endDistance"`
    Discount string `json:"discount"`
}
type ServiceDeliveryRangeDetails struct {
    ServiceCode string `json:"serviceCode"`
    DeliveryRanges []DeliveryRanges `json:"deliveryRanges"`
}
type GroupDeliveryPoiDeliveryRuleQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *GroupDeliveryPoiDeliveryRuleQueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupDeliveryPoiDeliveryRuleQueryResponse, error) {
    resp, err := client.InvokeApi(group_delivery_poi_delivery_rule_query_url, 46, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupDeliveryPoiDeliveryRuleQueryResponse
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

func (response *GroupDeliveryPoiDeliveryRuleQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

