package groupdeliverypoibizdeliveryrulefeeupdate

/**
* 门店配送规则更新自配配送费
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const group_delivery_poi_biz_delivery_rule_fee_update_url = "/dcps/fulfill/poi/biz/deliveryRule/fee/update"

type TimeSpan struct {
    BeginTime string `json:"beginTime"`
    EndTime string `json:"endTime"`
}
type TieredDeliveryFee struct {
    Tiered int32 `json:"tiered"`
    StartDistance string `json:"startDistance"`
    EndDistance string `json:"endDistance"`
    Fee string `json:"fee"`
}
type GroupDeliveryPoiBizDeliveryRuleFeeUpdateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}
type GroupDeliveryPoiBizDeliveryRuleFeeUpdateRequest struct {
    DeliveryFeeType int32 `json:"deliveryFeeType"`
    DeliveryFee string `json:"deliveryFee"`
    TieredDeliveryFee []TieredDeliveryFee `json:"tieredDeliveryFee"`
    AdditionalDeliveryFeeHours []AdditionalDeliveryFeeHours `json:"additionalDeliveryFeeHours"`
}
type AdditionalDeliveryFeeHours struct {
    Hour TimeSpan `json:"hour"`
    AdditionalDeliveryFee string `json:"additionalDeliveryFee"`
}



func (req *GroupDeliveryPoiBizDeliveryRuleFeeUpdateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupDeliveryPoiBizDeliveryRuleFeeUpdateResponse, error) {
    resp, err := client.InvokeApi(group_delivery_poi_biz_delivery_rule_fee_update_url, 46, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupDeliveryPoiBizDeliveryRuleFeeUpdateResponse
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

func (response *GroupDeliveryPoiBizDeliveryRuleFeeUpdateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

