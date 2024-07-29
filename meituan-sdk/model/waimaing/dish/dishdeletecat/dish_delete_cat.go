package dishdeletecat

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

const dish_delete_cat_url = "/waimai/dish/deleteCat"

type DishDeleteCatRequest struct {
    /**
    *  菜品分类名称 
    */
    CatName string `json:"catName"`
}
type DishDeleteCatResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *DishDeleteCatRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DishDeleteCatResponse, error) {
    resp, err := client.InvokeApi(dish_delete_cat_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DishDeleteCatResponse
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

func (response *DishDeleteCatResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

