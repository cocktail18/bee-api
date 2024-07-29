package orderkcreport

/**
* 上报卡餐
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_kc_report_url = "/waimai/ng/order/kc/report"

type OrderKcReportRequest struct {
    Biz string `json:"biz"`
}
type OrderKcReportResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 操作结果
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *OrderKcReportRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderKcReportResponse, error) {
    resp, err := client.InvokeApi(order_kc_report_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderKcReportResponse
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

func (response *OrderKcReportResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

