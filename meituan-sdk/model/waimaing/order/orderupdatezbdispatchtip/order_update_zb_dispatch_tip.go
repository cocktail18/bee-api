package orderupdatezbdispatchtip

/**
* 众包配送场景－配送单加小费
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_update_zb_dispatch_tip_url = "/waimai/order/updateZbDispatchTip"

type OrderUpdateZbDispatchTipRequest struct {
    /**
    *  订单号 
    */
    OrderId int64 `json:"orderId"`
    /**
    *  小费 
    */
    TipAmount float64 `json:"tipAmount"`
}
type OrderUpdateZbDispatchTipResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 操作结果
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *OrderUpdateZbDispatchTipRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderUpdateZbDispatchTipResponse, error) {
    resp, err := client.InvokeApi(order_update_zb_dispatch_tip_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderUpdateZbDispatchTipResponse
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

func (response *OrderUpdateZbDispatchTipResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

