package dishquerylistbyepoiid

/**
* 根据ERP的门店id查询门店下的菜品【不包含美团的菜品Id】
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const dish_query_list_by_epoiid_url = "/waimai/dish/queryListByEPoiId"

type Skus struct {
    /**
    * 配送时间， 默认门店营业时间
    */
    AvailableTimes AvailableTimes `json:"availableTimes"`
    /**
    * 打包盒数量
    */
    BoxNum float32 `json:"boxNum"`
    /**
    * 打包盒价格
    */
    BoxPrice float32 `json:"boxPrice"`
    /**
    * sku的料位码
    */
    LocationCode string `json:"locationCode"`
    /**
    * 价格
    */
    Price string `json:"price"`
    /**
    * 售卖属性，spuAttr与skuAttr中的属性数量相同且属性内容相同（mode\u003d2)
    */
    SkuAttr string `json:"skuAttr"`
    /**
    * ERP sku ID
    */
    SkuId string `json:"skuId"`
    /**
    * 规格排序
    */
    SkuSequence int32 `json:"sku_sequence"`
    /**
    * 规格
    */
    Spec string `json:"spec"`
    /**
    * 库存
    */
    Stock string `json:"stock"`
    /**
    * 菜品upc码
    */
    Upc string `json:"upc"`
    /**
    * 重量或份量数值
    */
    Weight int64 `json:"weight"`
    WeightUnit string `json:"weightUnit"`
    /**
    * 每M个商品需N元包装费，本字段代表M个商品。
    */
    LadderNum string `json:"ladderNum"`
    /**
    * 每M个商品需N元包装费，本字段代表N元。
    */
    LadderPrice string `json:"ladderPrice"`
}
type AvailableTimes struct {
    Friday string `json:"friday"`
    Monday string `json:"monday"`
    Saturday string `json:"saturday"`
    Sunday string `json:"sunday"`
    Thursday string `json:"thursday"`
    Tuesday string `json:"tuesday"`
    Wednesday string `json:"wednesday"`
}
type DishQueryListByEpoiidResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []DishInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type DishQueryListByEpoiidRequest struct {
    /**
    *  起始条目数 
    */
    Offset int32 `json:"offset"`
    /**
    *  每页大小，须小于200 
    */
    Limit int32 `json:"limit"`
}
type DishInfo struct {
    /**
    * 餐盒数量
    */
    BoxNum float32 `json:"boxNum"`
    /**
    * 餐盒单价
    */
    BoxPrice float32 `json:"boxPrice"`
    /**
    * 菜品分类
    */
    CategoryName string `json:"categoryName"`
    /**
    * 菜品描述
    */
    Description string `json:"description"`
    /**
    * 菜名
    */
    DishName string `json:"dishName"`
    /**
    * ERP方菜品Id
    */
    EDishCode string `json:"eDishCode"`
    /**
    * 服务商门店id
    */
    EpoiId string `json:"epoiId"`
    /**
    * 是否设置为单点不送，1-是，2-否，默认0，不修改该信息
    */
    IsNotSingle int32 `json:"isNotSingle"`
    /**
    * 菜品上下架状态。0-上架，1-下架
    */
    IsSoldOut int32 `json:"isSoldOut"`
    /**
    * 最小购买数量
    */
    MinOrderCount int32 `json:"minOrderCount"`
    /**
    * 图片id或地址
    */
    Picture string `json:"picture"`
    /**
    * 图片URL,多张图片用逗号隔开
    */
    Pictures string `json:"pictures"`
    /**
    * 价格
    */
    Price float32 `json:"price"`
    /**
    * 当前分类下的排序序号
    */
    Sequence int32 `json:"sequence"`
    /**
    * ERP方菜品的skus，代表菜品下的多个sku信息
    */
    Skus []Skus `json:"skus"`
    /**
    * 是否设置为招牌商品，最多可以设置15个。1-是，2-否，默认0，不修改该信息
    */
    Speciality int32 `json:"speciality"`
    /**
    * spu售卖属性，spuAttr与skuAttr中的属性数量相同且属性内容相同（mode\u003d2)
    */
    SpuAttr string `json:"spuAttr"`
    /**
    * 单位/规格
    */
    Unit string `json:"unit"`
    /**
    * 图文详情URL 
    */
    LongPictures string `json:"longPictures"`
}



func (req *DishQueryListByEpoiidRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DishQueryListByEpoiidResponse, error) {
    resp, err := client.InvokeApi(dish_query_list_by_epoiid_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DishQueryListByEpoiidResponse
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

func (response *DishQueryListByEpoiidResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

