package wmoperfoodqueryfoodcatlist

/**
* 查询门店菜品分类列表
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const wmoper_food_query_foodcat_list_url = "/wmoper/ng/food/queryFoodCatList"

type WmoperFoodQueryFoodcatListResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Data struct {
    /**
    * APP方门店id
    */
    EpoiId string `json:"epoiId"`
    /**
    * 品类名称
    */
    Name string `json:"name"`
    /**
    * 菜品分类排序
    */
    Sequence int32 `json:"sequence"`
    /**
    * 分类描述
    */
    CategoryDescription string `json:"category_description"`
    /**
    * 分组信息。0-普通分类；1-必选分类；2-可单独结算分类
    */
    CategoryMode int32 `json:"category_mode"`
    /**
    * 置顶开关，0 关闭； 1 开启
    */
    TopFlag int32 `json:"top_flag"`
    /**
    * 置顶时间段
    */
    TimeZone string `json:"time_zone"`
    /**
    * 分类创建时间
    */
    Ctime int32 `json:"ctime"`
    /**
    * 分类更新时间
    */
    Utime int32 `json:"utime"`
}
type WmoperFoodQueryFoodcatListRequest struct {
}



func (req *WmoperFoodQueryFoodcatListRequest) DoInvoke(client mtclient.MeituanClient) (*WmoperFoodQueryFoodcatListResponse, error) {
    resp, err := client.InvokeApi(wmoper_food_query_foodcat_list_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response WmoperFoodQueryFoodcatListResponse
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

func (response *WmoperFoodQueryFoodcatListResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

