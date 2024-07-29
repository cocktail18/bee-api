package zbdispatch

/**
* 众包配送-发配送
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const zb_dispatch_url = "/wmoper/ng/order/zbDispatch"

type FeeDetails struct {
    /**
    *  折扣金额 
    */
    DiscountValue float64 `json:"discountValue"`
    /**
    *  费用项类型 
    */
    FeeType string `json:"feeType"`
    /**
    *  费用项价值 
    */
    FeeValue float64 `json:"feeValue"`
    /**
    *  设置抵扣金额 
    */
    SetDiscountValue bool `json:"setDiscountValue"`
    /**
    *  设置费用类型 
    */
    SetFeeType bool `json:"setFeeType"`
    /**
    *  设置费用值 
    */
    SetFeeValue bool `json:"setFeeValue"`
}
type ExtMap struct {
    /**
    *  title背景颜色 
    */
    OrderPageTitleBgColor string `json:"orderPageTitleBgColor"`
    /**
    *  按钮图案 
    */
    ButtonImg string `json:"buttonImg"`
    /**
    *  订单图标 
    */
    OrderIcon string `json:"orderIcon"`
    /**
    *  订单页面标题颜色 
    */
    OrderPageTitleColor string `json:"orderPageTitleColor"`
    /**
    *  权益token 
    */
    RightDisToken string `json:"rightDisToken"`
    /**
    *  提示内容 
    */
    TipsContent string `json:"tipsContent"`
    /**
    *  活动Id 
    */
    ActivityId string `json:"activityId"`
    /**
    *  提示内容 
    */
    TipsComment string `json:"tipsComment"`
    /**
    *  提示颜色 
    */
    TipsColor string `json:"tipsColor"`
    /**
    *  费用详细信息 
    */
    FeeDetails []FeeDetails `json:"feeDetails"`
    /**
    *  权益类型code 
    */
    RightTypeCode string `json:"rightTypeCode"`
    /**
    *  提示背景颜色 
    */
    TipsBgColor string `json:"tipsBgColor"`
    /**
    *  提示图标 
    */
    TipsIcon string `json:"tipsIcon"`
}
type ZbDispatchRequest struct {
    /**
    *  订单号 
    */
    OrderId string `json:"orderId"`
    /**
    *  配送费 
    */
    ShippingFee float64 `json:"shippingFee"`
    /**
    *  小费 
    */
    TipAmount float64 `json:"tipAmount"`
    /**
    *  配送优惠券id 
    */
    CouponViewId string `json:"couponViewId"`
    /**
    *  惊喜立减字段 
    */
    ReduceDetail string `json:"reduceDetail"`
    /**
    *  跑腿配送服务品牌 1003-美团跑腿经济送；2011-美团跑腿蛋糕送 ，默认为1003 
    */
    LogisticsCode string `json:"logisticsCode"`
    /**
    *  使用权益列表 众包发配送接口查询返回字段rights_detail_List 
    */
    RightsDetailList []RightsDetail `json:"rightsDetailList"`
}
type RightsDetail struct {
    /**
    *  权益token 
    */
    RightDisToken string `json:"rightDisToken"`
    ExtMap ExtMap `json:"extMap"`
}
type ZbDispatchResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 操作结果
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *ZbDispatchRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ZbDispatchResponse, error) {
    resp, err := client.InvokeApi(zb_dispatch_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ZbDispatchResponse
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

func (response *ZbDispatchResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

