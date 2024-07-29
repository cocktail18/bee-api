package saleattrcreate

/**
* 创建或更新售卖属性（必接）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const saleattr_create_url = "/foodmop/sku/saleattr/create"

type SpuSaleAttributeBasicDTO struct {
    /**
    *  售卖属性编码 
    */
    SaleAttributeCode string `json:"saleAttributeCode"`
    /**
    *  售卖属性名称 
    */
    SaleAttributeValue string `json:"saleAttributeValue"`
    /**
    *  售卖属性排序，默认填 1 即可 
    */
    Rank int32 `json:"rank"`
    /**
    *  售卖属性描述 
    */
    Description string `json:"description"`
    /**
    *  媒介信息（图片等），目前不支持售卖属性添加图片 
    */
    Media MediaDTO `json:"media"`
    /**
    *  全国上下架状态 1：上架，0：下架。默认填 1 即可，表示 售卖属性 全国上架 
    */
    ShelfStatus int32 `json:"shelfStatus"`
    /**
    *  份量大小，一般用于描述杯型的毫升数描述 
    */
    Weight string `json:"weight"`
    /**
    *  份量单位 ，一般用于描述杯型的毫升数的单位 
    */
    WeightUnit string `json:"weightUnit"`
    /**
    *  标识附加售卖属性还是标准售卖属性： 1：标准售卖属性 2：附加售卖属性 只适用于星巴克品牌，其他品牌只需要写死 type = 1 即可 
    */
    Type int32 `json:"type"`
    /**
    *  星巴克品牌特殊定制，其他品牌不关注 附加售卖属性关联的售卖属性编码，如微热的 relatedSaleAttributeCode = 热的属性编码 
    */
    RelatedSaleAttributeCode string `json:"relatedSaleAttributeCode"`
    /**
    *  是否可用会员积分抵扣 true：可以 false：不可 默认填 false 
    */
    UseMemberPoint bool `json:"useMemberPoint"`
    /**
    *  数据来源 1：品牌自建 2：品牌下发 3：门店自建 默认为1 
    */
    Source int32 `json:"source"`
    /**
    *  厂商门店 id source == 2 或source ==3时必填 
    */
    VendorShopId string `json:"vendorShopId"`
    /**
    *  售卖属性标识 source == 2 时，必填值为模板模板售卖属性 id 
    */
    TradeMarkCode string `json:"tradeMarkCode"`
}
type SaleattrCreateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * key：售卖属性值编码 value：同步结果，成功则返回 “成功”，失败则返回相应的异常信息
    */
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}
type MediaDTO struct {
    /**
    *  图片资源 
    */
    Picture PictureDTO `json:"picture"`
}
type ResultData struct {
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
type SaleattrCreateRequest struct {
    /**
    *  商品售卖属性列表 
    */
    SpuSaleAttributeBasicList []SpuSaleAttributeBasicDTO `json:"spuSaleAttributeBasicList"`
}



func (req *SaleattrCreateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*SaleattrCreateResponse, error) {
    resp, err := client.InvokeApi(saleattr_create_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response SaleattrCreateResponse
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

func (response *SaleattrCreateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

