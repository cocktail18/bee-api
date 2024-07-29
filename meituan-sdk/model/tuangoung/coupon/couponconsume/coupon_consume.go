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

const coupon_consume_url = "/tuangou/coupon/consume"

type CouponConsumeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data CouponConsumeInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type CouponConsumeInfo struct {
    /**
    * 验证券码数组
    */
    CouponCodes []string `json:"couponCodes"`
    /**
    * 操作状态
    */
    Result int32 `json:"result"`
    /**
    * 项目ID
    */
    DealId int64 `json:"dealId"`
    /**
    * 项目名称
    */
    DealTitle string `json:"dealTitle"`
    /**
    * 美团门店ID
    */
    Poiid int64 `json:"poiid"`
    /**
    * 返回值消息
    */
    Message string `json:"message"`
    /**
    * 开店宝套餐名
    */
    RawTitle string `json:"rawTitle"`
    /**
    * 团购项目在App端的展示标题
    */
    PlatformTitle string `json:"platformTitle"`
    /**
    * 券面值，即人们常说的市场价
    */
    DealValue float64 `json:"dealValue"`
}
type CouponConsumeRequest struct {
    /**
    *  商家登录ERP帐号ID 
    */
    EId string `json:"eId"`
    /**
    *  商家登录ERP帐号名称 
    */
    EName string `json:"eName"`
    /**
    *  第三方ERP订单号 
    */
    EOrderId string `json:"eOrderId"`
    /**
    *  数量 
    */
    Count int32 `json:"count"`
    /**
    *  券码 
    */
    CouponCode string `json:"couponCode"`
}



func (req *CouponConsumeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CouponConsumeResponse, error) {
    resp, err := client.InvokeApi(coupon_consume_url, 1, appAuthToken, req)

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

