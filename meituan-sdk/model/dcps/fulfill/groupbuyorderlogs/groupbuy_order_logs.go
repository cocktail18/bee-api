package groupbuyorderlogs

/**
* 套餐配送-查询轨迹
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const groupbuy_order_logs_url = "/dcps/fulfill/order/logs"

type LogDetail struct {
    OrderId int64 `json:"orderId"`
    Type int32 `json:"type"`
    Status int32 `json:"status"`
    OldStatus int32 `json:"oldStatus"`
    Describe string `json:"describe"`
    OperatorSource int32 `json:"operatorSource"`
    AddTime int64 `json:"addTime"`
    BookingOrderId string `json:"bookingOrderId"`
}
type Data struct {
    Message string `json:"message"`
    Code int32 `json:"code"`
    LogDetails []LogDetail `json:"logDetails"`
}
type GroupbuyOrderLogsRequest struct {
    BookingOrderId string `json:"bookingOrderId"`
    OpIp string `json:"opIp"`
    OpPlatform string `json:"opPlatform"`
    OpUuid string `json:"opUuid"`
}
type GroupbuyOrderLogsResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *GroupbuyOrderLogsRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupbuyOrderLogsResponse, error) {
    resp, err := client.InvokeApi(groupbuy_order_logs_url, 46, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupbuyOrderLogsResponse
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

func (response *GroupbuyOrderLogsResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

