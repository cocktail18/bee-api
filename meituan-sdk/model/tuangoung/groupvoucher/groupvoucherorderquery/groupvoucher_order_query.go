package groupvoucherorderquery

/**
* 查询代金券买单信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const groupvoucher_order_query_url = "/tuangou/ng/group_voucher/order/query"

type GroupVoucherOrderDTO struct {
    /**
    * 用户订单号
    */
    SerialNumber string `json:"serialNumber"`
    /**
    * 消息类型 （1：createOrder订单创建成功 2：payOrderSuccess订单支付成功 3：payOrderError订单支付失败 4：refundingOrder订单退款中 5：refundOrderSuccess订单退款成功 6：refundOrderError订单退款失败）
    */
    MsgType string `json:"msgType"`
    /**
    * 商品ID
    */
    ProductId int64 `json:"productId"`
    /**
    * 系统订单ID
    */
    OrderId int64 `json:"orderId"`
    /**
    * 支付时间
    */
    PayTime int64 `json:"payTime"`
    /**
    * 美团城市ID
    */
    MtCityId int64 `json:"mtCityId"`
    /**
    * 订单状态
    */
    OrderStatus int64 `json:"orderStatus"`
    /**
    * 用户手机号
    */
    MobileNo string `json:"mobileNo"`
    /**
    * 用户ID
    */
    UserId int64 `json:"userId"`
    /**
    * 用户支付金额
    */
    UserPayAmount int64 `json:"userPayAmount"`
    /**
    * 不享优惠金额
    */
    NoDiscountAmount int64 `json:"noDiscountAmount"`
    /**
    * 订单消息体（非结构化）
    */
    MsgContent string `json:"msgContent"`
    /**
    * 订单金额
    */
    OrderAmount int64 `json:"orderAmount"`
    /**
    * 订单更新时间
    */
    OrderUpdateTime int64 `json:"orderUpdateTime"`
    /**
    * 订单创建时间
    */
    OrderAddTime int64 `json:"orderAddTime"`
    /**
    * 代金券抵扣金额
    */
    GroupVoucherDiscountAmount int64 `json:"groupVoucherDiscountAmount"`
    /**
    * 退款详情
    */
    RefundInfo RefundInfo `json:"refundInfo"`
    /**
    * 美团门店ID
    */
    MtShopId int64 `json:"mtShopId"`
    /**
    * 用户类型
    */
    UserType int64 `json:"userType"`
    /**
    * 优惠详情
    */
    CouponInfoList []CouponInfoListSub `json:"couponInfoList"`
    /**
    * 扩展信息
    */
    ExtraInfo Map `json:"extraInfo"`
}
type RefundInfo struct {
    /**
    * 退款成功时间
    */
    RefundSuccessTime int64 `json:"refundSuccessTime"`
    /**
    * 退款原因
    */
    RefundReason string `json:"refundReason"`
    /**
    * 操作者IP
    */
    OperatorIp string `json:"operatorIp"`
    /**
    * 操作者
    */
    Operator string `json:"operator"`
    /**
    * 退款来源
    */
    PlatformSource int64 `json:"platformSource"`
}
type CouponInfoListSub struct {
    CouponInfoDTO CouponInfoDTO `json:"CouponInfoDTO"`
}
type CouponInfoDTO struct {
    /**
    * 优惠抵扣金额
    */
    CouponAmount int64 `json:"couponAmount"`
    /**
    * 优惠类型
    */
    CouponType int64 `json:"couponType"`
    /**
    * 代金券对应的商品ID
    */
    DealId int64 `json:"dealId"`
    /**
    * 扩展信息
    */
    ExtraMessage string `json:"extraMessage"`
    /**
    * 用户实付金额
    */
    RealPayAmount int64 `json:"realPayAmount"`
    /**
    * 优惠类型ID
    */
    CouponIdStr string `json:"couponIdStr"`
}
type GroupvoucherOrderQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data GroupVoucherOrderQueryResponse `json:"data"`
    TraceId string `json:"traceId"`
}
type GroupVoucherOrderQueryResponse struct {
    Code string `json:"code"`
    Data GroupVoucherOrderDTO `json:"data"`
    Success bool `json:"success"`
    Message string `json:"message"`
}
type Map struct {
}
type GroupvoucherOrderQueryRequest struct {
    /**
    *  用户/开店宝展示的订单ID 
    */
    SerialNumber string `json:"serialNumber"`
    /**
    *  订单ID 
    */
    OrderId int64 `json:"orderId"`
}



func (req *GroupvoucherOrderQueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupvoucherOrderQueryResponse, error) {
    resp, err := client.InvokeApi(groupvoucher_order_query_url, 1, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupvoucherOrderQueryResponse
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

func (response *GroupvoucherOrderQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

