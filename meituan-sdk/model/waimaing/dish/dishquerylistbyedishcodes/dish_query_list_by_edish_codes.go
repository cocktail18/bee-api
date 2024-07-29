package dishquerylistbyedishcodes

/**
* 根据eDishCode批量查询外卖菜品信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const dish_query_list_by_edish_codes_url = "/waimai/dish/queryListByEdishCodes"

type Skus struct {
    /**
    * 配送时间， 默认门店营业时间
    */
    AvailableTimes AvailableTimes `json:"availableTimes"`
    BoxNum float64 `json:"boxNum"`
    BoxPrice float64 `json:"boxPrice"`
    /**
    * sku的料位码
    */
    LocationCode string `json:"locationCode"`
    /**
    * 价格
    */
    Price string `json:"price"`
    SkuAttr string `json:"skuAttr"`
    /**
    * ERP sku ID
    */
    SkuId string `json:"skuId"`
    SkuSequence int64 `json:"sku_sequence"`
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
    Weight int64 `json:"weight"`
    WeightUnit string `json:"weightUnit"`
    LadderNum string `json:"ladderNum"`
    LadderPrice string `json:"ladderPrice"`
}
type DishQueryListByEdishCodesRequest struct {
    /**
    *  ERP方要查询菜品id（多个，以逗号隔开，最多100个） 
    */
    EDishCodes string `json:"eDishCodes"`
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
type DishQueryListByEdishCodesResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data DishQueryListByEdishCodesData `json:"data"`
    TraceId string `json:"traceId"`
}
type List struct {
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
    DishId int64 `json:"dishId"`
    /**
    * 菜名
    */
    DishName string `json:"dishName"`
    /**
    * ERP方菜品Id
    */
    EDishCode string `json:"eDishCode"`
    EpoiId string `json:"epoiId"`
    FoodId int64 `json:"foodId"`
    /**
    * 菜品上下架状态，0表上架，1表下架
    */
    IsSoldOut int32 `json:"isSoldOut"`
    MaxOrderCount int64 `json:"maxOrderCount"`
    /**
    * 最小购买数量
    */
    MinOrderCount int32 `json:"minOrderCount"`
    Operation int64 `json:"operation"`
    /**
    * 图片id或地址
    */
    Picture string `json:"picture"`
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
    SpecialPrice float64 `json:"specialPrice"`
    /**
    * 单位/规格
    */
    Unit string `json:"unit"`
}
type DishQueryListByEdishCodesData struct {
    Msg string `json:"msg"`
    /**
    * 数据
    */
    List []List `json:"list"`
    /**
    * 查询结果标识，0全部查询成功、1全部查询失败、2部分查询成功
    */
    Status int32 `json:"status"`
}



func (req *DishQueryListByEdishCodesRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DishQueryListByEdishCodesResponse, error) {
    resp, err := client.InvokeApi(dish_query_list_by_edish_codes_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DishQueryListByEdishCodesResponse
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

func (response *DishQueryListByEdishCodesResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

