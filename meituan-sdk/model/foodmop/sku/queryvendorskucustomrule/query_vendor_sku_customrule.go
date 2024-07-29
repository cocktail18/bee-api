package queryvendorskucustomrule

/**
* 查询商品客制化规则
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_vendor_sku_customrule_url = "/foodmop/sku/customrule/query"

type QueryVendorSkuCustomruleRequest struct {
    /**
    *  spuId 
    */
    VendorSkuId string `json:"vendorSkuId"`
    /**
    *  skuId 
    */
    VendorSpuId string `json:"vendorSpuId"`
}
type ResultData struct {
    /**
    * 商品 SPU 编码
    */
    VendorSpuId string `json:"vendorSpuId"`
    /**
    * 商品 SKU 编码
    */
    VendorSkuId string `json:"vendorSkuId"`
    /**
    * 客制化规则列表
    */
    SkuCustomRuleBasicList []SkuCustomRuleBasicDTO `json:"skuCustomRuleBasicList"`
}
type QueryVendorSkuCustomruleResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}
type SkuCustomRuleBasicDTO struct {
    /**
    * 规则类型： 10：配料组客制化规则 20：配料客制化规则
    */
    RuleType int32 `json:"ruleType"`
    /**
    * 配料组编码
    */
    GroupCode string `json:"groupCode"`
    /**
    * 配料编码
    */
    ItemCode string `json:"itemCode"`
    /**
    * 是否必选
    */
    MustSelect bool `json:"mustSelect"`
    /**
    * 是否默认选中
    */
    DefaultSelect bool `json:"defaultSelect"`
    /**
    * 属性组是否折叠
    */
    Fold bool `json:"fold"`
    /**
    * 属性组 or 属性值展示排序
    */
    Rank int32 `json:"rank"`
    /**
    * 最小选择数，默认 1
    */
    MinChoice int32 `json:"minChoice"`
    /**
    * 最大选择数，默认 1
    */
    MaxChoice int32 `json:"maxChoice"`
    /**
    * 配料 or 配料组是否在页面隐藏
    */
    Hide bool `json:"hide"`
    /**
    * 是否支持换购
    */
    Exchange bool `json:"exchange"`
    /**
    * 最大免费数，默认 0
    */
    MaxFreeNum int32 `json:"maxFreeNum"`
    /**
    * 默认选中数，默认 0
    */
    DefaultSelectNum int32 `json:"defaultSelectNum"`
    /**
    * 每份数量，按份计价用，一般品牌无特殊要求，默认填 1 即可
    */
    NumOfPortion int32 `json:"numOfPortion"`
}



func (req *QueryVendorSkuCustomruleRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryVendorSkuCustomruleResponse, error) {
    resp, err := client.InvokeApi(query_vendor_sku_customrule_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryVendorSkuCustomruleResponse
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

func (response *QueryVendorSkuCustomruleResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

