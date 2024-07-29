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

const coupon_cancel_url = "/tuangou/coupon/cancel"

type CouponCancelResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data CouponCancelData `json:"data"`
    TraceId string `json:"traceId"`
}
type CouponCancelData struct {
    /**
    * 操作状态0表示成功，其余表示失败
    */
    Result int32 `json:"result"`
    /**
    * 撤销结果描述信息
    */
    Message string `json:"message"`
}
type CouponCancelRequest struct {
    /**
    *  商家登录ERP帐号ID，长度不超过32。 
    */
    EId string `json:"eId"`
    /**
    *  商家登录ERP帐号名称，长度不超过32位 
    */
    EName string `json:"eName"`
    /**
    *  type值永远为1撤销验券 
    */
    Type int32 `json:"type"`
    /**
    *  券码 
    */
    CouponCode string `json:"couponCode"`
}



func (req *CouponCancelRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CouponCancelResponse, error) {
    resp, err := client.InvokeApi(coupon_cancel_url, 1, appAuthToken, req)

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

