package actdiscountstock

/**
* 批量更新折扣商品当日活动库存
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const act_discount_stock_url = "/waimai/ng/act/discount/stock"

type ActDiscountStockResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type ActDiscountStockRequest struct {
    /**
    *  活动数据数量上限为200 
    */
    ActData string `json:"actData"`
}



func (req *ActDiscountStockRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ActDiscountStockResponse, error) {
    resp, err := client.InvokeApi(act_discount_stock_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ActDiscountStockResponse
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

func (response *ActDiscountStockResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

