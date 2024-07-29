package groupdeliverypoideliveryrulerangeupdate

/**
* 门店配送规则更新配送范围
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const group_delivery_poi_delivery_rule_range_update_url = "/dcps/fulfill/poi/deliveryRule/range/update"

type GroupDeliveryPoiDeliveryRuleRangeUpdateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}
type GroupDeliveryPoiDeliveryRuleRangeUpdateRequest struct {
    NeedFix bool `json:"needFix"`
    ServiceDeliveryRangeDetails []ServiceDeliveryRangeDetails `json:"serviceDeliveryRangeDetails"`
}
type DeliveryRange struct {
    X string `json:"x"`
    Y string `json:"y"`
}
type ServiceDeliveryRangeDetails struct {
    DeliveryRanges []DeliveryRange `json:"deliveryRanges"`
    ServiceCode string `json:"serviceCode"`
}



func (req *GroupDeliveryPoiDeliveryRuleRangeUpdateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupDeliveryPoiDeliveryRuleRangeUpdateResponse, error) {
    resp, err := client.InvokeApi(group_delivery_poi_delivery_rule_range_update_url, 46, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupDeliveryPoiDeliveryRuleRangeUpdateResponse
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

func (response *GroupDeliveryPoiDeliveryRuleRangeUpdateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

