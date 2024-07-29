package wmoperfoodquerycategorylist

/**
* 查询所有类目，对应商家后台标准分类
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const wmoper_food_query_category_list_url = "/wmoper/ng/food/queryCategoryList"

type TargetObj struct {
    /**
    * 分类名称
    */
    Name string `json:"name"`
    /**
    * 对应类目模板ID
    */
    TemplateId int64 `json:"template_id"`
    /**
    * 对应类目id
    */
    CategoryId int64 `json:"category_id"`
    /**
    * 对应类目父id
    */
    ParentId int64 `json:"parentId"`
    /**
    * 子节点，结构相同
    */
    Child []TargetObj `json:"child"`
}
type WmoperFoodQueryCategoryListResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []TargetObj `json:"data"`
    TraceId string `json:"traceId"`
}
type WmoperFoodQueryCategoryListRequest struct {
}



func (req *WmoperFoodQueryCategoryListRequest) DoInvoke(client mtclient.MeituanClient) (*WmoperFoodQueryCategoryListResponse, error) {
    resp, err := client.InvokeApi(wmoper_food_query_category_list_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response WmoperFoodQueryCategoryListResponse
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

func (response *WmoperFoodQueryCategoryListResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

