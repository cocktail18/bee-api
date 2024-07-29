package couponprepare

/**
* 验券准备
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const coupon_prepare_url = "/tuangou/coupon/prepare"

type CouponPrepareResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data CouponPrepareData `json:"data"`
    TraceId string `json:"traceId"`
}
type CouponPrepareData struct {
    /**
    * 项目开始时间
    */
    DealBeginTime string `json:"dealBeginTime"`
    /**
    * 项目ID
    */
    DealId int32 `json:"dealId"`
    /**
    * 券面值，即人们常说的市场价
    */
    DealValue float64 `json:"dealValue"`
    /**
    * 项目名称
    */
    DealTitle string `json:"dealTitle"`
    /**
    * 最大可验证张数
    */
    Count int32 `json:"count"`
    /**
    * 量贩项目的单张券原价（普通券单张券原价与项目总价相同）
    */
    SingleValue float64 `json:"singleValue"`
    /**
    * 团购券价格，即商家促销前的券购买价格，如首次购买有更多优惠这类需要在开店宝设置的促销
    */
    DealPrice float64 `json:"dealPrice"`
    /**
    * 套餐类券码对应套餐详细内容，代金券券码包含适用范围(如酒水除外)
    */
    DealMenu [][]DealMenu `json:"dealMenu"`
    /**
    * 返回消息
    */
    Message string `json:"message"`
    /**
    * 最小消费张数
    */
    MinConsume int32 `json:"minConsume"`
    /**
    * 券码有效期
    */
    CouponEndTime string `json:"couponEndTime"`
    /**
    * 是否量贩：0：不是，1：是
    */
    Volume int32 `json:"volume"`
    /**
    * 操作状态,0表示成功
    */
    Result int32 `json:"result"`
    /**
    * 券购买价,即用户在购买团购券时所付的实际价格。返回-1则表示没查到couponBuyPrice信息。
    */
    CouponBuyPrice float64 `json:"couponBuyPrice"`
    /**
    * 券码
    */
    CouponCode string `json:"couponCode"`
    /**
    * 是否代金券,true代表代金券,false代表套餐券
    */
    IsVoucher bool `json:"isVoucher"`
    /**
    * 开店宝套餐名
    */
    RawTitle string `json:"rawTitle"`
    /**
    * 团购项目在App端的展示标题
    */
    PlatformTitle string `json:"platformTitle"`
}
type CouponPrepareRequest struct {
    /**
    *  券码 
    */
    CouponCode string `json:"couponCode"`
}
type DealMenu struct {
    /**
    * 字段内容
    */
    Content string `json:"content"`
    /**
    * 数量/规格
    */
    Specification string `json:"specification"`
    /**
    * 单价
    */
    Price string `json:"price"`
    /**
    * 小计
    */
    Total string `json:"total"`
    /**
    * 区别表头和行：0表示表头，128表示行
    */
    Type string `json:"type"`
}



func (req *CouponPrepareRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CouponPrepareResponse, error) {
    resp, err := client.InvokeApi(coupon_prepare_url, 1, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CouponPrepareResponse
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

func (response *CouponPrepareResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

