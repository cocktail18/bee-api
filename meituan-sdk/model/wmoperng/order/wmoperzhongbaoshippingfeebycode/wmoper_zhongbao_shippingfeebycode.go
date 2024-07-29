package wmoperzhongbaoshippingfeebycode

/**
* 批量查询跑腿配送费
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const wmoper_zhongbao_shippingfeebycode_url = "/wmoper/ng/order/zhongbao/shippingFeeByCode"

type FailOrders struct {
    /**
    * 订单ID
    */
    WmOrderId int64 `json:"wm_order_id"`
    /**
    * 错误原因
    */
    ErrorMsg string `json:"error_msg"`
}
type WmoperZhongbaoShippingfeebycodeRequest struct {
    OrderExtraList []OrderExtraList `json:"orderExtraList"`
}
type OrderExtraList struct {
    OrderId string `json:"order_id"`
    LogisticsCode string `json:"logistics_code"`
}
type ResultData struct {
    /**
    * 成功列表
    */
    SuccessOrders []SuccessOrders `json:"success_orders"`
    /**
    * 失败列表
    */
    FailOrders []FailOrders `json:"fail_orders"`
}
type SuccessOrders struct {
    /**
    * 订单ID
    */
    WmOrderId int64 `json:"wm_order_id"`
    /**
    * 配送费
    */
    ShippingFee float64 `json:"shipping_fee"`
    /**
    * 订单展示号
    */
    WmOrderViewId int64 `json:"wm_order_view_id"`
    /**
    * 配送费备注
    */
    ShippingTips string `json:"shipping_tips"`
    /**
    * 美团跑腿服务品牌code
    */
    LogisticsCode string `json:"logistics_code"`
}
type WmoperZhongbaoShippingfeebycodeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *WmoperZhongbaoShippingfeebycodeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*WmoperZhongbaoShippingfeebycodeResponse, error) {
    resp, err := client.InvokeApi(wmoper_zhongbao_shippingfeebycode_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response WmoperZhongbaoShippingfeebycodeResponse
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

func (response *WmoperZhongbaoShippingfeebycodeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

