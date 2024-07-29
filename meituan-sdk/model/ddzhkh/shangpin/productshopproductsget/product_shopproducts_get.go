package productshopproductsget

/**
* 店铺商品查询
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const product_shopproducts_get_url = "/ddzhkh/shangpin/shopProducts/get"

type ProductShopproductsGetRequest struct {
    OpPoiIds []string `json:"opPoiIds"`
    /**
    *  商品分类（存量可不传），医美预付-850，酒吧-2000327 
    */
    SpuType string `json:"spuType"`
    /**
    *  页码，从1开始 
    */
    PageNo int32 `json:"pageNo"`
    /**
    *  每页的大小，最大100 
    */
    PageSize int32 `json:"pageSize"`
}
type ProductShopproductsGetResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data PagingGetShopProductsResponse `json:"data"`
    TraceId string `json:"traceId"`
}
type ShopProductTO struct {
    /**
    * 状态，0-无效，1-上架，2-下架
    */
    Status int32 `json:"status"`
    /**
    * 产品ID
    */
    ProductId int64 `json:"productId"`
    /**
    * 产品名称
    */
    ProductName string `json:"productName"`
    AppProductId string `json:"appProductId"`
    /**
    * 流程状态，10-编辑中，12-已提交，20-审核中，22-审核通过，24-审核驳回，30-已发布
    */
    FlowStatus int32 `json:"flowStatus"`
    ShopProductItems []ShopProductItem `json:"shopProductItems"`
}
type PagingGetShopProductsResponse struct {
    /**
    * 页码 
    */
    PageNo int32 `json:"pageNo"`
    /**
    * 每页条数 
    */
    PageSize int32 `json:"pageSize"`
    ShopProducts []ShopProductTO `json:"shopProducts"`
    /**
    * 总条数
    */
    TotalCount int32 `json:"totalCount"`
}
type ShopProductItem struct {
    /**
    * 项目ID
    */
    ItemId int64 `json:"itemId"`
    /**
    * 项目名称
    */
    ItemName string `json:"itemName"`
    AppItemId string `json:"appItemId"`
}



func (req *ProductShopproductsGetRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ProductShopproductsGetResponse, error) {
    resp, err := client.InvokeApi(product_shopproducts_get_url, 59, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ProductShopproductsGetResponse
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

func (response *ProductShopproductsGetResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

