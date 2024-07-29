package tuangouproductqueryproductbytype

/**
* 按商品类型查询门店商品信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const tuangou_product_queryproductbytype_url = "/ddzh/tuangou/product/queryproductbytype"

type TuangouProductQueryproductbytypeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type ProductItemDTO struct {
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
    * 关联的服务项目，目前只有次卡类型的商品有该字段
    */
    ServiceSkuList []ServiceSkuDTO `json:"serviceSkuList"`
}
type TuangouProductQueryproductbytypeRequest struct {
    /**
    *  页码 
    */
    PageNo int32 `json:"pageNo"`
    /**
    *  单页大小，最大限制100 
    */
    PageSize int32 `json:"pageSize"`
    /**
    *  商品类型。500-款式一口价，723-运动健身拼团，722-足疗按摩拼团，792-洗浴拼团，793-密室拼团，795-体检拼团，850-医美预付，922-丽人拼团，923-生活服务拼团，924-教育培训拼团，926-K歌拼团，1081-次卡，1359-结婚代金券，2000：医美预付拼团。若不确定具体商品类型，可通过在线咨询 
    */
    Type int64 `json:"type"`
}
type Data struct {
    Result []ProductDTO `json:"result"`
}
type ServiceSkuDTO struct {
    /**
    * 服务项目ID，核销次卡时验券接口返回的productItemId对应此skuId
    */
    SkuId int32 `json:"skuId"`
    /**
    * 服务项目名称
    */
    SkuName string `json:"skuName"`
    /**
    * 数量
    */
    Quantity int32 `json:"quantity"`
    /**
    * 价格
    */
    Price string `json:"price"`
}
type ProductDTO struct {
    /**
    * 产品ID
    */
    ProductId int32 `json:"productId"`
    /**
    * 商品类型,次卡:1081,结婚代金券:1359,丽人拼团:922,教育培训拼团:924,生活服务拼团:923,洗浴拼团:792,款式一口价:500,体检拼团:795,足疗按摩拼团:722 若不确定具体商品类型，可通过在线咨询提问
    */
    Type int32 `json:"type"`
    /**
    * sku列表
    */
    ProductItemList []ProductItemDTO `json:"productItemList"`
}



func (req *TuangouProductQueryproductbytypeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*TuangouProductQueryproductbytypeResponse, error) {
    resp, err := client.InvokeApi(tuangou_product_queryproductbytype_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response TuangouProductQueryproductbytypeResponse
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

func (response *TuangouProductQueryproductbytypeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

