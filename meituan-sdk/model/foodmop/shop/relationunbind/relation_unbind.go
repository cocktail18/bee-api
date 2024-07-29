package relationunbind

/**
* 解绑门店（必接）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const relation_unbind_url = "/foodmop/shop/relation/unbind"

type RelationUnbindRequest struct {
    /**
    *  商家ERP上门店编码 
    */
    ErpShopId string `json:"erpShopId"`
}
type RelationUnbindResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *RelationUnbindRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*RelationUnbindResponse, error) {
    resp, err := client.InvokeApi(relation_unbind_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response RelationUnbindResponse
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

func (response *RelationUnbindResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

