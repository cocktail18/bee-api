package actdiscountactivityorderlimit

/**
* 更新折扣商品外卖门店维度每单限购数量
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const act_discount_activity_order_limit_url = "/waimai/ng/act/discount/activityOrderLimit"

type ActDiscountActivityOrderLimitResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type ActDiscountActivityOrderLimitRequest struct {
    /**
    *  本门店每单可购买的折扣商品数量 
    */
    ActivityOrderLimit int32 `json:"activityOrderLimit"`
}



func (req *ActDiscountActivityOrderLimitRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ActDiscountActivityOrderLimitResponse, error) {
    resp, err := client.InvokeApi(act_discount_activity_order_limit_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ActDiscountActivityOrderLimitResponse
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

func (response *ActDiscountActivityOrderLimitResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

