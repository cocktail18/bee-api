package groupvoucherorderrefund

/**
* 代金券买单申请退款
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const groupvoucher_order_refund_url = "/tuangou/ng/group_voucher/order/refund"

type GroupVoucherOrderRefundResponse struct {
    Code string `json:"code"`
    Data Data `json:"data"`
    Success bool `json:"success"`
    Message string `json:"message"`
}
type Data struct {
    /**
    * 退款类型，0:商家通过账号退款，1：系统直接退款
    */
    RefundType int64 `json:"refundType"`
    /**
    * 用户订单ID
    */
    SerialNumber string `json:"serialNumber"`
    /**
    * 厂商退款请求地址ip
    */
    VendorIp string `json:"vendorIp"`
    /**
    * 系统订单ID
    */
    OrderId int64 `json:"orderId"`
    /**
    * 退款原因
    */
    RefundReason string `json:"refundReason"`
    /**
    * 退款结果true成功false失败
    */
    RefundResult bool `json:"refundResult"`
    /**
    * 厂商ID
    */
    VendorId string `json:"vendorId"`
    /**
    * 厂商管理员账号id
    */
    VendorShopAdminId string `json:"vendorShopAdminId"`
    /**
    * 美团门店ID
    */
    MtShopId int64 `json:"mtShopId"`
    /**
    * 时间戳
    */
    Timestamp int64 `json:"timestamp"`
}
type GroupvoucherOrderRefundResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data GroupVoucherOrderRefundResponse `json:"data"`
    TraceId string `json:"traceId"`
}
type GroupvoucherOrderRefundRequest struct {
    /**
    *  门店管理员账号id 
    */
    VendorShopAdminId string `json:"vendorShopAdminId"`
    /**
    *  厂商退款请求地址ip 
    */
    VendorIp string `json:"vendorIp"`
    /**
    *  退款类型，0:商家通过开店宝账号退款，1：系统直接退款 
    */
    RefundType int64 `json:"refundType"`
    /**
    *  退款原因 
    */
    RefundReason string `json:"refundReason"`
    /**
    *  用户/开店宝展示的订单ID 
    */
    SerialNumber string `json:"serialNumber"`
    /**
    *  订单ID 
    */
    OrderId int64 `json:"orderId"`
}



func (req *GroupvoucherOrderRefundRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupvoucherOrderRefundResponse, error) {
    resp, err := client.InvokeApi(groupvoucher_order_refund_url, 1, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupvoucherOrderRefundResponse
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

func (response *GroupvoucherOrderRefundResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

