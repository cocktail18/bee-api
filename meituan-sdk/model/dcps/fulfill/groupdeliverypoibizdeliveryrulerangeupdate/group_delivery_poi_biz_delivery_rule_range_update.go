package groupdeliverypoibizdeliveryrulerangeupdate

/**
* 门店配送规则更新自配配送范围
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const group_delivery_poi_biz_delivery_rule_range_update_url = "/dcps/fulfill/poi/biz/deliveryRule/range/update"

type GroupDeliveryPoiBizDeliveryRuleRangeUpdateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}
type PeakHoursDeliveryRange struct {
    X string `json:"x"`
    Y string `json:"y"`
}
type PeakHours struct {
    BeginTime string `json:"beginTime"`
    EndTime string `json:"endTime"`
}
type RegularHourDeliveryRange struct {
    X string `json:"x"`
    Y string `json:"y"`
}
type GroupDeliveryPoiBizDeliveryRuleRangeUpdateRequest struct {
    NeedFix bool `json:"needFix"`
    RegularHourDeliveryRange []RegularHourDeliveryRange `json:"regularHourDeliveryRange"`
    PeakHours []PeakHours `json:"peakHours"`
    PeakHoursDeliveryRange []PeakHoursDeliveryRange `json:"peakHoursDeliveryRange"`
}



func (req *GroupDeliveryPoiBizDeliveryRuleRangeUpdateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupDeliveryPoiBizDeliveryRuleRangeUpdateResponse, error) {
    resp, err := client.InvokeApi(group_delivery_poi_biz_delivery_rule_range_update_url, 46, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupDeliveryPoiBizDeliveryRuleRangeUpdateResponse
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

func (response *GroupDeliveryPoiBizDeliveryRuleRangeUpdateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

