package queryshopproductshelfstatus

/**
* 查询门店上架的商品
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_shop_product_shelf_status_url = "/foodmop/sku/queryByShop"

type QueryShopProductShelfStatusRequest struct {
    /**
    *  门店id 
    */
    VendorShopId string `json:"vendorShopId"`
}
type QueryShopProductShelfStatusResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 上架商品 id 列表
    */
    Data []string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *QueryShopProductShelfStatusRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryShopProductShelfStatusResponse, error) {
    resp, err := client.InvokeApi(query_shop_product_shelf_status_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryShopProductShelfStatusResponse
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

func (response *QueryShopProductShelfStatusResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

