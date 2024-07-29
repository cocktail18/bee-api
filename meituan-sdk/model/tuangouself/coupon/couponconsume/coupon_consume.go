package couponconsume

/**
* 执行验券
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const coupon_consume_url = "/tuangouself/coupon/consume"

type CouponConsumeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data CouponConsumeData `json:"data"`
    TraceId string `json:"traceId"`
}
type CouponConsumeData struct {
    /**
    * 验券结果
    */
    Result bool `json:"result"`
}
type CouponConsumeRequest struct {
    /**
    *  券码 
    */
    Codes []string `json:"codes"`
    /**
    *  幂等因子（会与授权的门店一起构造唯一请求标识校验幂等，缓存时间为5分钟） 
    */
    Idempotent string `json:"idempotent"`
}



func (req *CouponConsumeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CouponConsumeResponse, error) {
    resp, err := client.InvokeApi(coupon_consume_url, 33, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CouponConsumeResponse
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

func (response *CouponConsumeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

