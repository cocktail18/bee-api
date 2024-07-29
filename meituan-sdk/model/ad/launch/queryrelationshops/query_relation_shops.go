package queryrelationshops

/**
* 查询账号关联门店
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_relation_shops_url = "/ad/launch/queryRelationShops"

type LaunchObj struct {
    /**
    * 门店名称
    */
    ShopName string `json:"shopName"`
    /**
    * 门店id
    */
    ShopId int64 `json:"shopId"`
    /**
    * 门店uuid
    */
    ShopUuid string `json:"shopUuid"`
    /**
    * 分店名称
    */
    ShopBranchName string `json:"shopBranchName"`
}
type QueryRelationShopsRequest struct {
    /**
    *  子账号商家id 
    */
    SonAccountId int32 `json:"sonAccountId"`
}
type QueryRelationShopsResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []LaunchObj `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *QueryRelationShopsRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryRelationShopsResponse, error) {
    resp, err := client.InvokeApi(query_relation_shops_url, 22, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryRelationShopsResponse
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

func (response *QueryRelationShopsResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

