package getcouponpriceinfo

/**
* 获取团购券交易快照
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const get_coupon_price_info_url = "/tuangou/ng/coupon/getCouponPriceInfo"

type GetCouponPriceInfoRequest struct {
    /**
    *  要验证的美团券密码、二维码 
    */
    Code string `json:"code"`
}
type GetCouponPriceResp struct {
    /**
    * 团购券售卖价。用户交易成功时，商家开店宝设置的团购促销前的购买价格，单位：分
    */
    SalePrice int64 `json:"salePrice"`
    /**
    * 团购券面值或门市价。用户交易成功时，商家开店宝设置的团购券面值，单位：分
    */
    OriginalPrice int64 `json:"originalPrice"`
    /**
    * 交易时间，即团购券交易成功时间，秒级
    */
    BuyTime int64 `json:"buyTime"`
}
type GetCouponPriceInfoResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data GetCouponPriceResp `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *GetCouponPriceInfoRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GetCouponPriceInfoResponse, error) {
    resp, err := client.InvokeApi(get_coupon_price_info_url, 1, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GetCouponPriceInfoResponse
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

func (response *GetCouponPriceInfoResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

