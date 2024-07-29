package actdiscountbatchsave

/**
* 批量创建或更新折扣商品
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const act_discount_batchsave_url = "/waimai/ng/act/discount/batchsave"

type ActDiscountBatchsaveRequest struct {
    ActData string `json:"actData"`
}
type ActDiscountBatchsaveResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *ActDiscountBatchsaveRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ActDiscountBatchsaveResponse, error) {
    resp, err := client.InvokeApi(act_discount_batchsave_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ActDiscountBatchsaveResponse
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

func (response *ActDiscountBatchsaveResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

