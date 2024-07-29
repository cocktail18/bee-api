package queryvendorpremium

/**
* 查询配料属性信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_vendor_premium_url = "/foodmop/sku/premium/query"

type SpuPremiumBasicDTO struct {
    /**
    * 配料编码
    */
    PremiumCode string `json:"premiumCode"`
    /**
    * 配料名称
    */
    PremiumValue string `json:"premiumValue"`
    /**
    * 排序
    */
    Rank int32 `json:"rank"`
    /**
    * 描述
    */
    Description string `json:"description"`
    /**
    * 媒介信息，配料暂不支持展示图片
    */
    Media MediaDTO `json:"media"`
    /**
    * 全国上下架状态 1：上架 0：下架
    */
    ShelfStatus int32 `json:"shelfStatus"`
    /**
    * 配料标签，该字段未实际投入使用，默认不填
    */
    TagList []string `json:"tagList"`
    /**
    * 单位，如“泵”
    */
    Unit string `json:"unit"`
    /**
    * 价格，单位分
    */
    Price int64 `json:"price"`
    /**
    * 会员价，单位分
    */
    MemberPrice int64 `json:"memberPrice"`
    /**
    * 是否可用会员积分抵扣
    */
    UseMemberPoint bool `json:"useMemberPoint"`
    /**
    * 配料类型 1：标准配料 2：特殊配料，目前只有【换购】是特殊配料 3：附加冷热配料 2、3 类型属于星巴克品牌特殊定制
    */
    Type int32 `json:"type"`
    /**
    * 星巴克品牌特殊定制，其他品牌不关注 附加冷热关联的售卖属性编码，如微热的 relatedSaleAttributeCode = 热的属性编码
    */
    RelatedSaleAttributeCode string `json:"relatedSaleAttributeCode"`
}
type MediaDTO struct {
    /**
    * 图片资源
    */
    Picture PictureDTO `json:"picture"`
}
type QueryVendorPremiumRequest struct {
    /**
    *  配料属性编码 
    */
    PremiumCode string `json:"premiumCode"`
}
type ResultData struct {
    /**
    * 配料属性基本信息
    */
    SpuPremiumBasic SpuPremiumBasicDTO `json:"spuPremiumBasic"`
    /**
    * 只适用于星巴克品牌
    */
    ChildPremiumCode string `json:"childPremiumCode"`
}
type PictureDTO struct {
    /**
    * 头像
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
type QueryVendorPremiumResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *QueryVendorPremiumRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryVendorPremiumResponse, error) {
    resp, err := client.InvokeApi(query_vendor_premium_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryVendorPremiumResponse
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

func (response *QueryVendorPremiumResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

