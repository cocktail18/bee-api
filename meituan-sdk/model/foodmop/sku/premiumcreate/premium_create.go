package premiumcreate

/**
* 创建或更新配料属性（必接）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const premium_create_url = "/foodmop/sku/premium/create"

type SpuPremiumBasicDTO struct {
    /**
    *  配料编码 不为空，编码不能重复 
    */
    PremiumCode string `json:"premiumCode"`
    /**
    *  配料名称 不为空 
    */
    PremiumValue string `json:"premiumValue"`
    /**
    *  排序，默认填 1 即可 不为空，rank &gt; 0 
    */
    Rank int32 `json:"rank"`
    /**
    *  描述 
    */
    Description string `json:"description"`
    /**
    *  媒介信息，配料暂不支持展示图片 
    */
    Media MediaDTO `json:"media"`
    /**
    *  全国上下架状态 1：上架 0：下架 默认填 1 即可，表示配料全国上架 不为空，只能为 0 或 1 
    */
    ShelfStatus int32 `json:"shelfStatus"`
    /**
    *  配料标签，该字段未实际投入使用，默认不填 
    */
    TagList []string `json:"tagList"`
    /**
    *  单位，如“泵” 
    */
    Unit string `json:"unit"`
    /**
    *  价格，单位分 不为空，price &gt;= 0 
    */
    Price int64 `json:"price"`
    /**
    *  会员价，单位分，目前会员价暂不支持，填 0 即可 memberPrice &gt;= 0 
    */
    MemberPrice int64 `json:"memberPrice"`
    /**
    *  是否可用会员积分抵扣 true：可以 false：不可以 默认 false 
    */
    UseMemberPoint bool `json:"useMemberPoint"`
    /**
    *  配料类型 1：标准配料 2：特殊配料，目前只有【换购】是特殊配料 3：附加冷热配料 4：做法 2、3 类型属于星巴克品牌特殊定制，其他品牌默认填 1 即可 
    */
    Type int32 `json:"type"`
    /**
    *  星巴克品牌特殊定制，其他品牌不关注 附加冷热关联的售卖属性编码，如微热的 relatedSaleAttributeCode = 热的属性编码 
    */
    RelatedSaleAttributeCode string `json:"relatedSaleAttributeCode"`
    /**
    *  数据来源 1：品牌自建 2：品牌下发 3：门店自建 默认为1 
    */
    Source int32 `json:"source"`
    /**
    *  厂商门店 id source == 2 或 source ==3时必填 
    */
    VendorShopId string `json:"vendorShopId"`
    /**
    *  售卖属性标识 source == 2 时，必填 值为模板配料属性 id 
    */
    TradeMarkCode string `json:"tradeMarkCode"`
}
type MediaDTO struct {
    /**
    *  图片资源 
    */
    Picture PictureDTO `json:"picture"`
}
type ResultData struct {
}
type PremiumCreateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * key：配料编码 value：同步结果，成功返回 “成功”，失败则返回相应的异常信息
    */
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}
type VendorSpuPremiumDTO struct {
    /**
    *  配料属性基本信息 
    */
    SpuPremiumBasic SpuPremiumBasicDTO `json:"spuPremiumBasic"`
    /**
    *  只适用于星巴克品牌，其他品牌不填即可 子属性 code，若「浓缩份数」属性下有子属性「换购」则「浓缩份数」的 childPremiumCode 是「换购」的配料编码 childPremiumCode 必须存在于配料池，childPremiumCode != spuPremiumBasic.premiumCode 
    */
    ChildPremiumCode string `json:"childPremiumCode"`
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
    /**
    *  背景图片 
    */
    BackgroundPictureList []string `json:"backgroundPictureList"`
}
type PremiumCreateRequest struct {
    /**
    *  商品配料属性列表 不为空，每次最多传 25 个 
    */
    VendorSpuPremiumList []VendorSpuPremiumDTO `json:"vendorSpuPremiumList"`
}



func (req *PremiumCreateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*PremiumCreateResponse, error) {
    resp, err := client.InvokeApi(premium_create_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response PremiumCreateResponse
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

func (response *PremiumCreateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

