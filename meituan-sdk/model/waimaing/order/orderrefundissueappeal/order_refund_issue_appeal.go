package orderrefundissueappeal

/**
* 商家申诉接口
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_refund_issue_appeal_url = "/waimai/ng/order/refund/issue/appeal"

type OrderRefundIssueAppealResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type OrderRefundIssueAppealRequest struct {
    /**
    *  订单Id 
    */
    OrderId int64 `json:"orderId"`
    /**
    *  退款单Id 
    */
    AfterSaleId int64 `json:"afterSaleId"`
    /**
    *  申诉原因 
    */
    Reason string `json:"reason"`
    /**
    *  申诉图片（多张图片逗号分隔）,最多5张 
    */
    AppealPictures string `json:"appealPictures"`
}



func (req *OrderRefundIssueAppealRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderRefundIssueAppealResponse, error) {
    resp, err := client.InvokeApi(order_refund_issue_appeal_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderRefundIssueAppealResponse
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

func (response *OrderRefundIssueAppealResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

