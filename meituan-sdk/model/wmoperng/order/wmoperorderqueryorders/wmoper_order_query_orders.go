package wmoperorderqueryorders

/**
* 查询订单列表
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const wmoper_order_query_orders_url = "/wmoper/order/queryOrders"

type WmoperOrderQueryOrdersData struct {
    /**
    * 是否还有数据
    */
    HasMore bool `json:"hasMore"`
    /**
    * 订单ID列表
    */
    OrderIds []int64 `json:"orderIds"`
}
type WmoperOrderQueryOrdersResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data WmoperOrderQueryOrdersData `json:"data"`
    TraceId string `json:"traceId"`
}
type WmoperOrderQueryOrdersRequest struct {
    /**
    *  日期格式，yyyy-MM-dd 
    */
    Date string `json:"date"`
    /**
    *  订单状态 -1代表全部 2代表新订单 4 代表已确认订单 8订单完成 9 订单取消 
    */
    OrderStatus int32 `json:"orderStatus"`
    /**
    *  页码 
    */
    PageNo int32 `json:"pageNo"`
    /**
    *  分页条数 ，最大20 
    */
    PageSize string `json:"pageSize"`
}



func (req *WmoperOrderQueryOrdersRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*WmoperOrderQueryOrdersResponse, error) {
    resp, err := client.InvokeApi(wmoper_order_query_orders_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response WmoperOrderQueryOrdersResponse
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

func (response *WmoperOrderQueryOrdersResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

