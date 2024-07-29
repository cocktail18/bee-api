package shopmenudeletespecial

/**
* 删除门店特殊类目
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const shopmenu_delete_special_url = "/foodmop/sku/shopmenu/deleteSpecial"

type ShopmenuDeleteSpecialResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type ShopmenuDeleteSpecialRequest struct {
    /**
    *  门店 id 
    */
    VendorShopId string `json:"vendorShopId"`
    /**
    *  类目 id 列表 
    */
    VendorSellerCategoryIdList []string `json:"vendorSellerCategoryIdList"`
}



func (req *ShopmenuDeleteSpecialRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ShopmenuDeleteSpecialResponse, error) {
    resp, err := client.InvokeApi(shopmenu_delete_special_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ShopmenuDeleteSpecialResponse
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

func (response *ShopmenuDeleteSpecialResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

