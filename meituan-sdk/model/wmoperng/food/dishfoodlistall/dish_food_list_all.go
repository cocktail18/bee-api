package dishfoodlistall

/**
* 查询门店菜品列表（包括套餐商品和普通商品）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const dish_food_list_all_url = "/wmoper/ng/food/dish/food/listAll"

type DishFoodListAllResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []FoodInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type FoodInfo struct {
    /**
    * APP方门店ID(保留字段，目前为空字符串)
    */
    EpoiId string `json:"epoiId"`
    /**
    * APP方菜品id
    */
    AppFoodCode string `json:"app_food_code"`
    /**
    * 是否为仅套餐售卖。true仅在套餐中售卖，false非仅套餐中售卖
    */
    OnlySellInCombo bool `json:"onlySellInCombo"`
    /**
    * 菜品名称
    */
    Name string `json:"name"`
    /**
    * 菜品价格
    */
    Price float64 `json:"price"`
    /**
    * 最小选购数量
    */
    MinOrderCount int32 `json:"min_order_count"`
    /**
    * 最大选购数量
    */
    MaxOrderCount int32 `json:"max_order_count"`
    /**
    * 单位
    */
    Unit string `json:"unit"`
    /**
    * 餐盒数量
    */
    BoxNum float32 `json:"box_num"`
    /**
    * 餐盒价格
    */
    BoxPrice float32 `json:"box_price"`
    /**
    * 1下架，0上架
    */
    IsSoldOut int32 `json:"is_sold_out"`
    /**
    * 菜品图片
    */
    Picture string `json:"picture"`
    /**
    * 菜品创建时间
    */
    Ctime int32 `json:"ctime"`
    /**
    * 菜品修改时间
    */
    Utime int32 `json:"utime"`
    /**
    * 透传给第三方的tag_name
    */
    TagName string `json:"tag_name"`
    /**
    * 菜品图片。多个图片间用逗号分隔
    */
    Pictures string `json:"pictures"`
    /**
    * 是否设置为招牌商品。最多可以设置15个。1是，2否，默认0未设置
    */
    Speciality int32 `json:"speciality"`
    /**
    * 是否设置为单点不送。1是，2否，默认0未设置
    */
    IsNotSingle int32 `json:"is_not_single"`
    /**
    * 菜品描述
    */
    Description string `json:"description"`
    /**
    * 菜品分类名
    */
    CategoryName string `json:"category_name"`
    /**
    * 当前分类下的排序序号
    */
    Sequence int32 `json:"sequence"`
    /**
    * 透传给第三方的偏移后的spu的id
    */
    MtSpuId int64 `json:"mt_spu_id"`
    /**
    * 透传给第三方的偏移后的tag_id
    */
    MtTagId int64 `json:"mt_tag_id"`
    /**
    * 美团真实SPUID
    */
    OriginSpuId int64 `json:"origin_spu_id"`
    HighLight string `json:"high_light"`
    MonthSaled int64 `json:"monthSaled"`
    /**
    * spu售卖属性。spuAttr与skuAttr中的属性数量相同且属性内容相同（mode=2)
    */
    SpuAttr string `json:"spuAttr"`
    /**
    * 三方菜品的skus
    */
    Skus string `json:"skus"`
    /**
    * 图文详情URL
    */
    LongPictures string `json:"longPictures"`
}
type DishFoodListAllRequest struct {
    /**
    *  分页查询偏移量 
    */
    Offset int32 `json:"offset"`
    /**
    *  分页查询每页查询的数量,须小于200 
    */
    Limit int32 `json:"limit"`
}



func (req *DishFoodListAllRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DishFoodListAllResponse, error) {
    resp, err := client.InvokeApi(dish_food_list_all_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DishFoodListAllResponse
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

func (response *DishFoodListAllResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

