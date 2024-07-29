package ordercreate

/**
* 快驴订单创建
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_create_url = "/kl/open/order/create"

type OrderCreateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}
type GoodsList struct {
    /**
    *  服务商订单明细标识 
    */
    OrderItemIdentify string `json:"orderItemIdentify"`
    /**
    *  服务商商品标识 
    */
    GoodsIdentify string `json:"goodsIdentify"`
    /**
    *  订单购买商品的数量 
    */
    Quantity int32 `json:"quantity"`
}
type OrderCreateRequest struct {
    /**
    *  服务商订单标识 
    */
    OrderIdentify string `json:"orderIdentify"`
    /**
    *  服务商门店标识 
    */
    ShopIdentify string `json:"shopIdentify"`
    /**
    *  服务商关联商家品牌 
    */
    BrandIdentify string `json:"brandIdentify"`
    /**
    *  服务商订单上需要在快驴订单上展示的信息组合 
    */
    Remarks string `json:"remarks"`
    /**
    *  订单商品列表 
    */
    GoodsList []GoodsList `json:"goodsList"`
}



func (req *OrderCreateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderCreateResponse, error) {
    resp, err := client.InvokeApi(order_create_url, 27, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderCreateResponse
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

func (response *OrderCreateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

