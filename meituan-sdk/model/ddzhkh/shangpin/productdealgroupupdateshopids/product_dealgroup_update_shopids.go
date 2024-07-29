package productdealgroupupdateshopids

/**
* 修改团单适用门店
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const product_dealgroup_update_shopids_url = "/ddzhkh/shangpin/dealgroup/updateshopids"

type Response struct {
    /**
    * 团单id
    */
    DealGroupId int32 `json:"dealGroupId"`
}
type ProductDealgroupUpdateShopidsResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Response `json:"data"`
    TraceId string `json:"traceId"`
}
type ProductDealgroupUpdateShopidsRequest struct {
    /**
    *  团单ID 
    */
    DealGroupId int32 `json:"dealGroupId"`
    /**
    *  来源 
    */
    Source string `json:"source"`
    /**
    *  商户ID列表 
    */
    OpPoiIds []string `json:"opPoiIds"`
}



func (req *ProductDealgroupUpdateShopidsRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ProductDealgroupUpdateShopidsResponse, error) {
    resp, err := client.InvokeApi(product_dealgroup_update_shopids_url, 59, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ProductDealgroupUpdateShopidsResponse
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

func (response *ProductDealgroupUpdateShopidsResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

