package foodcatdelete

/**
* 删除菜品分类
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const food_cat_delete_url = "/wmoper/ng/foodop/foodCat/delete"

type FoodCatDeleteResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type FoodCatDeleteRequest struct {
    /**
    *  菜品分类名称 
    */
    CategoryName string `json:"category_name"`
}



func (req *FoodCatDeleteRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*FoodCatDeleteResponse, error) {
    resp, err := client.InvokeApi(food_cat_delete_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response FoodCatDeleteResponse
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

func (response *FoodCatDeleteResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

