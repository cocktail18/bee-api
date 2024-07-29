package foodtextcheck

/**
* 菜品标题诊断
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const food_text_check_url = "/design/document/foodTextCheck"

type FoodTextCheckData struct {
    /**
    * 诊断结果
    */
    Result string `json:"result"`
}
type FoodTextCheckRequest struct {
    /**
    * 需要诊断的内容（菜名或者菜品描述）
    */
    Text string `json:"text"`
    /**
    *  0：菜名诊断 1：菜品描述诊断 
    */
    Type string `json:"type"`
}
type FoodTextCheckResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data FoodTextCheckData `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *FoodTextCheckRequest) DoInvoke(client mtclient.MeituanClient) (*FoodTextCheckResponse, error) {
    resp, err := client.InvokeApi(food_text_check_url, 24, "", req)

    if err != nil {
        return nil, err
    }
    var response FoodTextCheckResponse
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

func (response *FoodTextCheckResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

