package orderremindreply

/**
* 催单回复接口
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_remind_reply_url = "/waimai/ng/order/remindReply"

type OrderRemindReplyResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type OrderRemindReplyRequest struct {
    Biz string `json:"biz"`
}



func (req *OrderRemindReplyRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderRemindReplyResponse, error) {
    resp, err := client.InvokeApi(order_remind_reply_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderRemindReplyResponse
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

func (response *OrderRemindReplyResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

