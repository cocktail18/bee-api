package apporderlist

/**
* 应用购买记录
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const app_order_list_url = "/wmoper/ng/waimaiad/app/order/list"

type Order struct {
    OrderId int64 `json:"orderId"`
    OrderViewId int64 `json:"orderViewId"`
    AppId int64 `json:"appId"`
    ServiceCode string `json:"serviceCode"`
    PayStatus int64 `json:"payStatus"`
    Version string `json:"version"`
    VersionName string `json:"versionName"`
    Days int64 `json:"days"`
    StartTime int64 `json:"startTime"`
    EndTime int64 `json:"endTime"`
    RightStatus int64 `json:"rightStatus"`
    RealPayPrice int64 `json:"realPayPrice"`
    OrderTime int64 `json:"orderTime"`
}
type AppOrderListRequest struct {
    /**
    *  应用id 
    */
    AppId int64 `json:"appId"`
    /**
    *  数据的时间(以下单时间为准)范围，开始时间 
    */
    StartTime int32 `json:"startTime"`
    /**
    *  数据的时间(以下单时间为准)范围，结束时间 
    */
    EndTime int32 `json:"endTime"`
    /**
    *  页码 
    */
    PageNum int32 `json:"pageNum"`
    /**
    *  每页记录数 
    */
    PageSize int32 `json:"pageSize"`
}
type Data struct {
    Result Result `json:"result"`
    Total int64 `json:"total"`
    Order []Order `json:"order"`
}
type AppOrderListResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Result struct {
    Success bool `json:"success"`
    Code string `json:"code"`
    Message string `json:"message"`
}



func (req *AppOrderListRequest) DoInvoke(client mtclient.MeituanClient) (*AppOrderListResponse, error) {
    resp, err := client.InvokeApi(app_order_list_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response AppOrderListResponse
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

func (response *AppOrderListResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

