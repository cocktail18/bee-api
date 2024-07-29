package batchpullphonenumber

/**
* 隐私号-批量拉取用户手机号
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const batch_pull_phone_number_url = "/waimai/order/batchPullPhoneNumber"

type PhoneNumberInfo struct {
    /**
    * 门店下的订单流水号
    */
    DaySeq int32 `json:"daySeq"`
    /**
    * ERP门店ID 请求参数不传appAuthToken信息，默认拉取开发者下所有未完成订单时，会返回该字段
    */
    EPoiId string `json:"ePoiId"`
    /**
    * 订单号
    */
    OrderId int64 `json:"orderId"`
    /**
    * 订单展示号
    */
    OrderIdView int64 `json:"orderIdView"`
    /**
    * 真实手机号
    */
    RealPhoneNumber string `json:"realPhoneNumber"`
    /**
    * 骑手名字
    */
    RiderName string `json:"riderName"`
}
type BatchPullPhoneNumberRequest struct {
    /**
    *  分页查询的偏移量 
    */
    DegradOffset int32 `json:"degradOffset"`
    DeveloperId int64 `json:"developerId"`
    /**
    *  每页条数，需小于等于1000 
    */
    DegradLimit string `json:"degradLimit"`
}
type BatchPullPhoneNumberResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []PhoneNumberInfo `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *BatchPullPhoneNumberRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*BatchPullPhoneNumberResponse, error) {
    resp, err := client.InvokeApi(batch_pull_phone_number_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response BatchPullPhoneNumberResponse
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

func (response *BatchPullPhoneNumberResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

