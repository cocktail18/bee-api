package orderreserveorderquery

/**
* 上门预约订单信息查询
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_reserveorderquery_url = "/ddzhkh/dingdan/reserveorderquery"

type OrderReserveorderqueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ThirdReserveOrderDTO `json:"data"`
    TraceId string `json:"traceId"`
}
type FulfilInfoMap struct {
    /**
    * 取件快递名称
    */
    DeliveryExpressName string `json:"deliveryExpressName"`
    /**
    * 取件快递单号
    */
    DeliveryExpressNumber string `json:"deliveryExpressNumber"`
    /**
    * 取件快递名称
    */
    PickExpressName string `json:"pickExpressName"`
    /**
    * 取件快递单号
    */
    PickExpressNumber string `json:"pickExpressNumber"`
}
type OrderReserveorderqueryRequest struct {
    /**
    *  预约单id列表 
    */
    OrderIds []int64 `json:"orderIds"`
}
type ThirdReserveOrderDTO struct {
    Result []ResultSub `json:"result"`
}
type ResultSub struct {
    /**
    * 预约单id 
    */
    ReserveOrderId int64 `json:"reserveOrderId"`
    /**
    * 统一订单id
    */
    UniOrderId string `json:"uniOrderId"`
    /**
    * 门店名称
    */
    ShopName string `json:"shopName"`
    /**
    * 预约单主状态，1-待接单，2-已接单，3-预约失败，4-取消中 5-已取消，6-预约改期中，7-商家已服务 8-已消费
    */
    Status int32 `json:"status"`
    /**
    * 预约单履约状态
    */
    FulfilStatus int32 `json:"fulfilStatus"`
    /**
    * 预约单预约开始时间
    */
    BookStartTime int64 `json:"bookStartTime"`
    /**
    * 预约单预约结束时间
    */
    BookEndTime int64 `json:"bookEndTime"`
    /**
    * 商品名称
    */
    ProductName string `json:"productName"`
    /**
    * 下单平台,1-大众点评，2-美团
    */
    Platform int32 `json:"platform"`
    /**
    * 预约数量
    */
    ReserveNum int64 `json:"reserveNum"`
    /**
    * 券码信息
    */
    SerialNumbers string `json:"serialNumbers"`
    /**
    * 订单总价
    */
    TotalPrice float64 `json:"totalPrice"`
    /**
    * 履约信息
    */
    FulfilInfoMap FulfilInfoMap `json:"fulfilInfoMap"`
}



func (req *OrderReserveorderqueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderReserveorderqueryResponse, error) {
    resp, err := client.InvokeApi(order_reserveorderquery_url, 59, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderReserveorderqueryResponse
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

func (response *OrderReserveorderqueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

