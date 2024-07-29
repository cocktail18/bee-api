package create

/**
* 创建或更新商品（必接）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const create_url = "/foodmop/sku/create"

type SpuBasicDTO struct {
    /**
    *  SPU 编码 
    */
    VendorSpuId string `json:"vendorSpuId"`
    /**
    *  商品名称 
    */
    Name string `json:"name"`
    /**
    *  商品售卖时间，若不设值，商品全时段均可售卖 
    */
    SaleTime []TimeDTO `json:"saleTime"`
    /**
    *  商品描述 
    */
    Description string `json:"description"`
    /**
    *  商品主要原料 
    */
    Material string `json:"material"`
    /**
    *  商品简述- 最多 27 个字 
    */
    Brief string `json:"brief"`
    /**
    *  全国上下架状态 0 ：下架 1：上架 默认填 1 
    */
    ShelfStatus int32 `json:"shelfStatus"`
    /**
    *  是否可用会员积分抵扣 true：可以 false：不可 默认 false 
    */
    UseMemberPoint bool `json:"useMemberPoint"`
    /**
    *  商品标签，目前只支持新品标签，固定值为 "NEW" 
    */
    TagList []string `json:"tagList"`
    /**
    *  媒介信息 
    */
    Media MediaDTO `json:"media"`
    /**
    *  库存状态 1：有限库存 2：无限库存 一般推荐，饮品属于无限库存，食品有限库存。设置为无限库存的商品，才可以进行商品库存同步，无限库存的商品不用操作库存同步 
    */
    StockStatus int32 `json:"stockStatus"`
    /**
    *  扩展属性，一般品牌无特殊需求，不填即可 Map&lt;String,String&gt; Key Value 说明 HIDE_IN_SHOP_MENU true false 菜单隐藏商品 SALE_ATTRIBUTE_MUTEX_MAP Map&lt;String, List&lt;String&gt;&gt; 使用JSON.toJSONString序列化 售卖属性互斥，互斥的售卖属性不能同时选中。例如「中杯」和「热」互斥，那么用户选中「中杯」后，「热」会置灰不可选，反之亦然 key: 售卖属性 code value: 与 key 互斥的售卖属性 code 列表 
    */
    ExtendInfoMap Map `json:"extendInfoMap"`
    /**
    *  数据来源 1：品牌自建 2：品牌下发 3：门店自建 默认为1 
    */
    Source int32 `json:"source"`
    /**
    *  source == 2 或 source ==3时必填 
    */
    VendorShopId string `json:"vendorShopId"`
    /**
    *  source == 2 时，必填 值为模板模板售卖属性 id 
    */
    TradeMarkCode string `json:"tradeMarkCode"`
}
type CreateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}
type CreateRequest struct {
    /**
    *  商品信息 不为空，每次最多传 1 个 
    */
    VendorSpuList []VendorSpuDTO `json:"vendorSpuList"`
}
type SpuPremiumGroupBasicDTO struct {
    /**
    *  配料属性组编码 
    */
    GroupCode string `json:"groupCode"`
    /**
    *  配料属性组名称 
    */
    GroupName string `json:"groupName"`
    /**
    *  配料属性组排序 
    */
    Rank int32 `json:"rank"`
    /**
    *  配料组是否支持选多个配料 true：支持 false：不支持 
    */
    ChooseMultiPremium bool `json:"chooseMultiPremium"`
    /**
    *  配料组下配料是否支持选多份 true：支持 false：不支持 例如 chooseMultiPremium=true &amp; multiChoosePremium=true 时，配料组可以选多个配料，且每个配料可以选多份（每个配料最多能选几个，在 创建或更改商品客制化规则 中有明确描述） 
    */
    MultiChoosePremium bool `json:"multiChoosePremium"`
}
type TimeRangeDTO struct {
    /**
    *  开始时间 不能为空，只能为 xx:xx 格式 
    */
    Start string `json:"start"`
    /**
    *  结束时间 不能为空，只能为 xx:xx 格式 
    */
    End string `json:"end"`
}
type PictureDTO struct {
    /**
    *  头像 比例需为 1:1，尺寸不小于168×168，0.5M以内 主体物建议不超过红色区域，且位置在画面中心，边缘要留白 
    */
    HeadPictureList []string `json:"headPictureList"`
    /**
    *  描述图片 比例建议为 16:9，尺寸不小于750×416，0.5M以内 主体物建议不超过红色区域，且位置在画面中心，上方一定要留白，否则会被Iphone刘海挡住 
    */
    DescriptionPictureList []string `json:"descriptionPictureList"`
}
type BackgroundCategoryBasicDTO struct {
    /**
    *  后台类目 id 
    */
    VendorBackgroundCategoryId string `json:"vendorBackgroundCategoryId"`
    /**
    *  后台类目名称 
    */
    VendorBackgroundCategoryName string `json:"vendorBackgroundCategoryName"`
    /**
    *  后台类目英文名称 
    */
    VendorBackgroundCategoryEnglishName string `json:"vendorBackgroundCategoryEnglishName"`
}
type VendorSpuSaleAttributeGroupDTO struct {
    /**
    *  商品售卖属性基本信息 
    */
    SpuSaleAttributeGroupBasic SpuSaleAttributeGroupBasicDTO `json:"spuSaleAttributeGroupBasic"`
    /**
    *  售卖属性组关联的属性值 Map&lt;Integer,String&gt;类型 key：售卖属性于组中的展示顺序 value：售卖属性 code 不为空，属性值 code 必须对应于售卖属性编码，code 不能重复 
    */
    VendorSpuSaleAttributeCodeMap Map `json:"vendorSpuSaleAttributeCodeMap"`
    /**
    *  星巴克品牌定制，其他品牌不关注 售卖属性组关联的附加冷热配料 Map&lt;Integer,String&gt;类型 key：属性值于组中的展示顺序 value：附加冷热配料 code 
    */
    VendorSpuPremiumCodeMap Map `json:"vendorSpuPremiumCodeMap"`
}
type SkuSaleAttributeBasicDTO struct {
    /**
    *  售卖属性组名称 
    */
    AttributeGroupName string `json:"attributeGroupName"`
    /**
    *  售卖属性组编码 
    */
    AttributeGroupCode string `json:"attributeGroupCode"`
    /**
    *  售卖属性名称 
    */
    AttributeName string `json:"attributeName"`
    /**
    *  售卖属性编码 
    */
    AttributeCode string `json:"attributeCode"`
}
type VendorSpuDTO struct {
    /**
    *  商品基本信息 
    */
    SpuBasic SpuBasicDTO `json:"spuBasic"`
    /**
    *  商品类型 NORMAL_SPU(1, "普通商品") FIXED_COMBO(2, "固定套餐") DINNERWARE(3, "餐具") 
    */
    Type int32 `json:"type"`
    /**
    *  商品 SKU 列表 
    */
    VendorSkuList []VendorSkuDTO `json:"vendorSkuList"`
    /**
    *  默认选中的 skuId skuId 必须存在于 vendorSkuList 中 
    */
    DefaultVendorSkuId string `json:"defaultVendorSkuId"`
    /**
    *  商品后台类目基本信息，用于开店宝后台配置营销活动时，按商品后台分类进行商品圈选。后台类目品牌自定义即可 
    */
    BackgroundCategoryBasic BackgroundCategoryBasicDTO `json:"backgroundCategoryBasic"`
    /**
    *  商品售卖属性组列表 当 SKU 数 &gt; 1 时必填；当 SKU 数 = 1 时，可不填 售卖属性组中关联的售卖属性必须通过创建或更新售卖属性前置同步至美团 
    */
    VendorSpuSaleAttributeGroupList []VendorSpuSaleAttributeGroupDTO `json:"vendorSpuSaleAttributeGroupList"`
    /**
    *  商品配料属性组列表 若商品没有配料，则不填 配料组中关联的售卖属性必须通过创建或更新配料属性前置同步至美团 
    */
    VendorSpuPremiumGroupList []VendorSpuPremiumGroupDTO `json:"vendorSpuPremiumGroupList"`
}
type MediaDTO struct {
    Picture PictureDTO `json:"picture"`
}
type ResultData struct {
    /**
    * key：vendorSpuId value：同步结果，成功返回 “成功”，失败则返回相应的异常信息
    */
    Map Map `json:"map"`
}
type VendorSpuPremiumGroupDTO struct {
    /**
    *  配料组基本信息 
    */
    SpuPremiumGroupBasic SpuPremiumGroupBasicDTO `json:"spuPremiumGroupBasic"`
    /**
    *  配料组关联的配料 Map&lt;Integer,String&gt; key：配料于组中的展示顺序 value：配料 code 不为空，属性值 code 必须对应于配料属性编码，code 不能重复 
    */
    VendorSpuPremiumCodeMap Map `json:"vendorSpuPremiumCodeMap"`
    /**
    *  若配料组之间具有父子层级，则需要配置配料组对应的父配料 code 例如星巴克品牌中，配料「低因」下挂子配料性组【萃取方式】，那么【萃取方式】组的 parentPremiumCode 为「低因」的属性值编码 一般品牌不需要关注此字段，不填即可 属性值 code 必须对应于配料属性编码 
    */
    ParentPremiumCode string `json:"parentPremiumCode"`
}
type SkuBasicDTO struct {
    /**
    *  商品 SKU 编码 
    */
    VendorSkuId string `json:"vendorSkuId"`
    /**
    *  SKU 价格，单位分 
    */
    Price int64 `json:"price"`
    /**
    *  会员价，单位分 
    */
    MemberPrice int64 `json:"memberPrice"`
    /**
    *  全国上下架状态 0 ：下架 1：上架 默认填 1 即可 
    */
    ShelfStatus int32 `json:"shelfStatus"`
    /**
    *  扩展属性 Map&lt;String,String&gt; Key 说明 DEFAULT_CPU_SIZE_ML 默认杯型毫升数，星巴克品牌针对部分 SKU 需要展示默认杯型的毫升数。 
    */
    ExtendInfoMap Map `json:"extendInfoMap"`
    /**
    *  数据来源 1：品牌自建 2：品牌下发 3：门店自建 默认1，其他先联系美团PM 
    */
    Source int32 `json:"source"`
    /**
    *  source == 2 || source ==3时必填 
    */
    VendorShopId string `json:"vendorShopId"`
    /**
    *  source == 2 时必填，值为模版商品 vendorSpuId 
    */
    TradeMarkCode string `json:"tradeMarkCode"`
}
type TimeDTO struct {
    /**
    *  一周的第几天，eg： 1：星期一 2：星期二 ...以此类推 不为空，只能为 1~7 
    */
    DayOfWeek int32 `json:"dayOfWeek"`
    /**
    *  时间段 
    */
    Range []TimeRangeDTO `json:"range"`
}
type SpuSaleAttributeGroupBasicDTO struct {
    /**
    *  属性组编码 
    */
    GroupCode string `json:"groupCode"`
    /**
    *  属性组名称 
    */
    GroupName string `json:"groupName"`
    /**
    *  属性组排序 
    */
    Rank int32 `json:"rank"`
}
type VendorSkuDTO struct {
    /**
    *  商品 SKU 基本信息 
    */
    SkuBasic SkuBasicDTO `json:"skuBasic"`
    /**
    *  组成该 SKU 的售卖属性列表 
    */
    SkuSaleAttributeBasicList []SkuSaleAttributeBasicDTO `json:"skuSaleAttributeBasicList"`
}
type Map struct {
}



func (req *CreateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CreateResponse, error) {
    resp, err := client.InvokeApi(create_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CreateResponse
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

func (response *CreateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

