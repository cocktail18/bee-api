package groupbuyagreerefund

/**
* 套餐配送-同意退款
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const groupbuy_agree_refund_url = "/dcps/fulfill/agree/refund"

type GroupbuyAgreeRefundRequest struct {
    BookingOrderId string `json:"bookingOrderId"`
    OpIp string `json:"opIp"`
    OpPlatform string `json:"opPlatform"`
    OpUuid string `json:"opUuid"`
}
type Data struct {
    BookingOrderId string `json:"bookingOrderId"`
    OpKey string `json:"opKey"`
}
type GroupbuyAgreeRefundResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *GroupbuyAgreeRefundRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupbuyAgreeRefundResponse, error) {
    resp, err := client.InvokeApi(groupbuy_agree_refund_url, 46, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupbuyAgreeRefundResponse
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

func (response *GroupbuyAgreeRefundResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

