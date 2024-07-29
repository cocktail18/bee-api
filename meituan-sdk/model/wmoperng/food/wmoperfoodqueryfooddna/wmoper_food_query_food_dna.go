package wmoperfoodqueryfooddna

/**
* 查询商品DNA
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const wmoper_food_query_food_dna_url = "/wmoper/ng/food/queryFoodDna"

type WmoperFoodQueryFoodDnaData struct {
    /**
    * 后台类目ID
    */
    CategoryId int64 `json:"category_id"`
    /**
    * 模板ID
    */
    TemplateId int64 `json:"template_id"`
    /**
    * 后台类目对应模板下属性父ID
    */
    ParentPropertyId int64 `json:"parent_property_id"`
    /**
    * 后台类目对应模板下属性ID
    */
    Code int64 `json:"code"`
    /**
    * 后台类目对应模板下属性名称
    */
    Name string `json:"name"`
    /**
    * 后台类目对应模板下属性填的值
    */
    Value string `json:"value"`
    /**
    * 后台类目对应模板下属性填的值的ID
    */
    ValueId int64 `json:"value_id"`
    /**
    * 层级
    */
    Level int32 `json:"level"`
    /**
    * 是否叶子节点
    */
    IsLeaf int32 `json:"is_leaf"`
    /**
    * 排序
    */
    Sequence int32 `json:"sequence"`
}
type WmoperFoodQueryFoodDnaRequest struct {
    /**
    *  菜品id 
    */
    AppFoodCode string `json:"app_food_code"`
}
type WmoperFoodQueryFoodDnaResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data WmoperFoodQueryFoodDnaData `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *WmoperFoodQueryFoodDnaRequest) DoInvoke(client mtclient.MeituanClient) (*WmoperFoodQueryFoodDnaResponse, error) {
    resp, err := client.InvokeApi(wmoper_food_query_food_dna_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response WmoperFoodQueryFoodDnaResponse
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

func (response *WmoperFoodQueryFoodDnaResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

