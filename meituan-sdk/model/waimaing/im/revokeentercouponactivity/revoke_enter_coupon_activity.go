package revokeentercouponactivity

/**
* 停止当前的进群领券活动
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const revoke_enter_coupon_activity_url = "/waimai/ng/im/revokeEnterCouponActivity"

type RevokeEnterCouponActivityRequest struct {
    /**
    *  群ID 
    */
    GroupId int64 `json:"groupId"`
    /**
    *  活动ID 
    */
    ActivityId int32 `json:"activityId"`
}
type RevokeEnterCouponActivityResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}



func (req *RevokeEnterCouponActivityRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*RevokeEnterCouponActivityResponse, error) {
    resp, err := client.InvokeApi(revoke_enter_coupon_activity_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response RevokeEnterCouponActivityResponse
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

func (response *RevokeEnterCouponActivityResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

