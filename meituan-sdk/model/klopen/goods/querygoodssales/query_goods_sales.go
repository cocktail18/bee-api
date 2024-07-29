package querygoodssales

/**
* 查询sku售卖信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_goods_sales_url = "/kl/open/goods/sku/sales"

type Data struct {
    Total string `json:"total"`
    List string `json:"list"`
}
type QueryGoodsSalesResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type QueryGoodsSalesRequest struct {
    MerchantCode string `json:"merchantCode"`
    SkuCodeList []int32 `json:"skuCodeList"`
    /**
    *  品牌标识，除华住外其他服务商必传 
    */
    BrandIdentify string `json:"brandIdentify"`
}



func (req *QueryGoodsSalesRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryGoodsSalesResponse, error) {
    resp, err := client.InvokeApi(query_goods_sales_url, 27, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryGoodsSalesResponse
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

func (response *QueryGoodsSalesResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

