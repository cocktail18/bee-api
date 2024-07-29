package businessmetricsreport

/**
* 商家上报业务指标
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const business_metrics_report_url = "/waimai/ng/order/reportMetrics"

type BusinessMetricsReportRequest struct {
    /**
    *  指标类型，订单打印完成: ORDER_PRINT_COMPLETE 
    */
    Type string `json:"type"`
    /**
    *  指标值，不同的指标要求的value值格式不同 指标类型为 「 ORDER_PRINT_COMPLETE」订单打印完成时 
    */
    Value ValueInfo `json:"value"`
}
type BusinessMetricsReportResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 操作结果
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type ValueInfo struct {
    /**
    *  订单orderViewId 
    */
    OrderViewId string `json:"orderViewId"`
}



func (req *BusinessMetricsReportRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*BusinessMetricsReportResponse, error) {
    resp, err := client.InvokeApi(business_metrics_report_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response BusinessMetricsReportResponse
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

func (response *BusinessMetricsReportResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

