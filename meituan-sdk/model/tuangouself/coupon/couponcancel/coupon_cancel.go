package couponcancel

/**
* 撤销验券
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const coupon_cancel_url = "/tuangouself/coupon/cancel"

type CouponCancelResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data CouponCancelData `json:"data"`
    TraceId string `json:"traceId"`
}
type CouponCancelData struct {
    /**
    * 撤销验券结果
    */
    Result bool `json:"result"`
}
type CouponCancelRequest struct {
    /**
    *  券码 
    */
    Code string `json:"code"`
}



func (req *CouponCancelRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CouponCancelResponse, error) {
    resp, err := client.InvokeApi(coupon_cancel_url, 33, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CouponCancelResponse
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

func (response *CouponCancelResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

