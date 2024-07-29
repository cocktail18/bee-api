package specialfoodbindspuandskucode

/**
* 绑定SPUcode以及SKUcode
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const special_food_bind_spu_and_sku_code_url = "/waimai/ng/special/food/bindSpuAndSkuCode"

type SpecialFoodBindSpuAndSkuCodeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type SpecialFoodBindSpuAndSkuCodeRequest struct {
    /**
    * 门店id 
    */
    EpoiId string `json:"epoiId"`
    /**
    * 业务标识，1-特价版
    */
    BusinessIdentify int32 `json:"businessIdentify"`
    /**
    * 方菜品id，当eDishCode为default时，表示清空该值，当eDishCode为空字符串时，表示不处理该字段
    */
    EDishCode string `json:"eDishCode"`
    /**
    *  美团SPUID
    */
    MtSpuId int64 `json:"mtSpuId"`
    /**
    * 三方规格ID，当sku_id为default时，表示清空该值，当sku_id为空字符串时，表示不处理该字段
    */
    SkuId string `json:"skuId"`
    /**
    * 美团SKUID
    */
    MtSkuId int64 `json:"mtSkuId"`
}



func (req *SpecialFoodBindSpuAndSkuCodeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*SpecialFoodBindSpuAndSkuCodeResponse, error) {
    resp, err := client.InvokeApi(special_food_bind_spu_and_sku_code_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response SpecialFoodBindSpuAndSkuCodeResponse
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

func (response *SpecialFoodBindSpuAndSkuCodeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

