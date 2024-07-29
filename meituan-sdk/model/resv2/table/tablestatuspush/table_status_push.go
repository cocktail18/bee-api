package tablestatuspush

/**
* ERB桌态状态变更推送
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const table_status_push_url = "/resv2/table/status/erbpush"

type TableStatusPushResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 返回数据
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type TableStatusPushRequest struct {
    Data []Data `json:"data"`
}
type Data struct {
    /**
    *  桌位id 
    */
    TableId string `json:"tableId"`
    /**
    *  第三方订单id唯一标示，用于区分桌态订单 
    */
    BizOrderId string `json:"bizOrderId"`
    /**
    *  锁台开始时间 unix时间戳 
    */
    StartTime int32 `json:"startTime"`
    /**
    *  锁台结束时间 unix时间戳 
    */
    EndTime int32 `json:"endTime"`
    /**
    *  桌位状态 1：空台 2：已开台 3：已点菜 4：已结账 5：待清台 6: 已预结 
    */
    Status int32 `json:"status"`
    /**
    *  桌台对应的点菜数据 
    */
    Dishes []Dish `json:"dishes"`
    /**
    *  订单支付金额 
    */
    Payment float64 `json:"payment"`
}
type Dish struct {
    /**
    *  商家菜品id 
    */
    DishId string `json:"dishId"`
    /**
    *  菜品速记code 
    */
    DishCode string `json:"dishCode"`
    /**
    *  菜品名称 
    */
    DishName string `json:"dishName"`
    /**
    *  "份","斤","盒"等等 
    */
    DishSpec string `json:"dishSpec"`
    /**
    *  菜品价格 
    */
    DishPrice float64 `json:"dishPrice"`
    /**
    *  菜品数量 
    */
    DishCount int32 `json:"dishCount"`
}



func (req *TableStatusPushRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*TableStatusPushResponse, error) {
    resp, err := client.InvokeApi(table_status_push_url, 7, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response TableStatusPushResponse
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

func (response *TableStatusPushResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

