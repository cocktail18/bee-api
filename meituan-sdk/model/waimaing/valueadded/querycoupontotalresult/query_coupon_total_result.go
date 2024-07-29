package querycoupontotalresult

/**
* 活动效果汇总查询接口
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_coupon_total_result_url = "/waimai/ng/valueadded/queryCouponTotalResult"

type QueryCouponTotalResultRequest struct {
    /**
    *  发券返回的id。 支持手机号发券生成的任务id 
    */
    QueryId string `json:"queryId"`
}
type Data struct {
    SendNum int64 `json:"sendNum"`
    UseNum int64 `json:"useNum"`
    UnUseNum int64 `json:"unUseNum"`
    ExpireNum int64 `json:"expireNum"`
}
type QueryCouponTotalResultResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *QueryCouponTotalResultRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryCouponTotalResultResponse, error) {
    resp, err := client.InvokeApi(query_coupon_total_result_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryCouponTotalResultResponse
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

func (response *QueryCouponTotalResultResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

