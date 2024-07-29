package productdealgroupupdatestock

/**
* 修改团单库存
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const product_dealgroup_updatestock_url = "/ddzhkh/shangpin/dealgroup/updatestock"

type Response struct {
    DealGroupId int32 `json:"dealGroupId"`
}
type ProductDealgroupUpdatestockRequest struct {
    /**
    *  来源 
    */
    Source string `json:"source"`
    /**
    *  团单ID 
    */
    DealGroupId int32 `json:"dealGroupId"`
    /**
    *  库存 
    */
    Stock int32 `json:"stock"`
}
type ProductDealgroupUpdatestockResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Response `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *ProductDealgroupUpdatestockRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ProductDealgroupUpdatestockResponse, error) {
    resp, err := client.InvokeApi(product_dealgroup_updatestock_url, 59, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ProductDealgroupUpdatestockResponse
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

func (response *ProductDealgroupUpdatestockResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

