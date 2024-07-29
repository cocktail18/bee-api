package querycategorylist

/**
* 查询所有类目
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_category_list_url = "/waimai/ng/dish/queryCategoryList"

type Category struct {
    CategoryId int64 `json:"categoryId"`
    TemplateId int64 `json:"templateId"`
    /**
    * 分类名称
    */
    Name string `json:"name"`
    /**
    * 对应类目父id
    */
    ParentId int64 `json:"parentId"`
    Child []Category `json:"child"`
}
type QueryCategoryListRequest struct {
}
type QueryCategoryListResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []Category `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *QueryCategoryListRequest) DoInvoke(client mtclient.MeituanClient) (*QueryCategoryListResponse, error) {
    resp, err := client.InvokeApi(query_category_list_url, 2, "", req)

    if err != nil {
        return nil, err
    }
    var response QueryCategoryListResponse
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

func (response *QueryCategoryListResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

