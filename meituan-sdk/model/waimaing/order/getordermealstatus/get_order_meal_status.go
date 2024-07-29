package getordermealstatus

/**
* 查询出餐超时的订单
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const get_order_meal_status_url = "/waimai/ng/order/getOrderMealStatus"

type ResultData struct {
    /**
    * 结果数量
    */
    Total int32 `json:"total"`
    /**
    * 出餐超时订单明细
    */
    OrderFoodDoneAssessments []OrderFoodDoneAssessments `json:"orderFoodDoneAssessments"`
}
type OrderFoodDoneAssessments struct {
    /**
    * 订单Id
    */
    OrderId int64 `json:"order_id"`
    /**
    * 出餐超时原因：1-未上报出餐 2-上报出餐行为被判定虚假 3-上报出餐且行为真实，但未在考核时间内上报 4-豁免
    */
    Status int32 `json:"status"`
}
type GetOrderMealStatusRequest struct {
    OrderTimeStart int32 `json:"orderTimeStart"`
    OrderTimeEnd int32 `json:"orderTimeEnd"`
    /**
    *  页数 
    */
    PageNum int32 `json:"pageNum"`
    /**
    *  每页数量 
    */
    PageSize int32 `json:"pageSize"`
}
type GetOrderMealStatusResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *GetOrderMealStatusRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GetOrderMealStatusResponse, error) {
    resp, err := client.InvokeApi(get_order_meal_status_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GetOrderMealStatusResponse
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

func (response *GetOrderMealStatusResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

