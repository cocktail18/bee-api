package wmoperfoodquerydetail

/**
* 查询菜品详情
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const wmoper_food_query_detail_url = "/wmoper/ng/food/detail"

type WmoperFoodQueryDetailData struct {
    EpoiId string `json:"epoiId"`
    /**
    * APP方菜品id
    */
    AppFoodCode string `json:"app_food_code"`
    Name string `json:"name"`
    Description string `json:"description"`
    /**
    * sku价格
    */
    Price float32 `json:"price"`
    /**
    * 最小购买量，最小为1
    */
    MinOrderCount int32 `json:"min_order_count"`
    /**
    * 单位
    */
    Unit string `json:"unit"`
    /**
    * sku餐盒数量
    */
    BoxNum float32 `json:"box_num"`
    /**
    * sku餐盒价格
    */
    BoxPrice float32 `json:"box_price"`
    /**
    * 菜品分类
    */
    CategoryName string `json:"category_name"`
    /**
    * 菜品上下架状态，0表上架，1表下架
    */
    IsSoldOut int32 `json:"is_sold_out"`
    /**
    * 菜品图片id
    */
    Picture string `json:"picture"`
    Sequence int64 `json:"sequence"`
    /**
    * 代表菜品下的多个sku信息，使用json格式传递参数
    */
    Skus string `json:"skus"`
    /**
    * 创建时间（当前距离Epoch（1970年1月1日） 以秒计算的时间，即unix-timestamp）
    */
    Ctime int32 `json:"ctime"`
    /**
    * 更新时间（当前距离Epoch（1970年1月1日） 以秒计算的时间，即unix-timestamp）
    */
    Utime int32 `json:"utime"`
    /**
    * spu属性，该字段为JSON String
    */
    SpuAttr string `json:"spuAttr"`
    MtSpuId int64 `json:"mt_spu_id"`
    MtTagId int64 `json:"mt_tag_id"`
    TagName string `json:"tag_name"`
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
    /**
    * 月销量
    */
    MonthSaled int32 `json:"monthSaled"`
    /**
    * 图文详情URL
    */
    LongPictures string `json:"longPictures"`
}
type WmoperFoodQueryDetailResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data WmoperFoodQueryDetailData `json:"data"`
    TraceId string `json:"traceId"`
}
type WmoperFoodQueryDetailRequest struct {
    /**
    *  菜品id 
    */
    AppFoodCode string `json:"app_food_code"`
}



func (req *WmoperFoodQueryDetailRequest) DoInvoke(client mtclient.MeituanClient) (*WmoperFoodQueryDetailResponse, error) {
    resp, err := client.InvokeApi(wmoper_food_query_detail_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response WmoperFoodQueryDetailResponse
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

func (response *WmoperFoodQueryDetailResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

