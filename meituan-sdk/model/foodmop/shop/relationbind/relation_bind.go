package relationbind

/**
* 绑定门店（必接）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const relation_bind_url = "/foodmop/shop/relation/bind"

type RelationBindResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}
type RelationBindRequest struct {
    /**
    *  商家ERP上门店编码 
    */
    ErpShopId string `json:"erpShopId"`
    /**
    *  商家POS上门店编码，如果没有则默认为ERP上门店编码 
    */
    PosShopId string `json:"posShopId"`
    /**
    *  美团门店ID 
    */
    PoiId int32 `json:"poiId"`
}



func (req *RelationBindRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*RelationBindResponse, error) {
    resp, err := client.InvokeApi(relation_bind_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response RelationBindResponse
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

func (response *RelationBindResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

