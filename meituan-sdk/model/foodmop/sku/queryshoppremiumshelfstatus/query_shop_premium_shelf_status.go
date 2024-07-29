package queryshoppremiumshelfstatus

/**
* 查询门店上架的配料
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_shop_premium_shelf_status_url = "/foodmop/sku/shelf/premium/query"

type QueryShopPremiumShelfStatusRequest struct {
    /**
    *  门店id 
    */
    VendorShopId string `json:"vendorShopId"`
}
type QueryShopPremiumShelfStatusResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 上架的配料编码列表
    */
    Data []string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *QueryShopPremiumShelfStatusRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryShopPremiumShelfStatusResponse, error) {
    resp, err := client.InvokeApi(query_shop_premium_shelf_status_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryShopPremiumShelfStatusResponse
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

func (response *QueryShopPremiumShelfStatusResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

