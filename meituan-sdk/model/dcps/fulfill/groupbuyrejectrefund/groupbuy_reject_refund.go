package groupbuyrejectrefund

/**
* 套餐配送-拒绝退款
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const groupbuy_reject_refund_url = "/dcps/fulfill/reject/refund"

type GroupbuyRejectRefundResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type GroupbuyRejectRefundRequest struct {
    BookingOrderId string `json:"bookingOrderId"`
    OpIp string `json:"opIp"`
    OpPlatform string `json:"opPlatform"`
    OpUuid string `json:"opUuid"`
    RefuseReason string `json:"refuseReason"`
}
type Data struct {
    BookingOrderId string `json:"bookingOrderId"`
    OpKey string `json:"opKey"`
}



func (req *GroupbuyRejectRefundRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupbuyRejectRefundResponse, error) {
    resp, err := client.InvokeApi(groupbuy_reject_refund_url, 46, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupbuyRejectRefundResponse
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

func (response *GroupbuyRejectRefundResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

