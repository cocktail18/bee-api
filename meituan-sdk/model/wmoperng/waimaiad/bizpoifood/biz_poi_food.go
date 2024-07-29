package bizpoifood

/**
* 经营分析-菜品信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const biz_poi_food_url = "/wmoper/ng/waimaiad/biz/foodlist"

type BizPoiFoodResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Data struct {
    Result Result `json:"result"`
    FoodList []FoodInfo `json:"foodList"`
}
type FoodInfo struct {
    WmFoodId int64 `json:"wmFoodId"`
    FoodName string `json:"foodName"`
    FoodPrice int64 `json:"foodPrice"`
    OriginFoodPrice int64 `json:"originFoodPrice"`
    FoodSaleCnt int64 `json:"foodSaleCnt"`
    FoodSaleAmt int64 `json:"foodSaleAmt"`
    Dt int64 `json:"dt"`
}
type Result struct {
    Success bool `json:"success"`
    Code string `json:"code"`
    Message string `json:"message"`
}
type BizPoiFoodRequest struct {
    /**
    * 日期，yyyyMMdd格式
    */
    Dt string `json:"dt"`
}



func (req *BizPoiFoodRequest) DoInvoke(client mtclient.MeituanClient) (*BizPoiFoodResponse, error) {
    resp, err := client.InvokeApi(biz_poi_food_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response BizPoiFoodResponse
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

func (response *BizPoiFoodResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

