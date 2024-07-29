package cancelzblogisticsbywmorderid

/**
* 提交取消跑腿
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const cancel_zb_logistics_by_wm_order_id_url = "/waimai/order/cancelZbLogisticsByWmOrderId"

type CancelZbLogisticsByWmOrderIdResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 操作结果
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type CancelZbLogisticsByWmOrderIdRequest struct {
    /**
    *  具体原因 
    */
    DetailContent string `json:"detailContent"`
    /**
    *  订单号 
    */
    OrderId string `json:"orderId"`
    /**
    *  原因代码 
    */
    ReasonCode string `json:"reasonCode"`
}



func (req *CancelZbLogisticsByWmOrderIdRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CancelZbLogisticsByWmOrderIdResponse, error) {
    resp, err := client.InvokeApi(cancel_zb_logistics_by_wm_order_id_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CancelZbLogisticsByWmOrderIdResponse
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

func (response *CancelZbLogisticsByWmOrderIdResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

