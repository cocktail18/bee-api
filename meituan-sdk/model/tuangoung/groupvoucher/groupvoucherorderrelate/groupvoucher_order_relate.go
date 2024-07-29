package groupvoucherorderrelate

/**
* 代金券买单信息与厂商收银订单（或者收银键位）关联
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const groupvoucher_order_relate_url = "/tuangou/ng/group_voucher/order/relate"

type GroupvoucherOrderRelateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data GroupVoucherOrderRelateResponse `json:"data"`
    TraceId string `json:"traceId"`
}
type GroupvoucherOrderRelateRequest struct {
    /**
    *  是否关联厂商订单,true关联false不关联 
    */
    IsRelate bool `json:"isRelate"`
    /**
    *  用户或开店宝展示的订单ID 
    */
    SerialNumber string `json:"serialNumber"`
    /**
    *  订单ID 
    */
    OrderId int64 `json:"orderId"`
    /**
    *  厂商管理员账号id 
    */
    VendorShopAdminId string `json:"vendorShopAdminId"`
    /**
    *  厂商桌台号 
    */
    TableNo string `json:"tableNo"`
    /**
    *  厂商订单号 
    */
    VendorOrderId string `json:"vendorOrderId"`
}
type GroupVoucherOrderRelateResponse struct {
    Code string `json:"code"`
    Data Data `json:"data"`
    Success bool `json:"success"`
    Message string `json:"message"`
}
type Data struct {
    /**
    * 用户订单号
    */
    SerialNumber string `json:"serialNumber"`
    /**
    * 是否关联厂商订单,true关联false不关联
    */
    IsRelate bool `json:"isRelate"`
    /**
    * 系统订单号
    */
    OrderId int64 `json:"orderId"`
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
    /**
    * 关联的结果
    */
    IsRelateResult bool `json:"isRelateResult"`
    /**
    * 厂商收银订单ID
    */
    VendorOrderId string `json:"vendorOrderId"`
    /**
    * 桌台号
    */
    TableNo string `json:"tableNo"`
}



func (req *GroupvoucherOrderRelateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupvoucherOrderRelateResponse, error) {
    resp, err := client.InvokeApi(groupvoucher_order_relate_url, 1, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupvoucherOrderRelateResponse
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

func (response *GroupvoucherOrderRelateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

