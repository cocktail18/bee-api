package specialfoodbatchquery

/**
* 批量查询商品
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const special_food_batch_query_url = "/waimai/ng/special/food/batchQuery"

type SpecialFoodBatchQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type WmOpenProductSpus struct {
    MtSpuId int64 `json:"mt_spu_id"`
    Name string `json:"name"`
    AppFoodCode string `json:"app_food_code"`
    Picture string `json:"picture"`
    TagName string `json:"tag_name"`
    MtTagId int64 `json:"mt_tag_id"`
    Description string `json:"description"`
    Tag string `json:"tag"`
    Skus []Skus `json:"skus"`
    Propertys []Propertys `json:"propertys"`
    Setmeals []Setmeals `json:"setmeals"`
    SpuType int64 `json:"spu_type"`
    Sequence int64 `json:"sequence"`
    Ctime int64 `json:"ctime"`
    Utime int64 `json:"utime"`
    SellStatus int64 `json:"sell_status"`
    MaxNum int64 `json:"max_num"`
}
type Propertys struct {
    PropertyName string `json:"property_name"`
    Values []int64 `json:"values"`
    RequiredNum int64 `json:"required_num"`
}
type Skus struct {
    MtSkuId int64 `json:"mt_sku_id"`
    SkuId string `json:"sku_id"`
    Spec string `json:"spec"`
    Price float64 `json:"price"`
    Stock int64 `json:"stock"`
    MaxStock int64 `json:"max_stock"`
}
type Setmeals struct {
    Name string `json:"name"`
    Price float64 `json:"price"`
}
type SpecialFoodBatchQueryRequest struct {
    Biz string `json:"biz"`
}
type Data struct {
    Total int64 `json:"total"`
    WmOpenProductSpus []WmOpenProductSpus `json:"wmOpenProductSpus"`
}



func (req *SpecialFoodBatchQueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*SpecialFoodBatchQueryResponse, error) {
    resp, err := client.InvokeApi(special_food_batch_query_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response SpecialFoodBatchQueryResponse
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

func (response *SpecialFoodBatchQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

