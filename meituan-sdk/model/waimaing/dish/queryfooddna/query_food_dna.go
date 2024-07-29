package queryfooddna

/**
* 查询菜品DNA
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_food_dna_url = "/waimai/ng/dish/queryFoodDna"

type QueryFoodDnaResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data QueryFoodDnaData `json:"data"`
    TraceId string `json:"traceId"`
}
type QueryFoodDnaRequest struct {
    Biz string `json:"biz"`
}
type PropertiesValues struct {
    /**
    * 后台类目对应模板下属性ID
    */
    Code int64 `json:"code"`
    /**
    * 模板ID
    */
    TemplateId int64 `json:"template_id"`
    /**
    * 后台类目ID
    */
    CategoryId int64 `json:"category_id"`
    /**
    * 后台类目对应模板下属性父ID
    */
    ParentPropertyId int64 `json:"parent_property_id"`
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
    Level int64 `json:"level"`
    /**
    * 是否叶子节点
    */
    IsLeaf int64 `json:"is_leaf"`
    /**
    * 排序
    */
    Sequence int64 `json:"sequence"`
    Child []PropertiesValues `json:"child"`
}
type QueryFoodDnaData struct {
    /**
    * 后台类目ID
    */
    CategoryId int64 `json:"categoryId"`
    Customizable int64 `json:"customizable"`
    PropertiesValues []PropertiesValues `json:"properties_values"`
}



func (req *QueryFoodDnaRequest) DoInvoke(client mtclient.MeituanClient) (*QueryFoodDnaResponse, error) {
    resp, err := client.InvokeApi(query_food_dna_url, 2, "", req)

    if err != nil {
        return nil, err
    }
    var response QueryFoodDnaResponse
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

func (response *QueryFoodDnaResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

