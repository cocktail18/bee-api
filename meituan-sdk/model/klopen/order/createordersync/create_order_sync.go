package createordersync

/**
* 快驴同步创建订单
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const create_order_sync_url = "/kl/open/order/create/sync"

type CreateOrderSyncResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 快驴订单号
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type GoodsList struct {
    /**
    *  sku编码 
    */
    SkuCode string `json:"skuCode"`
    /**
    *  sku名称 
    */
    SkuName string `json:"skuName"`
    /**
    *  数量 
    */
    Quantity int32 `json:"quantity"`
}
type CreateOrderSyncRequest struct {
    /**
    *  服务商订单标识 
    */
    OrderSn string `json:"orderSn"`
    /**
    *  快驴门店码 
    */
    MerchantCode string `json:"merchantCode"`
    /**
    *  服务商关联商家品牌 
    */
    BrandIdentify string `json:"brandIdentify"`
    /**
    *  服务商订单上需要在快驴订单上展示的信息组合 
    */
    Remark string `json:"remark"`
    /**
    *  加盟店0代表创建订单，1代表支付订单 直营店1代表创建订单 
    */
    Type int32 `json:"type"`
    /**
    *  订单商品列表 
    */
    CartProducts []GoodsList `json:"cartProducts"`
    /**
    *  订单金额 
    */
    TotalAmount string `json:"totalAmount"`
}



func (req *CreateOrderSyncRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CreateOrderSyncResponse, error) {
    resp, err := client.InvokeApi(create_order_sync_url, 27, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CreateOrderSyncResponse
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

func (response *CreateOrderSyncResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

