package tuangouproductqueryproduct

/**
* 查询门店消费码商品信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const tuangou_product_queryproduct_url = "/ddzh/tuangou/product/queryproduct"

type TuangouProductQueryproductRequest struct {
    /**
    *  页码 
    */
    PageNo int32 `json:"pageNo"`
    /**
    *  单页大小 
    */
    PageSize int32 `json:"pageSize"`
}
type Data struct {
    Result []ProductItem `json:"result"`
}
type TuangouProductQueryproductResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type ProductItem struct {
    /**
    * 商品ID
    */
    ProductItemId int32 `json:"productItemId"`
    /**
    * 产品ID
    */
    ProductId int32 `json:"productId"`
    /**
    * 商品名称
    */
    Name string `json:"name"`
    /**
    * 价格
    */
    Price string `json:"price"`
    /**
    * 市场价格
    */
    MarketPrice string `json:"marketPrice"`
    /**
    * 商品类型，1081: 次卡
    */
    Type int32 `json:"type"`
}



func (req *TuangouProductQueryproductRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*TuangouProductQueryproductResponse, error) {
    resp, err := client.InvokeApi(tuangou_product_queryproduct_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response TuangouProductQueryproductResponse
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

func (response *TuangouProductQueryproductResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

