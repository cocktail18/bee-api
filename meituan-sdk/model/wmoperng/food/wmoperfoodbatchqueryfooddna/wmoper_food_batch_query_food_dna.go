package wmoperfoodbatchqueryfooddna

/**
* 批量查询商品DNA
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const wmoper_food_batch_query_food_dna_url = "/wmoper/ng/food/batchQueryFoodDna"

type WmoperFoodBatchQueryFoodDnaResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data WmoperFoodBatchQueryFoodDnaData `json:"data"`
    TraceId string `json:"traceId"`
}
type WmoperFoodBatchQueryFoodDnaData struct {
    /**
    * 商品DNA信息
    */
    Data []DNAData `json:"data"`
}
type DNAData struct {
    /**
    * 后台类目id
    */
    CategoryId int64 `json:"categoryId"`
    /**
    * 菜品id
    */
    AppFoodCode string `json:"app_food_code"`
    /**
    * 模板信息
    */
    PropertiesValues []Property `json:"properties_values"`
}
type WmoperFoodBatchQueryFoodDnaRequest struct {
    /**
    *  菜品id列表，使用逗号隔开，最多200个。 
    */
    EDishCodes string `json:"eDishCodes"`
}
type Property struct {
    /**
    * 后台类目对应模板下属性ID
    */
    Code int64 `json:"code"`
    /**
    * 模板ID
    */
    TemplateId int64 `json:"template_id"`
    /**
    * 后台类目id
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
    Level int32 `json:"level"`
    /**
    * 是否叶子节点
    */
    IsLeaf int32 `json:"is_leaf"`
    /**
    * 排序
    */
    Sequence int32 `json:"sequence"`
    /**
    * 子节点
    */
    Child []string `json:"child"`
}



func (req *WmoperFoodBatchQueryFoodDnaRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*WmoperFoodBatchQueryFoodDnaResponse, error) {
    resp, err := client.InvokeApi(wmoper_food_batch_query_food_dna_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response WmoperFoodBatchQueryFoodDnaResponse
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

func (response *WmoperFoodBatchQueryFoodDnaResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

