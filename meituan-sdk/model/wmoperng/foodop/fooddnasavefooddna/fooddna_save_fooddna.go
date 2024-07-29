package fooddnasavefooddna

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

const fooddna_save_fooddna_url = "/wmoper/ng/foodop/foodDna/saveFoodDna"

type FooddnaSaveFooddnaResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type FooddnaSaveFooddnaRequest struct {
    /**
    *  dna信息，json字符串： [ { "app_food_code": "FOOD_0005", "wmProductSpuExtendList": [ { "category_id": 23, // 后台类目ID "code": "1000000012", //后台类目对应模板下属性ID "is_leaf": 1, // 是否叶子节点 "level": 0, // 层级 "name": "荤素", // 后台类目对应模板下属性名称 "parent_property_id": 0, // 后台类目对应模板下属性父ID "sequence": 0, // 排序 "template_id": 10, // 模板ID "value": "", // 后台类目对应模板下属性填的值 "value_id": 0 // 后台类目对应模板下属性填的值的ID } ] } ] 
    */
    FoodData string `json:"food_data"`
}



func (req *FooddnaSaveFooddnaRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*FooddnaSaveFooddnaResponse, error) {
    resp, err := client.InvokeApi(fooddna_save_fooddna_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response FooddnaSaveFooddnaResponse
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

func (response *FooddnaSaveFooddnaResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

