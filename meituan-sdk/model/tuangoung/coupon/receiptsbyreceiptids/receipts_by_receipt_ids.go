package receiptsbyreceiptids

/**
* 查询券结算信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const receipts_by_receipt_ids_url = "/tuangou/coupon/queryReceiptsByReceiptIds"

type ReceiptsByReceiptIdsResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}
type ReceiptsByReceiptIdsRequest struct {
}



func (req *ReceiptsByReceiptIdsRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ReceiptsByReceiptIdsResponse, error) {
    resp, err := client.InvokeApi(receipts_by_receipt_ids_url, 1, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ReceiptsByReceiptIdsResponse
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

func (response *ReceiptsByReceiptIdsResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

