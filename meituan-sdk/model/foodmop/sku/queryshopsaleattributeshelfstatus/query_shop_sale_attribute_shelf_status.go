package queryshopsaleattributeshelfstatus

/**
* 查询门店上架的售卖属性
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_shop_sale_attribute_shelf_status_url = "/foodmop/sku/shelf/saleattr/query"

type QueryShopSaleAttributeShelfStatusResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 上架售卖属性编码列表
    */
    Data []string `json:"data"`
    TraceId string `json:"traceId"`
}
type QueryShopSaleAttributeShelfStatusRequest struct {
    /**
    *  门店id 
    */
    VendorShopId string `json:"vendorShopId"`
}



func (req *QueryShopSaleAttributeShelfStatusRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryShopSaleAttributeShelfStatusResponse, error) {
    resp, err := client.InvokeApi(query_shop_sale_attribute_shelf_status_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryShopSaleAttributeShelfStatusResponse
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

func (response *QueryShopSaleAttributeShelfStatusResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

