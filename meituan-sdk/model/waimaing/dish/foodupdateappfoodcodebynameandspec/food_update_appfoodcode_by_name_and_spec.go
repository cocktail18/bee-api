package foodupdateappfoodcodebynameandspec

/**
* 根据商品名称和规格名称更换新的商品编码
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const food_update_appfoodcode_by_name_and_spec_url = "/waimai/ng/dish/food/updateAppFoodCodeByNameAndSpec"

type FoodUpdateAppfoodcodeByNameAndSpecResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type FoodUpdateAppfoodcodeByNameAndSpecRequest struct {
    /**
    *  商品名称， 如变更商品编码，商品名称需同线上完全一致， 如不一致将无法匹配 
    */
    Name string `json:"name"`
    /**
    *  分类名称， 如变更商品编码，分类名称需同线上完全一致，如不一致将无法匹配 
    */
    CategoryName string `json:"categoryName"`
    /**
    *  新appFoodCode， 为商品的服务商方的商品id，不同门店可以重复，同一门店内不能重复，最大长度128 
    */
    AppFoodCode string `json:"appFoodCode"`
    /**
    *  新skuId，为商品sku的唯一标示,如商品有规格，则必填 
    */
    SkuId string `json:"skuId"`
    /**
    *  规格， 如变更商品编码，规格需同线上完全一致，如不一致将无法匹配,如商品有规格，则必填 
    */
    Spec string `json:"spec"`
}



func (req *FoodUpdateAppfoodcodeByNameAndSpecRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*FoodUpdateAppfoodcodeByNameAndSpecResponse, error) {
    resp, err := client.InvokeApi(food_update_appfoodcode_by_name_and_spec_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response FoodUpdateAppfoodcodeByNameAndSpecResponse
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

func (response *FoodUpdateAppfoodcodeByNameAndSpecResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

