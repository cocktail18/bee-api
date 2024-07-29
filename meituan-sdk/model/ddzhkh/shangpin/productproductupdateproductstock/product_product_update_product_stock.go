package productproductupdateproductstock

/**
* 更新商品库存
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const product_product_update_product_stock_url = "/ddzhkh/shangpin/product/updateproductstock"

type ShopStock struct {
    /**
    *  门店ID 
    */
    OpPoiId string `json:"opPoiId"`
    /**
    *  更新分门店总库存为，shopStock 
    */
    Stock int32 `json:"stock"`
}
type Resp struct {
    /**
    * 美团商品ID
    */
    ProductId int64 `json:"productId"`
    /**
    * 售卖SKU，注：类型为对象类型，详见扩展参数item_stock_infos 
    */
    ItemStockInfos []ItemStockInfo `json:"itemStockInfos"`
    /**
    * 商品的流程状态，10-编辑中，12-已提交，20-审核中，22-审核通过，24-审核驳回，30-已发布
    */
    FlowStatus int32 `json:"flowStatus"`
}
type ItemStock struct {
    /**
    *  售卖SKU ID 
    */
    ItemId int64 `json:"itemId"`
    /**
    *  SKU商品库存 
    */
    Stock int32 `json:"stock"`
    /**
    *  SKU分门店库存更新数组（和 stock 有且只有一个） 注：类型为对象类型，详见扩展参数shop_stocks 
    */
    ShopStocks []ShopStock `json:"shopStocks"`
}
type ItemStockInfo struct {
    /**
    * 售卖SKU ID
    */
    ItemId int64 `json:"itemId"`
    /**
    * SKU的商品库存
    */
    Stock int32 `json:"stock"`
    /**
    * 售卖SKU属性，业务参考属性文档，注：类型为对象类型，详见扩展参数shop_stock_infos 
    */
    ShopStockInfos []ShopStockInfo `json:"shopStockInfos"`
}
type ShopStockInfo struct {
    /**
    * 门店ID
    */
    OpPoiId string `json:"opPoiId"`
    /**
    * 美团门店名称
    */
    ShopName string `json:"shopName"`
    /**
    * 总库存
    */
    Stock int32 `json:"stock"`
    /**
    * 已售库存
    */
    SaleStock int32 `json:"saleStock"`
    /**
    * 剩余库存
    */
    RemainStock int32 `json:"remainStock"`
}
type ProductProductUpdateProductStockResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Resp `json:"data"`
    TraceId string `json:"traceId"`
}
type ProductProductUpdateProductStockRequest struct {
    /**
    *  美团商品ID 
    */
    ProductId int64 `json:"productId"`
    /**
    *  售卖SKU 注：类型为对象类型，详见扩展参数item_stocks 
    */
    ItemStocks []ItemStock `json:"itemStocks"`
}



func (req *ProductProductUpdateProductStockRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ProductProductUpdateProductStockResponse, error) {
    resp, err := client.InvokeApi(product_product_update_product_stock_url, 59, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ProductProductUpdateProductStockResponse
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

func (response *ProductProductUpdateProductStockResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

