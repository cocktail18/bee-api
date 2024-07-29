package orderqueryorder

/**
* 订单批量查询
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_queryorder_url = "/ddzhkh/dingdan/queryorder"

type Data struct {
    Result []ResultSub `json:"result"`
}
type OrderQueryorderRequest struct {
    /**
    *  需要查询的业务类型，默认为团购1 
    */
    OrderType int32 `json:"orderType"`
    /**
    *  订单创建时间开始（毫秒级unix时间戳，从1970年1月1日（UTC/GMT的午夜）开始所经过的毫秒数，例：1612236808000） 注：订单查询必须过滤创建时间，当不传入add_time的时候会使用buy_success_time作为订单创建时间查询条件,若为轮询所有订单,建议传递改参数,并设置为buy_success_time_from前15分钟 
    */
    AddTimeFrom int64 `json:"addTimeFrom"`
    /**
    *  订单创建时间截止（毫秒级unix时间戳，从1970年1月1日（UTC/GMT的午夜）开始所经过的毫秒数，例：1612236808000） 注：订单查询必须过滤创建时间，当不传入add_time的时候会使用buy_success_time作为订单创建时间查询条件 
    */
    AddTimeTo int64 `json:"addTimeTo"`
    /**
    *  购买成功时间开始（毫秒级unix时间戳，从1970年1月1日（UTC/GMT的午夜）开始所经过的毫秒数，例：1612236808000） 
    */
    BuySuccessTimeFrom int64 `json:"buySuccessTimeFrom"`
    /**
    *  购买成功时间截止 （ 毫秒级unix时间戳，从1970年1月1日（UTC/GMT的午夜）开始所经过的毫秒数，例：1612236808000，时间间隔不能超过三个月 ） 
    */
    BuySuccessTimeTo int64 `json:"buySuccessTimeTo"`
    /**
    *  第pageNo页，默认为1 
    */
    PageNo int32 `json:"pageNo"`
    /**
    *  每页大小，默认20，不得大于50 
    */
    PageSize int32 `json:"pageSize"`
}
type OrderQueryorderResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type ResultSub struct {
    /**
    * 统一订单号 
    */
    OrderId string `json:"orderId"`
    /**
    * 订单业务类型
    */
    OrderType int32 `json:"orderType"`
    /**
    * 购买成功时间 
    */
    BuySuccessTime int64 `json:"buySuccessTime"`
    /**
    * 过期时间
    */
    ExpireTime int64 `json:"expireTime"`
    /**
    * 商品名称
    */
    ProductName string `json:"productName"`
    /**
    * 商品id（套餐ID）
    */
    ProductItemId int64 `json:"productItemId"`
    /**
    * 订单状态，0-可使用，1-已完成（已全部使用/退款），2-已过期（取决于业务方是否有该状态）
    */
    Status int32 `json:"status"`
    /**
    * 退款状态，-1-非退款订单，1-退款中，2-退款成功，3-退款失败
    */
    RefundStatus int32 `json:"refundStatus"`
    /**
    * 用户手机号（加密）
    */
    Mobile string `json:"mobile"`
    /**
    * 关联门店（具体含义咨询相关业务方）
    */
    ShopName string `json:"shopName"`
    /**
    * 关联门店id
    */
    OpPoiId string `json:"opPoiId"`
    /**
    * 1-美团，2-大众点评
    */
    Channel int32 `json:"channel"`
    /**
    * 订单金额
    */
    TotalAmount float64 `json:"totalAmount"`
    /**
    * 商家优惠金额
    */
    ShopAmount float64 `json:"shopAmount"`
    /**
    * 团购商品id(团购ID)
    */
    SpugId int64 `json:"spugId"`
}



func (req *OrderQueryorderRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderQueryorderResponse, error) {
    resp, err := client.InvokeApi(order_queryorder_url, 59, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderQueryorderResponse
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

func (response *OrderQueryorderResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

