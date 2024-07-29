package supplyrulesavebookingtimerule

/**
* 设置门店预订时间规则
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const supply_rule_save_booking_time_rule_url = "/resv2/rule/supply/saveBookingTimeRule"

type BookingHoldRuleDTO struct {
    /**
    *  是否限制桌位保留时间 
    */
    HasLimit bool `json:"hasLimit"`
    /**
    *  最小可订间隔,范围为[1,12]，单位十分钟 
    */
    BookingHoldMinutes int32 `json:"bookingHoldMinutes"`
}
type BookingHoldAndTimingIntervalRuleDTO struct {
    /**
    *  桌位保留时间 
    */
    BookingHoldRuleDTO BookingHoldRuleDTO `json:"bookingHoldRuleDTO"`
    /**
    *  最小可订间隔,15-间隔为15分钟 30-间隔为30分钟 
    */
    MinIntervalTime int32 `json:"minIntervalTime"`
}
type SupplyRuleSaveBookingTimeRuleResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}
type SupplyRuleSaveBookingTimeRuleRequest struct {
    BookingHoldRuleDTO BookingHoldAndTimingIntervalRuleDTO `json:"bookingHoldRuleDTO"`
}



func (req *SupplyRuleSaveBookingTimeRuleRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*SupplyRuleSaveBookingTimeRuleResponse, error) {
    resp, err := client.InvokeApi(supply_rule_save_booking_time_rule_url, 7, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response SupplyRuleSaveBookingTimeRuleResponse
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

func (response *SupplyRuleSaveBookingTimeRuleResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

