package supplystockquery

/**
* 查询库存
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const supply_stock_query_url = "/resv2/stock/supply/queryStock"

type SupplyStockQueryRequest struct {
    /**
    *  日期 秒级时间戳 
    */
    StartDay int32 `json:"startDay"`
    /**
    *  日期 秒级时间戳 
    */
    EndDay int32 `json:"endDay"`
}
type SupplyStockQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Map `json:"data"`
    TraceId string `json:"traceId"`
}
type Map struct {
}



func (req *SupplyStockQueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*SupplyStockQueryResponse, error) {
    resp, err := client.InvokeApi(supply_stock_query_url, 7, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response SupplyStockQueryResponse
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

func (response *SupplyStockQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

