package querysettleresult

/**
* 查询结算结果
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_settle_result_url = "/dcps/settlement/querySettleResultForTGDeliver"

type QuerySettleResultRequest struct {
    /**
    *  订单ID 
    */
    OrderId int64 `json:"orderId"`
}
type QuerySettleResultResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []SettleDTO `json:"data"`
    TraceId string `json:"traceId"`
}
type SettleDTO struct {
}



func (req *QuerySettleResultRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QuerySettleResultResponse, error) {
    resp, err := client.InvokeApi(query_settle_result_url, 46, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QuerySettleResultResponse
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

func (response *QuerySettleResultResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

