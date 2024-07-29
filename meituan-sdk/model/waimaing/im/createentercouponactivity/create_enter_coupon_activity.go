package createentercouponactivity

/**
* 创建进群领券活动
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const create_enter_coupon_activity_url = "/waimai/ng/im/createEnterCouponActivity"

type CreateEnterCouponActivityRequest struct {
    /**
    *  群ID 
    */
    GroupId int64 `json:"groupId"`
    /**
    *  优惠券金额 单位元 
    */
    Price int32 `json:"price"`
    /**
    *  优惠券门槛 单位元 
    */
    LimitPrice int32 `json:"limitPrice"`
    /**
    *  券有效期 
    */
    Validity int32 `json:"validity"`
}
type CreateEnterCouponActivityResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}



func (req *CreateEnterCouponActivityRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CreateEnterCouponActivityResponse, error) {
    resp, err := client.InvokeApi(create_enter_coupon_activity_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CreateEnterCouponActivityResponse
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

func (response *CreateEnterCouponActivityResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

