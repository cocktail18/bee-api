package dishgetdetail

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

const dish_get_detail_url = "/waimai/ng/dish/getDetail"

type DishGetDetailData struct {
    /**
    * 服务商门店id
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
    /**
    * 当前分类下的排序序号
    */
    Sequence int32 `json:"sequence"`
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
    * spu售卖属性，spuAttr与skuAttr中的属性数量相同且属性内容相同（mode=2)
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
type DishGetDetailResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data DishGetDetailData `json:"data"`
    TraceId string `json:"traceId"`
}
type DishGetDetailRequest struct {
    /**
    *  菜品id(三方商品spu_code值, 不同门店可以重复，同一门店内不能重复)最大长度128 
    */
    EDishCode string `json:"eDishCode"`
    /**
    *  标识订单是否为团餐定制菜单。 1：团餐定制菜单，0或空：外卖菜单 
    */
    OrderEntranceType int32 `json:"orderEntranceType"`
}



func (req *DishGetDetailRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DishGetDetailResponse, error) {
    resp, err := client.InvokeApi(dish_get_detail_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DishGetDetailResponse
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

func (response *DishGetDetailResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

