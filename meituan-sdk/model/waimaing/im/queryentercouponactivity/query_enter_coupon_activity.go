package queryentercouponactivity

/**
* 查询当前的进群领券活动
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_enter_coupon_activity_url = "/waimai/ng/im/queryEnterCouponActivity"

type QueryEnterCouponActivity struct {
    /**
    * 查询游标
    */
    Index int64 `json:"index"`
    /**
    * 是否还有记录
    */
    HasMore bool `json:"hasMore"`
    /**
    * 活动ID
    */
    ActivityId int64 `json:"activityId"`
    /**
    * 优惠券金额 单位元
    */
    Price int32 `json:"price"`
    /**
    * 优惠券门槛 单位元
    */
    LimitPrice int32 `json:"limitPrice"`
    /**
    * 有效性：-1无效，1有效
    */
    Valid int32 `json:"valid"`
    /**
    * 券有效期 7代表7天
    */
    Validity int32 `json:"validity"`
    /**
    * 券创建时间（秒）
    */
    CTime int32 `json:"cTime"`
    /**
    * 券结束时间，未结束为0（秒）
    */
    EndTime int32 `json:"endTime"`
}
type QueryEnterCouponActivityResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data QueryEnterCouponActivity `json:"data"`
    TraceId string `json:"traceId"`
}
type QueryEnterCouponActivityRequest struct {
    /**
    *  查询游标，首次传0 
    */
    Index int64 `json:"index"`
    /**
    *  查询个数，大于1 
    */
    Size int32 `json:"size"`
    /**
    *  优化券状态 （有效 1 ，无效 -1） 
    */
    Status string `json:"status"`
}



func (req *QueryEnterCouponActivityRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryEnterCouponActivityResponse, error) {
    resp, err := client.InvokeApi(query_enter_coupon_activity_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryEnterCouponActivityResponse
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

func (response *QueryEnterCouponActivityResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

