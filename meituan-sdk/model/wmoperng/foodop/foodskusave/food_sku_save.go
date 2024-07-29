package foodskusave

/**
* 创建/更新SKU信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const food_sku_save_url = "/wmoper/ng/foodop/food/sku/save"

type Skus struct {
    /**
    * sku唯一标识
    */
    SkuId string `json:"sku_id"`
    /**
    * sku的规格
    */
    Spec string `json:"spec"`
    /**
    * upc码
    */
    Upc string `json:"upc"`
    /**
    * sku的价格，不能为负数，不能超过10个字
    */
    Price string `json:"price"`
    /**
    * sku的库存数量，不能为负数，也不能为小数，传'*'表示表示库存无限
    */
    Stock string `json:"stock"`
    /**
    * 规格排序
    */
    SkuSequence int32 `json:"sku_sequence"`
    /**
    * sku起售时间，要保证不同时间段之间不存在交集
    */
    AvailableTimes AvailableTimes `json:"available_times"`
    /**
    * 料位码
    */
    LocationCode string `json:"locationCode"`
    /**
    * 打包盒数量
    */
    BoxNum string `json:"box_num"`
    /**
    * 打包盒价格，不能为负数
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
    * 填写详细的份量数字。1.填数量(对应weight_unit中 1，2)。2.无需再填数量(对应weight_unit中3)。用法：① 0表示清空 weight，weight_unit。② -1和null 表示不更新weight，weight_unit。③ 其他正常设置weight，weight_unit
    */
    Weight int64 `json:"weight"`
    /**
    * 份量单位。1.克、千克、两、斤、磅、盎司、毫升、升、寸、厘米。2.个、串、枚、粒、 块、只、副、卷、片、贯、碗、杯、袋、瓶、盒、包、锅、罐、扎。3.1人份、2人份、3人份、4人份、5人份、6人份、7人份、8人份、9人份、10人份
    */
    WeightUnit string `json:"weight_unit"`
    /**
    *  售卖属性 spuAttr与skuAttr中的属性数量相同且属性内容相同（mode=2) 该字段为JSON String，具体字段如下： 字段 类型 描述 示例值 skuId string sku唯一标识 10000 no int 属性编号 1 name string 属性名称，长度不能大于10 冷热 value string 属性值, 长度不能大于10 冷 
    */
    SkuAttr string `json:"skuAttr"`
}
type FoodSkuSaveResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type AvailableTimes struct {
    /**
    * 周一
    */
    Monday string `json:"monday"`
    /**
    * 周二
    */
    Tuesday string `json:"tuesday"`
    /**
    * 周三
    */
    Wednesday string `json:"wednesday"`
    /**
    * 周四
    */
    Thursday string `json:"thursday"`
    /**
    * 周五
    */
    Friday string `json:"friday"`
    /**
    * 周六
    */
    Saturday string `json:"saturday"`
    /**
    * 周日
    */
    Sunday string `json:"sunday"`
}
type FoodSkuSaveRequest struct {
    /**
    * 服务商方的菜品id，(三方商品spu_code值， 不同门店可以重复，同一门店内不能重复)，最大长度128
    */
    AppFoodCode string `json:"app_food_code"`
    /**
    * 服务商方的菜品skus，代表菜品下的多个sku信息，使用json格式传递参数。如果skus传递box_num、box_price以skus为准，未传递box_num、box_price以上级参数为准。如果skus未传，会重置菜品的skus信息
    */
    Skus []Skus `json:"skus"`
}



func (req *FoodSkuSaveRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*FoodSkuSaveResponse, error) {
    resp, err := client.InvokeApi(food_sku_save_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response FoodSkuSaveResponse
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

func (response *FoodSkuSaveResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

