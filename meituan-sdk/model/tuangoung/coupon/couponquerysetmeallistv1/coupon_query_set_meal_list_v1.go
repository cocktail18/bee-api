package couponquerysetmeallistv1

/**
* 门店套餐映射；相比于coupon_query_set_meal_list拆分了隐藏状态和售卖状态
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const coupon_query_set_meal_list_v1_url = "/tuangou/ng/coupon/querySetMealListV1"

type DealInfosSub struct {
    /**
    * 项目结束时间戳，单位是秒
    */
    DealEndTime int64 `json:"dealEndTime"`
    /**
    * 项目ID
    */
    DealId int64 `json:"dealId"`
    /**
    * 市场价格
    */
    DealValue float64 `json:"dealValue"`
    /**
    * 项目名称
    */
    DealTitle string `json:"dealTitle"`
    /**
    * 项目开始时间戳，单位是秒
    */
    BeginTime int64 `json:"beginTime"`
    /**
    * 售卖价格
    */
    DealPrice float64 `json:"dealPrice"`
    /**
    * json字符串数组
    */
    DealMenu string `json:"dealMenu"`
    /**
    * 是否是代金券，true为代金券，false为套餐
    */
    IsVoucher bool `json:"isVoucher"`
    /**
    * C端项目名称
    */
    PlatformDealTitle string `json:"platformDealTitle"`
    /**
    * 是否隐藏
    */
    IsHide bool `json:"isHide"`
    /**
    * 售卖状态
    */
    SaleStatus int32 `json:"saleStatus"`
    /**
    * 券有效期模式 0表示按券结束日期，1表示按购买之后couponValidDays天有效
    */
    CouponDateType int32 `json:"couponDateType"`
    /**
    * 券有效期固定结束时间 单位分
    */
    CouponEndTime int64 `json:"couponEndTime"`
    /**
    * 券有效期多少天有效
    */
    CouponValidDays int32 `json:"couponValidDays"`
    /**
    * C端场景中提单页、支付结果页、查看券码和订单页面的项目标题以及开店宝核销后展示的标题。C端场景包括美团App，点评App，美团小程序
    */
    ShortAttrTitle string `json:"shortAttrTitle"`
    /**
    * 开店宝商品列表标题
    */
    PlatformTitle string `json:"platformTitle"`
    /**
    * 第三方商品ID(三方上单填入)
    */
    ThirdProductId string `json:"thirdProductId"`
}
type CouponQuerySetMealListV1Request struct {
    /**
    *  售卖状态；不传返回全部 状态值 状态描述 0 结束售卖的deal 1 正在售卖的deal 4 未开始售卖的deal 
    */
    SaleStatus int32 `json:"saleStatus"`
    /**
    *  隐藏状态 ；不传默认查询非隐藏单(渠道单) 状态值 状态描述 -1 所有deal 0 非隐藏单 1 隐藏单 
    */
    HideStatus int32 `json:"hideStatus"`
    /**
    *  分页偏移量，默认0 
    */
    Offset int64 `json:"offset"`
    /**
    *  分页大小，默认一个门店只返回20，最多可配置100 
    */
    Limit int64 `json:"limit"`
}
type Data struct {
    /**
    * 是否还有下一页
    */
    HasMore bool `json:"hasMore"`
    DealInfos []DealInfosSub `json:"dealInfos"`
}
type CouponQuerySetMealListV1Response struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *CouponQuerySetMealListV1Request) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CouponQuerySetMealListV1Response, error) {
    resp, err := client.InvokeApi(coupon_query_set_meal_list_v1_url, 1, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CouponQuerySetMealListV1Response
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

func (response *CouponQuerySetMealListV1Response) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

