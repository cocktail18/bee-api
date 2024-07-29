package querycouponstock

/**
* 查询商家当前剩余可发券数量
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_coupon_stock_url = "/waimai/ng/valueadded/queryCouponStock"

type QueryCouponStockResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type QueryCouponStockRequest struct {
}
type Data struct {
    Total int64 `json:"total"`
    Used int64 `json:"used"`
}



func (req *QueryCouponStockRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryCouponStockResponse, error) {
    resp, err := client.InvokeApi(query_coupon_stock_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryCouponStockResponse
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

func (response *QueryCouponStockResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

