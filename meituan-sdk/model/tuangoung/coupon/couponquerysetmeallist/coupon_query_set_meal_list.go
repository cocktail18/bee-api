package couponquerysetmeallist

/**
* 门店套餐映射
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const coupon_query_set_meal_list_url = "/tuangou/coupon/querySetMealList"

type CouponQuerySetMealListRequest struct {
    /**
    *  套餐状态，可选的值为： 状态值 套餐状态dealStatus含义 -1 全部状态（含以下所有） 0 结束售卖 1 正在售卖 2 隐藏单，前台通过POI列表及deal列表都不可见，但是通过收藏或url可以直接访问并购买 4 未开始售卖 5 不可购买 
    */
    DealStatus int32 `json:"dealStatus"`
}
type DealInfo struct {
    /**
    * 项目结束时间戳，单位是秒
    */
    DealEndTime int64 `json:"dealEndTime"`
    /**
    * 项目ID
    */
    DealId int32 `json:"dealId"`
    /**
    * 市场价格
    */
    DealValue float64 `json:"dealValue"`
    /**
    * 项目名称
    */
    DealTitle string `json:"dealTitle"`
    /**
    * 请求的套餐状态
    */
    DealStatus int32 `json:"dealStatus"`
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
}
type CouponQuerySetMealListResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []DealInfo `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *CouponQuerySetMealListRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CouponQuerySetMealListResponse, error) {
    resp, err := client.InvokeApi(coupon_query_set_meal_list_url, 1, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CouponQuerySetMealListResponse
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

func (response *CouponQuerySetMealListResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

