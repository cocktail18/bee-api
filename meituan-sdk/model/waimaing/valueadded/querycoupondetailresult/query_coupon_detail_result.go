package querycoupondetailresult

/**
* 根据发券任务ID查询发券活动效果明细
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_coupon_detail_result_url = "/waimai/ng/valueadded/queryCouponDetailResult"

type QueryCouponDetailResultRequest struct {
    /**
    *  发券返回的id。 支持手机号发券生成的任务id。 
    */
    QueryId string `json:"queryId"`
}
type QueryCouponDetailResultResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Data struct {
    Used []int64 `json:"used"`
    UnUsed []int64 `json:"unUsed"`
    Expire []int64 `json:"expire"`
}



func (req *QueryCouponDetailResultRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryCouponDetailResultResponse, error) {
    resp, err := client.InvokeApi(query_coupon_detail_result_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryCouponDetailResultResponse
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

func (response *QueryCouponDetailResultResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

