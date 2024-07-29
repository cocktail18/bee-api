package querycoupondetailextraresult

/**
* 精准营销-活动效果汇总查询接口（含券使用时间）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_coupon_detail_extra_result_url = "/waimai/ng/valueadded/queryCouponDetailExtraResult"

type QueryCouponDetailExtraResultResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data QueryCouponDetailExtraResultData `json:"data"`
    TraceId string `json:"traceId"`
}
type CouponUseInfoVo struct {
    /**
    *  用户id
    */
    UserId int64 `json:"userId"`
    /**
    *  订单号
    */
    OrderViewId int64 `json:"orderViewId"`
    /**
    * 使用时间
    */
    UsedTime int64 `json:"usedTime"`
}
type QueryCouponDetailExtraResultData struct {
    /**
    * 已使用券的用户列表
    */
    Used []int64 `json:"used"`
    /**
    * 未使用券的用户列表
    */
    UnUsed []int64 `json:"unUsed"`
    /**
    * 过期的券的用户列表
    */
    Expire []int64 `json:"expire"`
    UsedInfo CouponUseInfoVo `json:"usedInfo"`
}
type QueryCouponDetailExtraResultRequest struct {
    /**
    *  发券返回的id。 支持手机号发券生成的任务id。 
    */
    QueryId string `json:"queryId"`
}



func (req *QueryCouponDetailExtraResultRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryCouponDetailExtraResultResponse, error) {
    resp, err := client.InvokeApi(query_coupon_detail_extra_result_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryCouponDetailExtraResultResponse
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

func (response *QueryCouponDetailExtraResultResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

