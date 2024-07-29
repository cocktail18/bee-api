package productdealgroupupdateprice

/**
* 修改团单价格
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const product_dealgroup_updateprice_url = "/ddzhkh/shangpin/dealgroup/updateprice"

type Response struct {
    /**
    * 团单id
    */
    DealGroupId int32 `json:"dealGroupId"`
}
type ProductDealgroupUpdatepriceRequest struct {
    /**
    *  团单id 
    */
    DealGroupId int32 `json:"dealGroupId"`
    /**
    *  来源 
    */
    Source string `json:"source"`
    /**
    *  团单售价 
    */
    Price string `json:"price"`
}
type ProductDealgroupUpdatepriceResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Response `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *ProductDealgroupUpdatepriceRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ProductDealgroupUpdatepriceResponse, error) {
    resp, err := client.InvokeApi(product_dealgroup_updateprice_url, 59, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ProductDealgroupUpdatepriceResponse
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

func (response *ProductDealgroupUpdatepriceResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

