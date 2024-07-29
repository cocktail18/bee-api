package actdiscountdelete

/**
* 批量删除折扣商品
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const act_discount_delete_url = "/waimai/ng/act/discount/delete"

type ActDiscountDeleteResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type ActDiscountDeleteRequest struct {
    /**
    *  APP方商品id，删除商品数量上限为100，若需删除门店中的多个商品折扣活动，请用半角逗号分隔。 
    */
    AppFoodCodes string `json:"appFoodCodes"`
}



func (req *ActDiscountDeleteRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ActDiscountDeleteResponse, error) {
    resp, err := client.InvokeApi(act_discount_delete_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ActDiscountDeleteResponse
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

func (response *ActDiscountDeleteResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

