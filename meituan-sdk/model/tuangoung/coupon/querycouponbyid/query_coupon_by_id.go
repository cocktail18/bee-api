package querycouponbyid

/**
* 已验券码查询
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_coupon_by_id_url = "/tuangou/coupon/queryById"

type QueryCouponByIdRequest struct {
    /**
    *  券码 
    */
    CouponCode string `json:"couponCode"`
}
type QueryCouponByIdResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data QueryCouponByIdData `json:"data"`
    TraceId string `json:"traceId"`
}
type QueryCouponByIdData struct {
    /**
    * 团购券码购买价格
    */
    CouponBuyPrice float64 `json:"couponBuyPrice"`
    /**
    * 券码是否可撤销，表示可撤销，0表示不可撤销
    */
    CouponCancelStatus int32 `json:"couponCancelStatus"`
    /**
    * 券码
    */
    CouponCode string `json:"couponCode"`
    /**
    * 券状态
    */
    CouponStatusDesc string `json:"couponStatusDesc"`
    /**
    * 券码使用时间
    */
    CouponUseTime string `json:"couponUseTime"`
    /**
    * 项目开始时间
    */
    DealBeginTime string `json:"dealBeginTime"`
    /**
    * 项目id
    */
    DealId int32 `json:"dealId"`
    /**
    * 项目名称
    */
    DealTitle string `json:"dealTitle"`
    /**
    * 市场价
    */
    DealValue float64 `json:"dealValue"`
    /**
    * 是否代金券，代表代金券,false代表套餐券
    */
    IsVoucher bool `json:"isVoucher"`
    /**
    * 量贩项目的单张券原价（普通券单张券原价与项目总价相同）
    */
    SingleValue float64 `json:"singleValue"`
    /**
    * 验券帐号
    */
    VerifyAcct string `json:"verifyAcct"`
    /**
    * 验券方式
    */
    VerifyType string `json:"verifyType"`
    /**
    * 是否量贩：0：不是，1：是
    */
    Volume int32 `json:"volume"`
}



func (req *QueryCouponByIdRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryCouponByIdResponse, error) {
    resp, err := client.InvokeApi(query_coupon_by_id_url, 1, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryCouponByIdResponse
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

func (response *QueryCouponByIdResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

