package productproductloadproduct

/**
* 查询商品信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const product_product_loadproduct_url = "/ddzhkh/shangpin/product/loadproduct"

type Item struct {
    /**
    * 售卖SKU ID
    */
    ItemId int64 `json:"itemId"`
    /**
    * 售卖SKU名称
    */
    ItemName string `json:"itemName"`
    /**
    * 售卖SKU价格模型，注：类型为对象类型
    */
    ItemPrice ItemPrice `json:"itemPrice"`
    /**
    * 售卖SKU市场价格（原价），取值范围：0-99999999.99，精确到两位小数，单位：元。 
    */
    MarketPrice float64 `json:"marketPrice"`
    /**
    * 三方售卖SKU id，长度64
    */
    AppItemId string `json:"appItemId"`
    Attributes string `json:"attributes"`
}
type ItemPrice struct {
    /**
    * 售卖价格，取值范围：0-99999999.99，精确到两位小数，单位：元。
    */
    Price float64 `json:"price"`
}
type Product struct {
    /**
    * 美团商品ID
    */
    ProductId int64 `json:"productId"`
    /**
    * 美团商品名称
    */
    ProductName string `json:"productName"`
    /**
    * 三方商品ID
    */
    AppProductId string `json:"appProductId"`
    /**
    * 美团商品品类，三方需确定要接入商品品类和品类ID，例如体检预订[100016,735]
    */
    CategoryIds []int64 `json:"categoryIds"`
    OpPoiIds []string `json:"opPoiIds"`
    /**
    * 商品属性，业务参考属性文档,字典参考「创建商品」接口商品属性文档
    */
    Attributes string `json:"attributes"`
    /**
    * 商品的流程状态，10-编辑中，12-已提交，20-审核中，22-审核通过，24-审核驳回，30-已发布 
    */
    FlowStatus int32 `json:"flowStatus"`
    /**
    * 售卖SKU，注：类型为对象类型
    */
    Items []Item `json:"items"`
}
type ProductProductLoadproductRequest struct {
    /**
    *  美团商品ID 
    */
    ProductId int64 `json:"productId"`
}
type ProductProductLoadproductResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Product `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *ProductProductLoadproductRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ProductProductLoadproductResponse, error) {
    resp, err := client.InvokeApi(product_product_loadproduct_url, 59, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ProductProductLoadproductResponse
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

func (response *ProductProductLoadproductResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

