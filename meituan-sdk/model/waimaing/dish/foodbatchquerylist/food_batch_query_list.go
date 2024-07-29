package foodbatchquerylist

/**
* 批量查询外卖菜品
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const food_batch_query_list_url = "/waimai/ng/dish/batchQuery"

type FoodBatchQueryListRequest struct {
    Biz string `json:"biz"`
}
type FoodBatchQueryListResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []DishInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type DishInfo struct {
    /**
    * APP方门店ID(保留字段，目前为空字符串)
    */
    EpoiId string `json:"epoiId"`
    /**
    * APP方菜品id
    */
    AppFoodCode string `json:"app_food_code"`
    /**
    * 菜品名称
    */
    Name string `json:"name"`
    /**
    * 菜品描述
    */
    Description string `json:"description"`
    Price float64 `json:"price"`
    /**
    * 最小购买量
    */
    MinOrderCount int32 `json:"min_order_count"`
    /**
    * 单位
    */
    Unit string `json:"unit"`
    /**
    * 打包盒数量
    */
    BoxNum float32 `json:"box_num"`
    /**
    * 餐盒价格
    */
    BoxPrice float32 `json:"box_price"`
    /**
    * 菜品分类名
    */
    CategoryName string `json:"category_name"`
    /**
    * 1：下架，0：上架
    */
    IsSoldOut int32 `json:"is_sold_out"`
    /**
    * 菜品图片
    */
    Picture string `json:"picture"`
    /**
    * 当前分类下的排序序号
    */
    Sequence int32 `json:"sequence"`
    /**
    * App方菜品的skus
    */
    Skus string `json:"skus"`
    /**
    * 菜品创建时间
    */
    Ctime int32 `json:"ctime"`
    /**
    * 菜品修改时间
    */
    Utime int32 `json:"utime"`
    /**
    * 菜品属性
    */
    SpuAttr string `json:"spuAttr"`
    /**
    * 透传给第三方的偏移后的spu的id
    */
    MtSpuId int64 `json:"mt_spu_id"`
    /**
    * 透传给第三方的偏移后的tag_id
    */
    MtTagId int64 `json:"mt_tag_id"`
    /**
    * 透传给第三方的tag_name
    */
    TagName string `json:"tag_name"`
    OriginSpuId int64 `json:"origin_spu_id"`
    Pictures string `json:"pictures"`
    HighLight string `json:"high_light"`
    Speciality int64 `json:"speciality"`
    IsNotSingle int64 `json:"is_not_single"`
    MonthSaled int64 `json:"monthSaled"`
    /**
    * 图文详情URL
    */
    LongPictures string `json:"longPictures"`
}



func (req *FoodBatchQueryListRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*FoodBatchQueryListResponse, error) {
    resp, err := client.InvokeApi(food_batch_query_list_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response FoodBatchQueryListResponse
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

func (response *FoodBatchQueryListResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

