package batchupdateskuprice

/**
* 批量更新门店维度 SKU 价格
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const batch_update_sku_price_url = "/foodmop/sku/batchUpdateSkuPrice"

type ResultData struct {
}
type PriceBusinessDimensionDTO struct {
    /**
    *  1：商品基础价格；2：门店聚合维度价格，3：门店维度 
    */
    BusinessDimensionType int32 `json:"businessDimensionType"`
    /**
    *  价格维度 id-value key value businessDimensionType POI_AGGREGATION 商家门店聚合 id（TAG_CODE门店标签接口) 2 时的 key POI 商家门店Id(即VendorShopId) 3 时的 key businessDimensionType == 2 时必传 value 必须为数字类型 例如瑞幸，POI_AGGREGATION = 门店价格等级 
    */
    BusinessDimensionMap BusinessDimension `json:"businessDimensionMap"`
    /**
    *  商品 spuId 长度&lt;=xx字符 
    */
    VendorSpuId string `json:"vendorSpuId"`
    /**
    *  商品 skuId 1、长度&lt;=xx字符 2、目前不支持一个品牌下不同vendorSpuId存在相同vendorSkuId 的情况 
    */
    VendorSkuId string `json:"vendorSkuId"`
}
type BatchUpdateSkuPriceRequest struct {
    /**
    *  SKU 价格 最多传 100 个，要求所有vendorShopId都在美团侧存在映射，若存在无映射vendorShopId,失败并在错误信息中返回无映射vendorShopId 
    */
    VendorSkuPriceList []VendorSkuPriceDTO `json:"vendorSkuPriceList"`
}
type BatchUpdateSkuPriceResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * key：商品价格主体维度 value：失败原因 若存在门店映射未映射，直接返回失败{“code”:300,”msg”:”部分门店未映射xxx”}
    */
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}
type VendorSkuPriceDTO struct {
    /**
    *  sku 价格维度 
    */
    PriceBusinessDimensionDTO PriceBusinessDimensionDTO `json:"priceBusinessDimensionDTO"`
    /**
    *  sku 价格，单位分 单位分，非负数 
    */
    Price int64 `json:"price"`
}
type BusinessDimension struct {
}



func (req *BatchUpdateSkuPriceRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*BatchUpdateSkuPriceResponse, error) {
    resp, err := client.InvokeApi(batch_update_sku_price_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response BatchUpdateSkuPriceResponse
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

func (response *BatchUpdateSkuPriceResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

