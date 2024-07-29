package queryvendorsaleattributecode

/**
* 查询品牌所有售卖属性编码
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_vendor_sale_attribute_code_url = "/foodmop/sku/saleattr/queryAll"

type ResultData struct {
}
type QueryVendorSaleAttributeCodeRequest struct {
}
type QueryVendorSaleAttributeCodeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 品牌所有售卖属性编码
    */
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *QueryVendorSaleAttributeCodeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryVendorSaleAttributeCodeResponse, error) {
    resp, err := client.InvokeApi(query_vendor_sale_attribute_code_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryVendorSaleAttributeCodeResponse
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

func (response *QueryVendorSaleAttributeCodeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

