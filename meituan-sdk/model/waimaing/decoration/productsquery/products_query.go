package productsquery

/**
* 商家开放平台商品查询
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const products_query_url = "/waimai/ng/decoration/productsQuery"

type Products struct {
    /**
    * ERP 商品id
    */
    AppFoodCode string `json:"app_food_code"`
    /**
    * 商品名称
    */
    FoodName string `json:"food_name"`
    /**
    * 图片url
    */
    Picture string `json:"picture"`
    /**
    * 是否可选
    */
    Selectable bool `json:"selectable"`
}
type ProductsQueryData struct {
    /**
    * 海报id
    */
    Total int64 `json:"total"`
    Products []Products `json:"products"`
}
type ProductsQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ProductsQueryData `json:"data"`
    TraceId string `json:"traceId"`
}
type ProductsQueryRequest struct {
    /**
    *  场景类型： 1.海报 2.老板推荐 
    */
    Scene int32 `json:"scene"`
    /**
    *  页码,默认-1 
    */
    PageNum int32 `json:"pageNum"`
    /**
    *  每页展示数 
    */
    PageSize int32 `json:"pageSize"`
}



func (req *ProductsQueryRequest) DoInvoke(client mtclient.MeituanClient) (*ProductsQueryResponse, error) {
    resp, err := client.InvokeApi(products_query_url, 2, "", req)

    if err != nil {
        return nil, err
    }
    var response ProductsQueryResponse
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

func (response *ProductsQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

