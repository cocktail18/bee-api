package groupdeliverypoideliveryrulereservationupdate

/**
* 门店配送规则更新预约规则
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const group_delivery_poi_delivery_rule_reservation_update_url = "/dcps/fulfill/poi/deliveryRule/reservation/update"

type GroupDeliveryPoiDeliveryRuleReservationUpdateRequest struct {
    AvailableDays []int32 `json:"availableDays"`
    BeyondDeliveryTimeOrder bool `json:"beyondDeliveryTimeOrder"`
    ReservationRemind int32 `json:"reservationRemind"`
    ReservationSwitch bool `json:"reservationSwitch"`
}
type GroupDeliveryPoiDeliveryRuleReservationUpdateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *GroupDeliveryPoiDeliveryRuleReservationUpdateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupDeliveryPoiDeliveryRuleReservationUpdateResponse, error) {
    resp, err := client.InvokeApi(group_delivery_poi_delivery_rule_reservation_update_url, 46, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupDeliveryPoiDeliveryRuleReservationUpdateResponse
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

func (response *GroupDeliveryPoiDeliveryRuleReservationUpdateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

