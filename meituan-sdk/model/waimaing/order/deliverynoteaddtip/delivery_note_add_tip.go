package deliverynoteaddtip

/**
* 配送单加小费
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const delivery_note_add_tip_url = "/waimai/ng/order/deliveryNoteAddTip"

type DeliveryNoteAddTipResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type DeliveryNoteAddTipRequest struct {
    /**
    *  订单ID 
    */
    OrderId int64 `json:"orderId"`
    /**
    *  小费金额 
    */
    TipAmount float64 `json:"tipAmount"`
}



func (req *DeliveryNoteAddTipRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DeliveryNoteAddTipResponse, error) {
    resp, err := client.InvokeApi(delivery_note_add_tip_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DeliveryNoteAddTipResponse
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

func (response *DeliveryNoteAddTipResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

