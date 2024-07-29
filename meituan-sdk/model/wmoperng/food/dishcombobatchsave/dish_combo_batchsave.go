package dishcombobatchsave

/**
* 批量创建/更新套餐商品（仅支持套餐商品）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const dish_combo_batchsave_url = "/wmoper/ng/food/combo/batchsave"

type DishComboBatchsaveResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type WmComboSkuRel struct {
    /**
    *  套餐下某一单品的第三方SPU code 
    */
    AppFoodCode string `json:"app_food_code"`
    /**
    *  套餐下某一单品的第三方规格code， 注意在创建/修改套餐前，套餐里单品商品必须要已创建 
    */
    SkuId string `json:"sku_id"`
    /**
    *  单品数量，必须&gt;0 
    */
    SkuCount int32 `json:"skuCount"`
    /**
    *  单品显示顺序，默认从1开始 
    */
    Sequence int32 `json:"sequence"`
    /**
    *  加价，保留两位小数 
    */
    AppendPrice float64 `json:"appendPrice"`
    /**
    *  是否默认选择0-否，1-是 ，只能有一个sku默认选择 
    */
    DefaultSelected int32 `json:"defaultSelected"`
    /**
    *  是否必选0-否，1-是，最多只能有一个必选，可以没有 
    */
    RequiredSelected int32 `json:"requiredSelected"`
}
type WmComboSpuPic struct {
    /**
    *  图片URL，URL会在调用上传图片接口后返回 长度限制:255 
    */
    PicLargeUrl string `json:"pic_large_url"`
    /**
    *  图片排序 
    */
    Sequence int32 `json:"sequence"`
}
type ComboSpu struct {
    /**
    *  第三方商品SPUcode 
    */
    AppFoodCode string `json:"app_food_code"`
    /**
    *  商品名称 
    */
    Name string `json:"name"`
    /**
    *  商品排序， 不需要则传0 
    */
    Sequence int32 `json:"sequence"`
    /**
    *  商品售卖时间 如果所有时间都可售卖传- 如果需要分时售卖，支持周一到周日配置。 
    */
    ShippingTime string `json:"shippingTime"`
    /**
    *  商品绑定的视频id，视频需要已经上传至美团 
    */
    WmProductVideoId string `json:"wmProductVideoId"`
    /**
    *  标签列表,用,分隔。招牌=1，单点不送=32 
    */
    Labels string `json:"labels"`
    /**
    *  套餐商品图片，商品的图片，包括头图、详情图，仅支持静态图，动图不支持 
    */
    PicList []WmComboSpuPic `json:"picList"`
    /**
    *  套餐商品属性只有份量。 长度限制 1 
    */
    SpuAttrList []WmComboSpuSalesAttr `json:"spuAttrList"`
    /**
    *  套餐商品的skuList的size必须为1，即只有一个元素 
    */
    SkuList []WmComboSku `json:"skuList"`
    /**
    *  套餐商品所在分组 
    */
    TagName string `json:"tagName"`
    /**
    *  图文详情URL 英文逗号分隔不同图片链接 不同图片没有no排序，C端按照逗号分隔后的顺序展示 
    */
    LongPictures string `json:"longPictures"`
}
type WmComboSku struct {
    /**
    *  第三方规格code 长度限制128 
    */
    SkuId string `json:"sku_id"`
    /**
    *  sku价格，单位元/两位小数，对于套餐商品下的sku，代表其基础价。 限制： 1.基础价需小于等于5000； 2.基础价需小于等于原商品原价之和最低值，最低值=【固定搭配】分组下单品价格*份数+【自由选配】分组下最低价单品n价格之和，n取值为m选n的n； 
    */
    Price float64 `json:"price"`
    /**
    *  售卖状态;0-恢复售卖,1-暂停售卖，默认0 
    */
    SellStatus int32 `json:"sell_status"`
    /**
    *  库存，-1代表不限库存，超卖传0 
    */
    Stock int32 `json:"stock"`
    /**
    *  sku销售属性 
    */
    SkuAttrList []WmComboSkuSalesAttr `json:"skuAttrList"`
    /**
    *  打包盒个数 
    */
    LadderNum int32 `json:"ladder_num"`
    /**
    *  打包费，打包费最小为0 
    */
    LadderPrice float64 `json:"ladder_price"`
    /**
    *  分组可选套餐 的分组信息 套餐为可选分组套餐时，需要传入 当套餐为固定搭配套餐时，不传入 限制分组数量最多为5个 
    */
    ComboGroupList []WmComboGroup `json:"comboGroupList"`
    /**
    *  固定搭配套餐的单品关联信息 固定搭配套餐必须传入 可选分组套餐不传入 
    */
    CombineSkuRelList []WmCombineSkuRel `json:"combineSkuRelList"`
}
type WmComboSkuSalesAttr struct {
    /**
    *  属性名， 套餐商品固定传"份量" 
    */
    Name string `json:"name"`
    /**
    *  属性值， 套餐商品固定传"x人份" 
    */
    Value string `json:"value"`
}
type WmCombineSkuRel struct {
    /**
    *  固定搭配套餐里单品显示顺序，默认从1开始 
    */
    Sequence int32 `json:"sequence"`
    /**
    *  该单品对应的数量,必须&gt;0 固定搭配套餐里的单个单品的skuCount最多为10，固定搭配套餐里的所有单品的skuCount累加不超过20 
    */
    SkuCount int32 `json:"skuCount"`
    /**
    *  套餐下某一个单品的第三方商品SPUcode 
    */
    AppFoodCode string `json:"app_food_code"`
    /**
    *  套餐下某一单品的第三方规格code， 注意在创建/修改套餐前，套餐里单品商品必须要已创建 
    */
    SkuId string `json:"sku_id"`
}
type WmComboGroup struct {
    /**
    *  套餐分组展示顺序，默认从1开始 
    */
    Sequence int32 `json:"sequence"`
    /**
    *  分组名称 长度限制 10 
    */
    GroupName string `json:"groupName"`
    /**
    *  传1，代表非固定分组 传0，代表固定分组 
    */
    GroupType int32 `json:"groupType"`
    /**
    *  套餐分组规则，固定分组不需要传，非固定分组需要传递 
    */
    GroupRule WmComboGroupRule `json:"groupRule"`
    /**
    *  分组下套餐关系信息 
    */
    ComboSkuRelList []WmComboSkuRel `json:"comboSkuRelList"`
}
type DishComboBatchsaveRequest struct {
    /**
    *  该接口只支持套餐全量更新，每次都是全量覆盖。 
    */
    ComboDatas []ComboSpu `json:"comboDatas"`
}
type WmComboGroupRule struct {
    /**
    *  最小选购数量，当套餐分组是M选N时，最小选购数量为N 
    */
    MinOptionalQuantity int32 `json:"minOptionalQuantity"`
    /**
    *  同一单品是否可以多选，0为否，1为是 
    */
    CanBeReSelected bool `json:"canBeReSelected"`
}
type WmComboSpuSalesAttr struct {
    /**
    *  属性项，套餐商品只传“份量” 
    */
    Name string `json:"name"`
    /**
    *  属性值，“份量”对应的value， 套餐商品的value必须为"x人份" 
    */
    Value string `json:"value"`
    /**
    *  属性价格 单位元，支持小数点后两位。套餐价格上限暂定为5000元 
    */
    Price float64 `json:"price"`
    /**
    *  属性值上下架状态;0-上架，1-下架，默认0 
    */
    SellStatus int32 `json:"sell_status"`
}



func (req *DishComboBatchsaveRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DishComboBatchsaveResponse, error) {
    resp, err := client.InvokeApi(dish_combo_batchsave_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DishComboBatchsaveResponse
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

func (response *DishComboBatchsaveResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

