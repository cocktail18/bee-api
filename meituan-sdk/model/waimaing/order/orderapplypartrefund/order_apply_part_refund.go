package orderapplypartrefund

/**
* 部分退款-申请部分退款
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_apply_part_refund_url = "/waimai/order/applyPartRefund"

type OrderApplyPartRefundResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 操作结果
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type OrderApplyPartRefundRequest struct {
    /**
    *  申请部分退款的具体原因 
    */
    Reason string `json:"reason"`
    /**
    *  订单号 
    */
    OrderId int64 `json:"orderId"`
    /**
    *  部分退款菜品详情 
    */
    FoodData string `json:"foodData"`
}



func (req *OrderApplyPartRefundRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderApplyPartRefundResponse, error) {
    resp, err := client.InvokeApi(order_apply_part_refund_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderApplyPartRefundResponse
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

func (response *OrderApplyPartRefundResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

