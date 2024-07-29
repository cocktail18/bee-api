package remaincoupon

/**
* 查询门店剩余发券数
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const remain_coupon_url = "/waimai/ng/im/remainCoupon"

type RemainCouponResult struct {
    /**
    * 门店当月的剩余发券数量
    */
    RemainCouponNum int32 `json:"remainCouponNum"`
}
type RemainCouponResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data RemainCouponResult `json:"data"`
    TraceId string `json:"traceId"`
}
type RemainCouponRequest struct {
    /**
    *  粉丝群Id 
    */
    GroupId int64 `json:"groupId"`
}



func (req *RemainCouponRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*RemainCouponResponse, error) {
    resp, err := client.InvokeApi(remain_coupon_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response RemainCouponResponse
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

func (response *RemainCouponResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

