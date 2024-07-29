package productproductcreate

/**
* 新增商品
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const product_product_create_url = "/ddzhkh/shangpin/product/create"

type ProductProductCreateRequest struct {
    Product Product `json:"product"`
}
type Item struct {
    /**
    *  售卖SKU名称，长度限制&lt;=30 
    */
    ItemName string `json:"itemName"`
    /**
    *  售卖SKU价格模型 
    */
    ItemPrice Price `json:"itemPrice"`
    /**
    *  售卖SKU价格类型，0-一口价 
    */
    PriceType int32 `json:"priceType"`
    /**
    *  售卖SKU市场价格（原价），取值范围：0-99999999.99，精确到两位小数，单位：元。如101.01，表示101.01元1分 
    */
    MarketPrice float64 `json:"marketPrice"`
    /**
    *  ISV售卖SKU id，长度64 
    */
    AppItemId string `json:"appItemId"`
    /**
    *  售卖SKU属性 类型Map&lt;Long,List&lt;String&gt;&gt; 
    */
    Attributes string `json:"attributes"`
    /**
    *  售卖SKU库存类型，用来区分场景，1-三方直连，注：在化妆品商品中,该字段无意义,写1即可 
    */
    StockScene int32 `json:"stockScene"`
}
type ProductProductCreateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Resp `json:"data"`
    TraceId string `json:"traceId"`
}
type Price struct {
    /**
    *  售卖价格，取值范围：0-99999999.99，精确到两位小数，单位：元。如101.01，表示101.01元1分 
    */
    Price float64 `json:"price"`
}
type Resp struct {
    Attributes string `json:"attributes"`
    /**
    * 售卖SKU
    */
    Items []Item `json:"items"`
    ProductId int64 `json:"productId"`
    /**
    * 美团商品名称，长度限制<=30
    */
    ProductName string `json:"productName"`
    /**
    * 三方商品ID
    */
    AppProductId string `json:"appProductId"`
    /**
    * 美团商品品类，三方需确定要接入商品品类和品类ID，例如体检预订：[100016,735]，场馆预订：[100060,1650],化妆品电商[100048,1005]
    */
    CategoryIds []int64 `json:"categoryIds"`
    OpPoiIds []string `json:"opPoiIds"`
}
type Product struct {
    /**
    *  美团商品名称，长度限制&lt;=30 
    */
    ProductName string `json:"productName"`
    /**
    *  三方商品ID 
    */
    AppProductId string `json:"appProductId"`
    /**
    *  美团商品品类，三方需确定要接入商品品类和品类ID，例如体检预订：[100016,735]，场馆预订：[100060,1650],化妆品电商[100048,1005] 
    */
    CategoryIds []int64 `json:"categoryIds"`
    /**
    *  美团店铺id 
    */
    OpPoiIds []string `json:"opPoiIds"`
    /**
    *  商品属性，Map&lt;Long,List&lt;String&gt;&gt;格式 
    */
    Attributes string `json:"attributes"`
    /**
    *  售卖SKU 
    */
    Items []Item `json:"items"`
}



func (req *ProductProductCreateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ProductProductCreateResponse, error) {
    resp, err := client.InvokeApi(product_product_create_url, 59, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ProductProductCreateResponse
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

func (response *ProductProductCreateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

