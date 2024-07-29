package productupdate

/**
* 更新门店商品库存（必接）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const product_update_url = "/foodmop/sku/update"

type ProductUpdateRequest struct {
    /**
    *  SKU 库存 
    */
    VendorSkuStockDTO VendorSkuStockDTO `json:"vendorSkuStockDTO"`
}
type ResultData struct {
}
type ProductUpdateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * key：vendorShopId-vendorSpuId-vendorSkuId value：同步结果，成功返回 “成功”，失败则返回相应的异常信息
    */
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
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



func (req *ProductUpdateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ProductUpdateResponse, error) {
    resp, err := client.InvokeApi(product_update_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ProductUpdateResponse
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

func (response *ProductUpdateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

