package queryvendorsaleattribute

/**
* 查询售卖属性信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_vendor_sale_attribute_url = "/foodmop/sku/saleattr/query"

type MediaDTO struct {
    /**
    * 图片资源
    */
    Picture PictureDTO `json:"picture"`
}
type RequestData struct {
    /**
    * 售卖属性编码
    */
    SaleAttributeCode string `json:"saleAttributeCode"`
    /**
    * 售卖属性名称
    */
    SaleAttributeValue string `json:"saleAttributeValue"`
    /**
    * 售卖属性排序，默认填 1 即可
    */
    Rank int32 `json:"rank"`
    /**
    * 售卖属性描述
    */
    Description string `json:"description"`
    /**
    * 媒介信息（图片等），目前不支持售卖属性添加图片
    */
    Media MediaDTO `json:"media"`
    /**
    * 全国上下架状态 1：上架，0：下架。默认填 1 即可，表示配料全国上
    */
    ShelfStatus int32 `json:"shelfStatus"`
    /**
    * 份量大小，一般用于描述杯型的毫升数描述
    */
    Weight string `json:"weight"`
    /**
    * 份量单位 ，一般用于描述杯型的毫升数的单位
    */
    WeightUnit string `json:"weightUnit"`
    /**
    * 标识附加售卖属性还是标准售卖属性
    */
    Type int32 `json:"type"`
    /**
    * 星巴克品牌特殊定制，其他品牌不关注
    */
    RelatedSaleAttributeCode string `json:"relatedSaleAttributeCode"`
    /**
    * 是否可用会员积分抵扣
    */
    UseMemberPoint bool `json:"useMemberPoint"`
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
type QueryVendorSaleAttributeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data RequestData `json:"data"`
    TraceId string `json:"traceId"`
}
type QueryVendorSaleAttributeRequest struct {
    /**
    *  售卖属性编码 
    */
    SaleAttributeCode string `json:"saleAttributeCode"`
}



func (req *QueryVendorSaleAttributeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryVendorSaleAttributeResponse, error) {
    resp, err := client.InvokeApi(query_vendor_sale_attribute_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryVendorSaleAttributeResponse
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

func (response *QueryVendorSaleAttributeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

