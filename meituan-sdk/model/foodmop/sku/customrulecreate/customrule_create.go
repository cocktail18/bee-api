package customrulecreate

/**
* 创建或更改商品客制化规则（必接）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const customrule_create_url = "/foodmop/sku/customrule/create"

type CustomruleCreateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * key：vendorSpuId-vendorSkuId value：同步结果，成功返回 “成功”，失败则返回相应的异常信息
    */
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}
type ResultData struct {
}
type SkuCustomRuleBasicDTO struct {
    /**
    *  规则类型： 10：配料组客制化规则 20：配料客制化规则 
    */
    RuleType int32 `json:"ruleType"`
    /**
    *  配料组编码 ruleType == 10 时必填 
    */
    GroupCode string `json:"groupCode"`
    /**
    *  配料编码 ruleType == 20 时必填 
    */
    ItemCode string `json:"itemCode"`
    /**
    *  是否必选 true：是 false：否 
    */
    MustSelect bool `json:"mustSelect"`
    /**
    *  是否默认选中 true：是 false：否 默认 false ruleType == 20 时必填 若 mustSelect == true，则需保证 defaultSelect == true 
    */
    DefaultSelect bool `json:"defaultSelect"`
    /**
    *  属性组是否折叠 true：折叠 false：不折叠 默认 false ruleType == 10 时必填 
    */
    Fold bool `json:"fold"`
    /**
    *  属性组 or 属性值展示排序 不为空，rank &gt; 0 
    */
    Rank int32 `json:"rank"`
    /**
    *  最小选择数，默认 1 ruleType == 20 时必填 不为空，0 &lt;= minChoice &lt;= maxChoice 
    */
    MinChoice int32 `json:"minChoice"`
    /**
    *  最大选择数，默认 1 ruleType == 20 时必填 maxChoice &gt;= minChoice 
    */
    MaxChoice int32 `json:"maxChoice"`
    /**
    *  配料 or 配料组是否在页面隐藏 true：是 false：否 
    */
    Hide bool `json:"hide"`
    /**
    *  是否支持换购 true：是 false：否 星巴克品牌定制，其他品牌设为 false 
    */
    Exchange bool `json:"exchange"`
    /**
    *  最大免费数，默认 0 
    */
    MaxFreeNum int32 `json:"maxFreeNum"`
    /**
    *  默认选中数，默认 0 
    */
    DefaultSelectNum int32 `json:"defaultSelectNum"`
    /**
    *  每份数量，按份计价用，一般品牌无特殊要求，默认填 1 即可 
    */
    NumOfPortion int32 `json:"numOfPortion"`
}
type CustomruleCreateRequest struct {
    /**
    *  商品 SPU 客制化规则列表 
    */
    VendorSkuCustomRuleList []VendorSkuCustomRuleDTO `json:"vendorSkuCustomRuleList"`
}
type VendorSkuCustomRuleDTO struct {
    /**
    *  商品 SPU 编码 
    */
    VendorSpuId string `json:"vendorSpuId"`
    /**
    *  商品 SKU 编码 
    */
    VendorSkuId string `json:"vendorSkuId"`
    /**
    *  客制化规则列表 
    */
    SkuCustomRuleBasicList []SkuCustomRuleBasicDTO `json:"skuCustomRuleBasicList"`
    /**
    *  厂商门店 id 门店商品需要传入 
    */
    VendorShopId string `json:"vendorShopId"`
}



func (req *CustomruleCreateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CustomruleCreateResponse, error) {
    resp, err := client.InvokeApi(customrule_create_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CustomruleCreateResponse
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

func (response *CustomruleCreateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

