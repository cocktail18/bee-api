package adupdateplanbid

/**
* 修改广告投放计划出价
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const ad_update_plan_bid_url = "/wmoper/ng/waimaiad/plan/update/bid"

type AdUpdatePlanBidRequest struct {
    /**
    * 外卖广告的产品类型，1表示点金推广
    */
    Channel int32 `json:"channel"`
    /**
    * 出价，单位分
    */
    Bid int32 `json:"bid"`
}
type Data struct {
    Success bool `json:"success"`
    Code string `json:"code"`
}
type AdUpdatePlanBidResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *AdUpdatePlanBidRequest) DoInvoke(client mtclient.MeituanClient) (*AdUpdatePlanBidResponse, error) {
    resp, err := client.InvokeApi(ad_update_plan_bid_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response AdUpdatePlanBidResponse
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

func (response *AdUpdatePlanBidResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

