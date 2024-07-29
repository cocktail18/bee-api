package groupvouchershopmanage

/**
* 开通/关闭代金券买单
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const groupvoucher_shop_manage_url = "/tuangou/ng/group_voucher/shop/manage"

type GroupvoucherShopManageResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data GroupVoucherShopManageResponse `json:"data"`
    TraceId string `json:"traceId"`
}
type GroupVoucherShopManageResponse struct {
    Code string `json:"code"`
    Data Data `json:"data"`
    Success bool `json:"success"`
    Message string `json:"message"`
}
type Data struct {
    /**
    * 时间戳
    */
    TimeStamp int64 `json:"timeStamp"`
    /**
    * 是否成功：true成功 false失败
    */
    Result bool `json:"result"`
    /**
    * 厂商ID
    */
    VendorId string `json:"vendorId"`
    /**
    * 美团门店ID
    */
    MtShopId int64 `json:"mtShopId"`
    /**
    * 开通状态:0未合作1开通2关闭
    */
    Status int64 `json:"status"`
}
type GroupvoucherShopManageRequest struct {
    /**
    *  开通状态:1开通2关闭 
    */
    Status int32 `json:"status"`
    /**
    *  厂商管理员账号id 
    */
    VendorShopAdminId string `json:"vendorShopAdminId"`
}



func (req *GroupvoucherShopManageRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupvoucherShopManageResponse, error) {
    resp, err := client.InvokeApi(groupvoucher_shop_manage_url, 1, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupvoucherShopManageResponse
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

func (response *GroupvoucherShopManageResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

