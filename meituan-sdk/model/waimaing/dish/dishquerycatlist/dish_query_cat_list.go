package dishquerycatlist

/**
* 查询菜品分类
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const dish_query_cat_list_url = "/waimai/dish/queryCatList"

type DishCategoryInfo struct {
    /**
    * 分类描述
    */
    CategoryDescription string `json:"categoryDescription"`
    /**
    * 分组信息。0-普通分类；1-必选分类；2-可单独结算分类
    */
    CategoryMode int32 `json:"categoryMode"`
    /**
    * ERP门店id 最大长度100
    */
    EPoiId string `json:"ePoiId"`
    /**
    * 菜品分类名称
    */
    Name string `json:"name"`
    /**
    * 分类顺序,排序由小到大
    */
    Sequence int32 `json:"sequence"`
    /**
    * 置顶时间段
    */
    TimeZone string `json:"timeZone"`
    /**
    * 置顶开关，0 关闭； 1 开启
    */
    TopFlag int32 `json:"topFlag"`
}
type DishQueryCatListResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []DishCategoryInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type DishQueryCatListRequest struct {
}



func (req *DishQueryCatListRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DishQueryCatListResponse, error) {
    resp, err := client.InvokeApi(dish_query_cat_list_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DishQueryCatListResponse
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

func (response *DishQueryCatListResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

