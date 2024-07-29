package groupdeliverypoideliveryrulebaseupdate

/**
* 2.1.3.更新门店配送规则基础信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const group_delivery_poi_delivery_rule_base_update_url = "/dcps/fulfill/poi/deliveryRule/base/update"

type GroupDeliveryPoiDeliveryRuleBaseUpdateRequest struct {
    /**
    *  商家联系电话 
    */
    ContactPhone string `json:"contactPhone"`
    /**
    *  是否支持配送 1:支持， 0:不支持 
    */
    Deliverable int32 `json:"deliverable"`
    /**
    *  配送时间，结构同查询接口 
    */
    DeliveryHours []DeliveryHours `json:"deliveryHours"`
    OrderPhone []PhoneDTO `json:"orderPhone"`
    OuterPoiId string `json:"outerPoiId"`
}
type GroupDeliveryPoiDeliveryRuleBaseUpdateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}
type DeliveryHours struct {
    BeginTime string `json:"beginTime"`
    EndTime string `json:"endTime"`
}
type PhoneDTO struct {
    PhoneNumber string `json:"phoneNumber"`
    PhoneType int32 `json:"phoneType"`
}



func (req *GroupDeliveryPoiDeliveryRuleBaseUpdateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupDeliveryPoiDeliveryRuleBaseUpdateResponse, error) {
    resp, err := client.InvokeApi(group_delivery_poi_delivery_rule_base_update_url, 46, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupDeliveryPoiDeliveryRuleBaseUpdateResponse
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

func (response *GroupDeliveryPoiDeliveryRuleBaseUpdateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

