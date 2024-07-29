package wmoperfoodquerylist

/**
* 查询门店菜品列表
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const wmoper_food_query_list_url = "/wmoper/ng/food/queryFoodList"

type WmoperFoodQueryListRequest struct {
    /**
    *  分页查询偏移量 
    */
    Offset int32 `json:"offset"`
    /**
    *  分页查询每页查询的数量,须小于200 
    */
    Limit int32 `json:"limit"`
}
type WmoperFoodQueryListResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []WmoperFoodQueryListData `json:"data"`
    TraceId string `json:"traceId"`
}
type WmoperFoodQueryListData struct {
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
    * App方菜品的skus，该字段为JSON String
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
    * spu属性，该字段为JSON String
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
    /**
    * 美团真实SPUID
    */
    OriginSpuId int64 `json:"origin_spu_id"`
    /**
    * 图片URL,多张图片用逗号隔开
    */
    Pictures string `json:"pictures"`
    HighLight string `json:"high_light"`
    /**
    * 是否设置为招牌商品，最多可以设置15个。1-是，2-否，默认0，不修改该信息
    */
    Speciality int32 `json:"speciality"`
    /**
    * 是否设置为单点不送，1-是，2-否，默认0，不修改该信息
    */
    IsNotSingle int32 `json:"is_not_single"`
    MonthSaled int64 `json:"monthSaled"`
    /**
    * 图文详情URL
    */
    LongPictures string `json:"longPictures"`
}



func (req *WmoperFoodQueryListRequest) DoInvoke(client mtclient.MeituanClient) (*WmoperFoodQueryListResponse, error) {
    resp, err := client.InvokeApi(wmoper_food_query_list_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response WmoperFoodQueryListResponse
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

func (response *WmoperFoodQueryListResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

