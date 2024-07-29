package couponquery

/**
* 查询用户授权门店可核销券
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const coupon_query_url = "/tuangouself/coupon/query"

type UserRule struct {
    /**
    * 是否可以和其他优惠同享
    */
    CanShare CanShare `json:"canShare"`
    /**
    * 美团券使用限制
    */
    UseLimit UseLimit `json:"useLimit"`
    /**
    * 代金券可用张数，非代金券:null，不可叠加:1，可叠加不限制:-1，可叠加限制:n
    */
    VoucherOverlayNum int32 `json:"voucherOverlayNum"`
}
type CouponQueryRequest struct {
}
type Meal struct {
    /**
    * 菜品名称
    */
    MealName string `json:"mealName"`
    /**
    * 菜品数量
    */
    MealNum string `json:"mealNum"`
}
type CouponQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []CouponResponse `json:"data"`
    TraceId string `json:"traceId"`
}
type DealMenu struct {
    /**
    * 项目ID
    */
    DealId int64 `json:"dealId"`
    /**
    * 菜品分类
    */
    MenuTitle string `json:"menuTitle"`
    /**
    * 菜品列表
    */
    Meals []Meal `json:"meals"`
}
type CouponResponse struct {
    /**
    * 项目ID
    */
    DealId int32 `json:"dealId"`
    /**
    * 项目名称
    */
    DealTitle string `json:"dealTitle"`
    /**
    * 项目类型，0套餐，1代金券
    */
    DealType int32 `json:"dealType"`
    /**
    * 券码
    */
    Code string `json:"code"`
    /**
    * 券有效期开始时间，单位秒
    */
    CouponStartTime int64 `json:"couponStartTime"`
    /**
    * 券有效期结束时间，单位秒
    */
    CouponEndTime int64 `json:"couponEndTime"`
    /**
    * 券对应的订单id
    */
    OrderId int64 `json:"orderId"`
    /**
    * 项目的市场价，用于展示
    */
    BuyPrice string `json:"buyPrice"`
    /**
    * 下单时团购券售卖价格
    */
    SalePrice string `json:"salePrice"`
    /**
    * 用户购买价
    */
    SettlePrice string `json:"settlePrice"`
    /**
    * 券码购买平台（0美团、1点评）
    */
    Platform int32 `json:"platform"`
    /**
    * 券使用规则
    */
    UseRule UserRule `json:"useRule"`
    /**
    * deal对应的菜品信息
    */
    MenuInfo []DealMenu `json:"menuInfo"`
    /**
    * 是否可用
    */
    CanUse bool `json:"canUse"`
    /**
    * 非使用时间是否可以验券。true:是；false:否
    */
    CanUseAnyTime bool `json:"canUseAnyTime"`
    /**
    * 不可用原因
    */
    UnavailableReason string `json:"unavailableReason"`
    /**
    * 开店宝上单商家录入的项目名称，仅套餐券使用
    */
    RawTitle string `json:"rawTitle"`
}
type CanShare struct {
    /**
    * 是否可同享优惠
    */
    CanShare bool `json:"canShare"`
    /**
    * 能享受的优惠限制
    */
    LimitPromotion string `json:"limitPromotion"`
}
type UseLimit struct {
    /**
    * 是否限制使用美团券，若为代金券，返回默认值false
    */
    IsLimit bool `json:"isLimit"`
    /**
    * 每人限制张数
    */
    PersonLimit int32 `json:"personLimit"`
    /**
    * 每桌限制张数
    */
    TableLimit int32 `json:"tableLimit"`
}



func (req *CouponQueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CouponQueryResponse, error) {
    resp, err := client.InvokeApi(coupon_query_url, 33, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CouponQueryResponse
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

func (response *CouponQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

