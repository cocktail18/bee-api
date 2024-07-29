package querycoupondetail

/**
* 查询发券的活动效果
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_coupon_detail_url = "/waimai/ng/im/queryReceiveCouponDetail"

type QueryCouponDetailRequest struct {
    /**
    *  配置ID 
    */
    CouponConfigId string `json:"couponConfigId"`
}
type QueryCouponDetailResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data QueryReceiveCouponDetail `json:"data"`
    TraceId string `json:"traceId"`
}
type QueryReceiveCouponDetail struct {
    /**
    * 发送成功数量
    */
    SendNum int32 `json:"sendNum"`
    /**
    * 已使用
    */
    UseNum int32 `json:"useNum"`
    /**
    * 未使用的
    */
    UnUseNum int32 `json:"unUseNum"`
    /**
    * 已过期的数量
    */
    ExpireNum int32 `json:"expireNum"`
    /**
    * 已使用用户列表
    */
    Used []int64 `json:"used"`
    /**
    * 未使用用户列表
    */
    UnUsed []int64 `json:"unUsed"`
    /**
    * 已过期用户列表
    */
    Expire []int64 `json:"expire"`
}



func (req *QueryCouponDetailRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryCouponDetailResponse, error) {
    resp, err := client.InvokeApi(query_coupon_detail_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryCouponDetailResponse
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

func (response *QueryCouponDetailResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

