package zhongbaoshippingfeebycode

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

const zhongbao_shippingfeebycode_url = "/waimai/ng/order/zhongbao/shippingFeeByCode"

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
type ZhongbaoShippingfeebycodeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Result `json:"data"`
    TraceId string `json:"traceId"`
}
type ResultData struct {
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
type ShippingFeeByCode struct {
    OrderId string `json:"order_id"`
    LogisticsCode string `json:"logistics_code"`
}
type ZhongbaoShippingfeebycodeRequest struct {
    OrderExtraList []ShippingFeeByCode `json:"orderExtraList"`
}
type Result struct {
    /**
    * 成功列表
    */
    SuccessOrders []ResultData `json:"success_orders"`
    /**
    * 失败列表
    */
    FailOrders []FailOrders `json:"fail_orders"`
}



func (req *ZhongbaoShippingfeebycodeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ZhongbaoShippingfeebycodeResponse, error) {
    resp, err := client.InvokeApi(zhongbao_shippingfeebycode_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ZhongbaoShippingfeebycodeResponse
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

func (response *ZhongbaoShippingfeebycodeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

