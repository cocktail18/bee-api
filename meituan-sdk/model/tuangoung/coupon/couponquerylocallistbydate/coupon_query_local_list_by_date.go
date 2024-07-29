package couponquerylocallistbydate

/**
* 门店本地验券历史
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const coupon_query_local_list_by_date_url = "/tuangou/coupon/queryLocalListByDate"

type CouponQueryLocalListByDateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data CouponQueryLocalListByDateData `json:"data"`
    TraceId string `json:"traceId"`
}
type CouponQueryLocalListByDateData struct {
    /**
    * ERP订单数组
    */
    EOrders []EOrders `json:"eOrders"`
    /**
    * ERP订单条数总数
    */
    Total int32 `json:"total"`
}
type EOrders struct {
    /**
    * couponCodes中券码个数
    */
    Count string `json:"count"`
    /**
    * 券码数组
    */
    CouponCodes []int64 `json:"couponCodes"`
    /**
    * 验券时间
    */
    CouponUseTime string `json:"couponUseTime"`
    /**
    * 项目名称
    */
    DealTitle string `json:"dealTitle"`
    /**
    * 第三方ERP订单号
    */
    EOrderId string `json:"eOrderId"`
    IsVerifyPay int64 `json:"isVerifyPay"`
}
type CouponQueryLocalListByDateRequest struct {
    /**
    *  日期 
    */
    Date string `json:"date"`
    /**
    *  查询起始位置，从0开始，最大不超过999 
    */
    Offset int32 `json:"offset"`
    /**
    *  查询条数，最大不超过999 
    */
    Limit int32 `json:"limit"`
}



func (req *CouponQueryLocalListByDateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CouponQueryLocalListByDateResponse, error) {
    resp, err := client.InvokeApi(coupon_query_local_list_by_date_url, 1, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CouponQueryLocalListByDateResponse
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

func (response *CouponQueryLocalListByDateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

