package zblogisticspreview

/**
* 查询众包配送费
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const zb_logistics_preview_url = "/waimai/ng/order/zbLogisticsPreview"

type ZbLogisticsPreviewResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ZbLogisticsPreviewData `json:"data"`
    TraceId string `json:"traceId"`
}
type ExtMap struct {
    /**
    * title背景颜色
    */
    OrderPageTitleBgColor string `json:"orderPageTitleBgColor"`
    /**
    * 按钮图案
    */
    ButtonImg string `json:"buttonImg"`
    /**
    * 订单图标
    */
    OrderIcon string `json:"orderIcon"`
    /**
    * 订单页面标题颜色
    */
    OrderPageTitleColor string `json:"orderPageTitleColor"`
    /**
    * 权益token
    */
    RightDisToken string `json:"rightDisToken"`
    /**
    * 提示内容
    */
    TipsContent string `json:"tipsContent"`
    /**
    * 活动ID
    */
    ActivityId string `json:"activityId"`
    /**
    * 提示注释
    */
    TipsComment string `json:"tipsComment"`
    /**
    * 提示颜色
    */
    TipsColor string `json:"tipsColor"`
    /**
    * 权益类型code
    */
    RightTypeCode string `json:"rightTypeCode"`
    /**
    * 提示背景颜色
    */
    TipsBgColor string `json:"tipsBgColor"`
    /**
    * 提示图标
    */
    TipsIcon string `json:"tipsIcon"`
    /**
    * 费用详细信息
    */
    FeeDetails []FeeDetail `json:"feeDetails"`
}
type FeeDetail struct {
    /**
    * 抵扣金额
    */
    DiscountValue int32 `json:"discountValue"`
    /**
    * 费用项类型
    */
    FeeType string `json:"feeType"`
    /**
    * 费用项价值
    */
    FeeValue int32 `json:"feeValue"`
    /**
    * 设置抵扣金额
    */
    SetDiscountValue bool `json:"setDiscountValue"`
    /**
    * 设置费用类型
    */
    SetFeeType bool `json:"setFeeType"`
    /**
    * 设置费用值
    */
    SetFeeValue bool `json:"setFeeValue"`
}
type WmUserTaskBMLDetail struct {
    /**
    * 惊喜立减id
    */
    UserTaskId int64 `json:"userTaskId"`
    /**
    * 惊喜立减taskId
    */
    TaskId int64 `json:"taskId"`
    /**
    * 惊喜立减金额 元
    */
    BmlAmount float64 `json:"bmlAmount"`
}
type RightsDetail struct {
    /**
    * 权益token
    */
    RightDisToken string `json:"rightDisToken"`
    ExtMap ExtMap `json:"extMap"`
}
type ZbLogisticsPreviewRequest struct {
    Biz string `json:"biz"`
}
type ZbLogisticsPreviewData struct {
    /**
    * 订单id
    */
    WmOrderId int64 `json:"wm_order_id"`
    /**
    * 配送费（基础+临时）
    */
    ShippingFee float64 `json:"shipping_fee"`
    /**
    * 配送费备注，浮动价格XX元
    */
    ShippingTips string `json:"shipping_tips"`
    /**
    * 实付金额（基础-优惠-活动+临时加价+小费）
    */
    PayAmount float64 `json:"pay_amount"`
    /**
    * 配送费优惠券id
    */
    CouponViewId string `json:"coupon_view_id"`
    /**
    * 优惠券名称
    */
    CouponName string `json:"coupon_name"`
    /**
    * 优惠券金额
    */
    CouponAmount float64 `json:"coupon_amount"`
    /**
    * 配送距离
    */
    Distance int32 `json:"distance"`
    CouponSource int64 `json:"coupon_source"`
    ReduceDetail WmUserTaskBMLDetail `json:"reduce_detail"`
    /**
    * 跑腿溢价权益列表 请求参数recommendRights传true返回
    */
    RightsDetailList []RightsDetail `json:"rights_detail_list"`
}



func (req *ZbLogisticsPreviewRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ZbLogisticsPreviewResponse, error) {
    resp, err := client.InvokeApi(zb_logistics_preview_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ZbLogisticsPreviewResponse
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

func (response *ZbLogisticsPreviewResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

