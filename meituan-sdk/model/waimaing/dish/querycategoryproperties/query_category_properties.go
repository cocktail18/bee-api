package querycategoryproperties

/**
* 根据类目查询模板下所有属性
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_category_properties_url = "/waimai/ng/dish/queryCategoryProperties"

type QueryCategoryPropertiesRequest struct {
    /**
    *  分类id 
    */
    CategoryId int64 `json:"categoryId"`
}
type QueryCategoryPropertiesResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Data struct {
    PropertiesKeys []PropertiesKeys `json:"propertiesKeys"`
}
type PropertiesKeys struct {
    /**
    *  后台类目对应模板下属性ID
    */
    Code int64 `json:"code"`
    /**
    * 后台类目对应模板下属性名称
    */
    Name string `json:"name"`
    /**
    * 后台类目对应模板下属性父ID
    */
    ParentPropertyId int64 `json:"parent_property_id"`
    /**
    *  1必填;2非必填
    */
    IsRequired int32 `json:"is_required"`
    /**
    * (1: "单选");(2: "多选");(3: "输入");(4: "单选可输入");(5: "多选可添加");(6: "输入可添加");(7: "单选可添加")
    */
    InputType int32 `json:"input_type"`
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
    * 子节点，结构相同
    */
    Child []PropertiesKeys `json:"child"`
    PromptDocument string `json:"prompt_document"`
}



func (req *QueryCategoryPropertiesRequest) DoInvoke(client mtclient.MeituanClient) (*QueryCategoryPropertiesResponse, error) {
    resp, err := client.InvokeApi(query_category_properties_url, 2, "", req)

    if err != nil {
        return nil, err
    }
    var response QueryCategoryPropertiesResponse
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

func (response *QueryCategoryPropertiesResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

