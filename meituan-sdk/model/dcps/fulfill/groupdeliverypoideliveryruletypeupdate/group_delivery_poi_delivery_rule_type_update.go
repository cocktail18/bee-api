package groupdeliverypoideliveryruletypeupdate

/**
* 门店配送规则更新切换配送服务
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const group_delivery_poi_delivery_rule_type_update_url = "/dcps/fulfill/poi/deliveryRule/type/update"

type GroupDeliveryPoiDeliveryRuleTypeUpdateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}
type GroupDeliveryPoiDeliveryRuleTypeUpdateRequest struct {
    NewDeliveryType int32 `json:"newDeliveryType"`
}



func (req *GroupDeliveryPoiDeliveryRuleTypeUpdateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupDeliveryPoiDeliveryRuleTypeUpdateResponse, error) {
    resp, err := client.InvokeApi(group_delivery_poi_delivery_rule_type_update_url, 46, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupDeliveryPoiDeliveryRuleTypeUpdateResponse
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

func (response *GroupDeliveryPoiDeliveryRuleTypeUpdateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

