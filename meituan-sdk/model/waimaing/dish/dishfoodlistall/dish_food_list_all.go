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

const dish_food_list_all_url = "/waimai/ng/dish/food/listAll"

type Skus struct {
    /**
    * 菜品可售时间
    */
    AvailableTimes AvailableTimes `json:"available_times"`
    /**
    * 打包盒数量
    */
    BoxNum string `json:"box_num"`
    /**
    * 餐盒价格
    */
    BoxPrice string `json:"box_price"`
    /**
    * 每M个商品需N元包装费，本字段代表M个商品。不可为0。如传“ladder_num”和“ladder_price”，则代表商家使用了阶梯价格，“box_num”和“box_price”中的内容不生效
    */
    LadderNum string `json:"ladder_num"`
    /**
    * 每M个商品需N元包装费，本字段代表N元。如传“ladder_num”和“ladder_price”，则代表商家使用了阶梯价格，“box_num”和“box_price”中的内容不生效
    */
    LadderPrice string `json:"ladder_price"`
    /**
    * 位置码
    */
    LocationCode string `json:"location_code"`
    /**
    * sku价格
    */
    Price string `json:"price"`
    /**
    * sku唯一标识
    */
    SkuId string `json:"sku_id"`
    /**
    * sku的规格
    */
    Spec string `json:"spec"`
    /**
    * sku库存数量，不能为负数或小数，传'*'表示库存无限
    */
    Stock string `json:"stock"`
    /**
    * 规格排序
    */
    SkuSequence int32 `json:"sku_sequence"`
    Weight int64 `json:"weight"`
    WeightUnit string `json:"weight_unit"`
    /**
    * 透传给第三方的偏移后的sku的id
    */
    MtSkuId int64 `json:"mt_sku_id"`
    /**
    * 美团真实SKUID
    */
    OriginSkuId int64 `json:"origin_sku_id"`
    /**
    * 售卖属性。spuAttr与skuAttr中的属性数量相同且属性内容相同（mode=2)
    */
    SkuAttr string `json:"skuAttr"`
    /**
    * upc码
    */
    Upc string `json:"upc"`
    /**
    * 分组可选套餐的分组信息。当套餐为可选分组套餐时，需要传入。当套餐为固定搭配套餐时，不传入。固定搭配套餐分组数量不能超过5个
    */
    ComboGroupList []WmComboGroupVo `json:"comboGroupList"`
    /**
    * 固定搭配套餐的单品关联信息。固定搭配套餐必须传入，可选分组套餐不传入
    */
    CombineSkuRelList []WmCombineSkuRelVo `json:"combineSkuRelList"`
}
type WmComboGroupVo struct {
    /**
    * 套餐分组展示顺序，默认从1开始
    */
    Sequence int32 `json:"sequence"`
    /**
    * 分组名称
    */
    GroupName string `json:"groupName"`
    /**
    * 传1，代表非固定分组。传0，代表固定分组
    */
    GroupType int32 `json:"groupType"`
    /**
    * 分组规则。固定分组不需要传分组规则，非固定分组需要传递分组规则
    */
    GroupRule WmComboGroupRuleVo `json:"groupRule"`
    /**
    * 分组下套餐与单品关系信息
    */
    ComboSkuRelList []WmComboSkuRelVo `json:"comboSkuRelList"`
}
type DishFoodListAllResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []FoodInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type WmCombineSkuRelVo struct {
    /**
    * 固定搭配套餐里单品显示顺序，从1开始
    */
    Sequence int32 `json:"sequence"`
    /**
    * 该单品对应的数量,必须>0。固定搭配套餐里的单个单品的skuCount最多为10，固定搭配套餐里的所有单品的skuCount累加不超过20。
    */
    SkuCount int32 `json:"skuCount"`
    /**
    * 套餐下某一个单品的第三方商品SPUcode
    */
    AppFoodCode string `json:"app_food_code"`
    /**
    * 套餐下某一单品的第三方规格code，注意在创建/修改套餐前，套餐里单品商品必须要已创建
    */
    SkuId string `json:"sku_id"`
}
type AvailableTimes struct {
    /**
    * 周五
    */
    Friday string `json:"friday"`
    /**
    * 周一
    */
    Monday string `json:"monday"`
    /**
    * 周六
    */
    Saturday string `json:"saturday"`
    /**
    * 周日
    */
    Sunday string `json:"sunday"`
    /**
    * 周四
    */
    Thursday string `json:"thursday"`
    /**
    * 周二
    */
    Tuesday string `json:"tuesday"`
    /**
    * 周三
    */
    Wednesday string `json:"wednesday"`
}
type WmComboGroupRuleVo struct {
    /**
    * 最小选购数量，当套餐分组是M选N时，最小选购数量为N
    */
    MinOptionalQuantity int32 `json:"minOptionalQuantity"`
    /**
    * 同一单品是否可以多选，0为否，1为是
    */
    CanBeReelected bool `json:"canBeReelected"`
}
type WmComboSkuRelVo struct {
    /**
    * 套餐下某一单品的第三方SPU code，注意在创建/修改套餐前，套餐里单品商品必须要已创建
    */
    AppFoodCode string `json:"app_food_code"`
    /**
    * 套餐下某一单品的第三方规格code，注意在创建/修改套餐前，套餐里单品商品必须要已创建
    */
    SkuId string `json:"sku_id"`
    /**
    * 单品数量，必须>0。可选套餐商品下单个单品的skuCount最多为10
    */
    SkuCount int32 `json:"skuCount"`
    /**
    * 单品显示顺序，默认从1开始
    */
    Sequence int32 `json:"sequence"`
    /**
    * 加价。保留两位小数
    */
    AppendPrice float64 `json:"appendPrice"`
    /**
    * 是否默认选择。0-否，1-是，只能有一个默认选择
    */
    DefaultSelected int32 `json:"defaultSelected"`
    /**
    * 是否必选。0-否，1-是，最多只能有一个必选
    */
    RequiredSelected int32 `json:"requiredSelected"`
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
    Skus []Skus `json:"skus"`
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
    resp, err := client.InvokeApi(dish_food_list_all_url, 2, appAuthToken, req)

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

