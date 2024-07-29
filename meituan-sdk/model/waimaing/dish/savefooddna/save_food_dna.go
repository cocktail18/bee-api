package savefooddna

/**
* 保存商品DNA
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const save_food_dna_url = "/waimai/ng/dish/saveFoodDna"

type SaveFoodDnaResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type SaveFoodDnaRequest struct {
    Biz string `json:"biz"`
}



func (req *SaveFoodDnaRequest) DoInvoke(client mtclient.MeituanClient) (*SaveFoodDnaResponse, error) {
    resp, err := client.InvokeApi(save_food_dna_url, 2, "", req)

    if err != nil {
        return nil, err
    }
    var response SaveFoodDnaResponse
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

func (response *SaveFoodDnaResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

