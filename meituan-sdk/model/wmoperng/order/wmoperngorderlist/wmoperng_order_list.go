package wmoperngorderlist

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

const wmoperng_order_list_url = "/wmoper/ng/order/queryOrders"

type WmoperngOrderListResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type WmoperngOrderListRequest struct {
    /**
    *  日期不包含时分秒，格式为yyyyMMdd 
    */
    DayInt int32 `json:"dayInt"`
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
    PageSize int32 `json:"pageSize"`
}
type Data struct {
    HashMore bool `json:"hashMore"`
    OrderIds []int64 `json:"orderIds"`
}



func (req *WmoperngOrderListRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*WmoperngOrderListResponse, error) {
    resp, err := client.InvokeApi(wmoperng_order_list_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response WmoperngOrderListResponse
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

func (response *WmoperngOrderListResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

