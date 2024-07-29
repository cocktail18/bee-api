package foodbatchbulksave

/**
* 批量创建或更新菜品
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const food_batch_bulk_save_url = "/wmoper/ng/foodop/food/batchbulksave"

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
type FoodData struct {
    /**
    * 服务商方的菜品id，(三方商品spu_code值， 不同门店可以重复，同一门店内不能重复)，最大长度128
    */
    AppFoodCode string `json:"app_food_code"`
    /**
    * 菜品名称，最多不超过30个字符
    */
    Name string `json:"name"`
    /**
    * 菜品描述，最多不超过255个字符
    */
    Description string `json:"description"`
    /**
    * 服务商方的菜品skus，代表菜品下的多个sku信息，使用json格式传递参数。如果skus传递box_num、box_price以skus为准，未传递box_num、box_price以上级参数为准。如果skus未传，会重置菜品的skus信息
    */
    Skus []Skus `json:"skus"`
    /**
    * 菜品价格，不能为负数
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
    * 打包盒数量
    */
    BoxNum float32 `json:"box_num"`
    /**
    * 打包盒价格，不能为负数
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
    * 菜品图片id（注意：①此处的图片id由步骤1上传图片获得②更新菜品图片时必须更新图片URL），只支持jpg格式，图片需要小于1600*1200
    */
    Picture string `json:"picture"`
    /**
    * 当前分类下的排序序号
    */
    Sequence int32 `json:"sequence"`
    /**
    *  true：该单品仅在套餐里售卖，不单独售卖 false：默认值，在套餐和单品都可售 
    */
    OnlySellInCombo bool `json:"onlySellInCombo"`
    /**
    *  图文详情URL 英文逗号分隔不同图片链接 不同图片没有no排序，C端按照逗号分隔后的顺序展示 
    */
    LongPictures string `json:"longPictures"`
}
type FoodBatchBulkSaveResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type FoodBatchBulkSaveRequest struct {
    /**
    *  菜品集合的 json 数据， 处理逻辑： 1）如果菜品原来不存在就新增； 2）如app_food_code存在就更新； 3）picture 可以是调用 上传菜品图片接口 后返回的图片ID, 也可以是url地址,推荐使用图片ID的形式。 如果skus传递box_num、box_price以skus为准, 未传递box_num、box_price以上级参数为准 
    */
    FoodData []FoodData `json:"food_data"`
}



func (req *FoodBatchBulkSaveRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*FoodBatchBulkSaveResponse, error) {
    resp, err := client.InvokeApi(food_batch_bulk_save_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response FoodBatchBulkSaveResponse
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

func (response *FoodBatchBulkSaveResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

