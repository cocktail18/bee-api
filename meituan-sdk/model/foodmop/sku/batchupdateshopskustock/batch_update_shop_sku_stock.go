package batchupdateshopskustock

/**
* 批量更新门店商品库存（非必接）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const batch_update_shop_sku_stock_url = "/foodmop/sku/batchUpdateSkuStock"

type BatchUpdateShopSkuStockResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * key：vendorShopId-vendorSpuId-vendorSkuId value：同步结果，成功返回 “成功”，失败则返回相应的异常信息
    */
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}
type ResultData struct {
}
type VendorSkuStockDTO struct {
    /**
    *  门店 id 
    */
    VendorShopId string `json:"vendorShopId"`
    /**
    *  商品 spuId 
    */
    VendorSpuId string `json:"vendorSpuId"`
    /**
    *  商品 skuId 
    */
    VendorSkuId string `json:"vendorSkuId"`
    /**
    *  库存数量 
    */
    Stock int32 `json:"stock"`
    /**
    *  0：品牌商品 1：门店商品 默认品牌商品 
    */
    ProductScope int32 `json:"productScope"`
}
type BatchUpdateShopSkuStockRequest struct {
    /**
    *  SKU 库存 最多传 50 个，要求所有 vendorShopId 都在美团侧存在映射，若存在无映射 vendorShopId，失败并在错误信息中返回无映射 vendorShopId 
    */
    VendorSkuStockList []VendorSkuStockDTO `json:"vendorSkuStockList"`
}



func (req *BatchUpdateShopSkuStockRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*BatchUpdateShopSkuStockResponse, error) {
    resp, err := client.InvokeApi(batch_update_shop_sku_stock_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response BatchUpdateShopSkuStockResponse
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

func (response *BatchUpdateShopSkuStockResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

