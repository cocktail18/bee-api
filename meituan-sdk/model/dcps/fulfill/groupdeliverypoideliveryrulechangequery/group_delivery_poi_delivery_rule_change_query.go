package groupdeliverypoideliveryrulechangequery

/**
* 门店配送规则变更详情查询
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const group_delivery_poi_delivery_rule_change_query_url = "/dcps/fulfill/poi/deliveryRule/change/query"

type GroupDeliveryPoiDeliveryRuleChangeQueryRequest struct {
}
type Data struct {
    Status int32 `json:"status"`
    ChangeContents []ChangeContents `json:"changeContents"`
    CreateTime string `json:"createTime"`
    UpdateTime string `json:"updateTime"`
    RejectMessage string `json:"rejectMessage"`
}
type ChangeContents struct {
    FieldName string `json:"fieldName"`
    NewValue string `json:"newValue"`
    OldValue string `json:"oldValue"`
}
type GroupDeliveryPoiDeliveryRuleChangeQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *GroupDeliveryPoiDeliveryRuleChangeQueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupDeliveryPoiDeliveryRuleChangeQueryResponse, error) {
    resp, err := client.InvokeApi(group_delivery_poi_delivery_rule_change_query_url, 46, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupDeliveryPoiDeliveryRuleChangeQueryResponse
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

func (response *GroupDeliveryPoiDeliveryRuleChangeQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

