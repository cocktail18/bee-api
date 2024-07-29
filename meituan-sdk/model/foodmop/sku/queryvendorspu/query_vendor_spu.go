package queryvendorspu

/**
* 查询门店商品信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_vendor_spu_url = "/foodmop/sku/queryVendorSpu"

type SpuBasicDTO struct {
    /**
    * SPU 编码
    */
    VendorSpuId string `json:"vendorSpuId"`
    /**
    * 商品名称
    */
    Name string `json:"name"`
    /**
    * 商品售卖时间，若不设值，商品全时段均可售卖
    */
    SaleTime []TimeDTO `json:"saleTime"`
    /**
    * 类目描述
    */
    Description string `json:"description"`
    /**
    * 商品主要原料
    */
    Material string `json:"material"`
    /**
    * 商品简述
    */
    Brief string `json:"brief"`
    /**
    * 全国上下架状态 0 ：下架 1：上架 默认填 1
    */
    ShelfStatus int32 `json:"shelfStatus"`
    /**
    * 是否可用会员积分抵扣 true：可以 false：不可 默认 false
    */
    UseMemberPoint bool `json:"useMemberPoint"`
    /**
    * 商品标签，目前只支持新品标签，固定值为 "NEW"
    */
    TagList []string `json:"tagList"`
    /**
    * 媒介信息
    */
    Media MediaDTO `json:"media"`
    /**
    * 库存状态 1：有限库存 2：无限库存 一般推荐，饮品属于无限库存，食品有限库存。设置为无限库存的商品，才可以进行商品库存同步，无限库存的商品不用操作库存同步
    */
    StockStatus int32 `json:"stockStatus"`
    /**
    * 扩展属性
    */
    ExtendInfoMap ExtendInfo `json:"extendInfoMap"`
}
type SpuPremiumGroupBasicDTO struct {
    /**
    * 配料属性组编码
    */
    GroupCode string `json:"groupCode"`
    /**
    * 配料属性组名称
    */
    GroupName string `json:"groupName"`
    /**
    * 配料属性组排序
    */
    Rank int32 `json:"rank"`
    /**
    * 配料组是否支持选多个配料 true：支持 false：不支持
    */
    ChooseMultiPremium bool `json:"chooseMultiPremium"`
    /**
    * 配料组下配料是否支持选多份 true：支持 false：不支持
    */
    MultiChoosePremium bool `json:"multiChoosePremium"`
}
type TimeRangeDTO struct {
    /**
    * 开始时间
    */
    Start string `json:"start"`
    /**
    * 结束时间
    */
    End string `json:"end"`
}
type PictureDTO struct {
    /**
    * 头图
    */
    HeadPictureList []string `json:"headPictureList"`
    /**
    * 描述图片
    */
    DescriptionPictureList []string `json:"descriptionPictureList"`
    /**
    * 背景图片
    */
    BackgroundPictureList []string `json:"backgroundPictureList"`
}
type BackgroundCategoryBasicDTO struct {
    /**
    * 后台类目 id
    */
    VendorBackgroundCategoryId string `json:"vendorBackgroundCategoryId"`
    /**
    * 后台类目名称
    */
    VendorBackgroundCategoryName string `json:"vendorBackgroundCategoryName"`
    /**
    * 后台类目英文名称
    */
    VendorBackgroundCategoryEnglishName string `json:"vendorBackgroundCategoryEnglishName"`
}
type VendorSkuStockDTO struct {
    /**
    * 门店 id
    */
    VendorShopId string `json:"vendorShopId"`
    /**
    * 商品 spuId
    */
    VendorSpuId string `json:"vendorSpuId"`
    /**
    * 商品 skuId
    */
    VendorSkuId string `json:"vendorSkuId"`
    /**
    * 库存数量
    */
    Stock int32 `json:"stock"`
    /**
    * 0：品牌商品 1：门店商品 默认品牌商品
    */
    ProductScope int32 `json:"productScope"`
}
type VendorSpuSaleAttributeGroupDTO struct {
    /**
    * 商品售卖属性组基本信息
    */
    SpuSaleAttributeGroupBasic SpuSaleAttributeGroupBasicDTO `json:"spuSaleAttributeGroupBasic"`
    /**
    * 售卖属性组关联的属性值 key：售卖属性于组中的展示顺序 value：售卖属性 code
    */
    VendorSpuSaleAttributeCodeMap string `json:"vendorSpuSaleAttributeCodeMap"`
    /**
    * 星巴克品牌定制，其他品牌不关注 售卖属性组关联的附加冷热配料 key：属性值于组中的展示顺序 value：附加冷热配料 code
    */
    VendorSpuPremiumCodeMap string `json:"vendorSpuPremiumCodeMap"`
}
type SkuSaleAttributeBasicDTO struct {
    /**
    * 售卖属性组名称
    */
    AttributeGroupName string `json:"attributeGroupName"`
    /**
    * 售卖属性组编码
    */
    AttributeGroupCode string `json:"attributeGroupCode"`
    /**
    * 售卖属性名称
    */
    AttributeName string `json:"attributeName"`
    /**
    * 售卖属性编码
    */
    AttributeCode string `json:"attributeCode"`
}
type VendorSpuDTO struct {
    /**
    * 商品基本信息
    */
    SpuBasic SpuBasicDTO `json:"spuBasic"`
    /**
    * 商品类型 NORMAL_SPU(1, "普通商品") FIXED_COMBO(2, "固定套餐") DINNERWARE(3, "餐具")
    */
    Type int32 `json:"type"`
    /**
    * 商品 SKU 列表
    */
    VendorSkuList []VendorSkuDTO `json:"vendorSkuList"`
    /**
    * defaultVendorSkuId
    */
    DefaultVendorSkuId string `json:"defaultVendorSkuId"`
    /**
    * 商品后台类目基本信息，用于开店宝后台配置营销活动时，按商品后台分类进行商品圈选。后台类目品牌自定义即可
    */
    BackgroundCategoryBasic BackgroundCategoryBasicDTO `json:"backgroundCategoryBasic"`
    /**
    * 商品售卖属性组列表 当 SKU 数 > 1 时必填；当 SKU 数 = 1 时，可不填
    */
    VendorSpuSaleAttributeGroupList []VendorSpuSaleAttributeGroupDTO `json:"vendorSpuSaleAttributeGroupList"`
    /**
    * 商品配料属性组列表 若商品没有配料，则不填
    */
    VendorSpuPremiumGroupList []VendorSpuPremiumGroupDTO `json:"vendorSpuPremiumGroupList"`
}
type MediaDTO struct {
    Picture PictureDTO `json:"picture"`
}
type ResultData struct {
    /**
    * 商品信息
    */
    VendorSpu VendorSpuDTO `json:"vendorSpu"`
    /**
    * 类目列表
    */
    VendorSellerCategoryNameList []string `json:"vendorSellerCategoryNameList"`
    /**
    * 商品门店上下架状态
    */
    ShelfStatus int32 `json:"shelfStatus"`
    /**
    * SKU 库存
    */
    VendorSkuStockList []VendorSkuStockDTO `json:"vendorSkuStockList"`
}
type VendorSpuPremiumGroupDTO struct {
    /**
    * 配料组基本信息
    */
    SpuPremiumGroupBasic SpuPremiumGroupBasicDTO `json:"spuPremiumGroupBasic"`
    /**
    * 配料组关联的配料 key：配料于组中的展示顺序 value：配料 code
    */
    VendorSpuPremiumCodeMap string `json:"vendorSpuPremiumCodeMap"`
    /**
    * 若配料组之间具有父子层级，则需要配置配料组对应的父配料 code 例如星巴克品牌中，配料「低因」下挂子配料性组【萃取方式】，那么【萃取方式】组的 parentPremiumCode 为「低因」的属性值编码 一般品牌不需要关注此字段，不填即可
    */
    ParentPremiumCode string `json:"parentPremiumCode"`
}
type SkuBasicDTO struct {
    /**
    * 商品 SKU 编码
    */
    VendorSkuId string `json:"vendorSkuId"`
    /**
    * SKU 价格，单位分
    */
    Price int64 `json:"price"`
    /**
    * 会员价，单位分
    */
    MemberPrice int64 `json:"memberPrice"`
    /**
    * 全国上下架状态 0 ：下架 1：上架 默认填 1 即可
    */
    ShelfStatus int32 `json:"shelfStatus"`
    /**
    * 扩展属性
    */
    ExtendInfoMap ExtendInfo `json:"extendInfoMap"`
}
type TimeDTO struct {
    /**
    * 一周的第几天，eg： 1：星期一 2：星期二 ...以此类推
    */
    DayOfWeek int32 `json:"dayOfWeek"`
    /**
    * 时间段，10:20-12:00
    */
    Range []TimeRangeDTO `json:"range"`
}
type QueryVendorSpuRequest struct {
    /**
    *  SPU 编码 
    */
    VendorSpuId string `json:"vendorSpuId"`
    /**
    *  门店Id 
    */
    VendorShopId string `json:"vendorShopId"`
}
type SpuSaleAttributeGroupBasicDTO struct {
    /**
    * 属性组编码
    */
    GroupCode string `json:"groupCode"`
    /**
    * 属性组名称
    */
    GroupName string `json:"groupName"`
    /**
    * 属性组排序
    */
    Rank int32 `json:"rank"`
}
type VendorSkuDTO struct {
    /**
    * 商品 SKU 基本信息
    */
    SkuBasic SkuBasicDTO `json:"skuBasic"`
    /**
    * 组成该 SKU 的售卖属性列表
    */
    SkuSaleAttributeBasicList []SkuSaleAttributeBasicDTO `json:"skuSaleAttributeBasicList"`
}
type QueryVendorSpuResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}
type ExtendInfo struct {
}



func (req *QueryVendorSpuRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryVendorSpuResponse, error) {
    resp, err := client.InvokeApi(query_vendor_spu_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryVendorSpuResponse
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

func (response *QueryVendorSpuResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

