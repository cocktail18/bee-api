package queryvendorspupool

/**
* 查询品牌商品池
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_vendor_spu_pool_url = "/foodmop/sku/queryVendorSpuPool"

type QueryVendorSpuPoolRequest struct {
    /**
    *  门店id 
    */
    VendorShopId string `json:"vendorShopId"`
}
type QueryVendorSpuPoolResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 商品 id 列表
    */
    Data []string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *QueryVendorSpuPoolRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryVendorSpuPoolResponse, error) {
    resp, err := client.InvokeApi(query_vendor_spu_pool_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryVendorSpuPoolResponse
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

func (response *QueryVendorSpuPoolResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

