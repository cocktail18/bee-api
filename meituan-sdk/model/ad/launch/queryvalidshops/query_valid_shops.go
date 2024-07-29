package queryvalidshops

/**
* 判断门店是否可以投放广告
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_valid_shops_url = "/ad/launch/queryValidShops"

type ShopExInfo struct {
}
type QueryValidShopsResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * validShopIds合法的门店列表，invalidShopIds无法投放的门店
    */
    Data ShopExInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type QueryValidShopsRequest struct {
    /**
    *  校验门店id列表 
    */
    ShopList []int64 `json:"shopList"`
}



func (req *QueryValidShopsRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryValidShopsResponse, error) {
    resp, err := client.InvokeApi(query_valid_shops_url, 22, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryValidShopsResponse
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

func (response *QueryValidShopsResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

