package productproductoffline

/**
* 商品下架
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const product_product_offline_url = "/ddzhkh/shangpin/product/offline"

type ProductProductOfflineResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Data struct {
    /**
    * 商品ID
    */
    ProductId int64 `json:"productId"`
}
type ProductProductOfflineRequest struct {
    /**
    *  商品ID 
    */
    ProductId int64 `json:"productId"`
}



func (req *ProductProductOfflineRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ProductProductOfflineResponse, error) {
    resp, err := client.InvokeApi(product_product_offline_url, 59, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ProductProductOfflineResponse
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

func (response *ProductProductOfflineResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

