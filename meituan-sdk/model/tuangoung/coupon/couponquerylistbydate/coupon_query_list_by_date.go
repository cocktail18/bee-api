package couponquerylistbydate

/**
* 门店验券历史
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const coupon_query_list_by_date_url = "/tuangou/coupon/queryListByDate"

type CouponQueryListByDateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data CouponQueryListByDateData `json:"data"`
    TraceId string `json:"traceId"`
}
type CouponQueryListByDateData struct {
    CouponCodes []CouponCodes `json:"couponCodes"`
    Total int64 `json:"total"`
}
type CouponCodes struct {
    /**
    * 团购券购买价格
    */
    CouponBuyPrice float64 `json:"couponBuyPrice"`
    /**
    * 团购券是否可撤销。1表示可撤销，0表示不可撤销
    */
    CouponCancelStatus int64 `json:"couponCancelStatus"`
    /**
    * 团购券码
    */
    CouponCode string `json:"couponCode"`
    /**
    * 团购券状态，包含：未使用、已使用、已冻结、作弊已冻结、已退款
    */
    CouponStatusDesc string `json:"couponStatusDesc"`
    /**
    * 验券时间
    */
    CouponUseTime string `json:"couponUseTime"`
    /**
    * 团购券对应项目开始售卖时间
    */
    DealBeginTime string `json:"dealBeginTime"`
    /**
    * 项目对应的标题
    */
    DealTitle string `json:"dealTitle"`
    /**
    * 项目售卖价
    */
    DealValue float64 `json:"dealValue"`
    IsVerifyPay int64 `json:"isVerifyPay"`
    IsVoucher bool `json:"isVoucher"`
    /**
    * 验券帐号
    */
    VerifyAcct string `json:"verifyAcct"`
    /**
    * 验券方式
    */
    VerifyType string `json:"verifyType"`
}
type CouponQueryListByDateRequest struct {
    /**
    *  日期 
    */
    Date string `json:"date"`
    /**
    *  查询起始位置，从0开始 
    */
    Offset int32 `json:"offset"`
    /**
    *  查询条数 
    */
    Limit int32 `json:"limit"`
}



func (req *CouponQueryListByDateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CouponQueryListByDateResponse, error) {
    resp, err := client.InvokeApi(coupon_query_list_by_date_url, 1, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CouponQueryListByDateResponse
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

func (response *CouponQueryListByDateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

