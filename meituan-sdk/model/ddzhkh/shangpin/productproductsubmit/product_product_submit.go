package productproductsubmit

/**
* 提交商品
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const product_product_submit_url = "/ddzhkh/shangpin/product/submit"

type ProductProductSubmitResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Data struct {
    /**
    * 美团商品ID。提交商品成功后，会异步返回商品流程状态变更消息，可能会因为审核等原因上架失败
    */
    ProductId int64 `json:"productId"`
}
type ProductProductSubmitRequest struct {
    /**
    *  美团商品ID。提交商品成功后，会异步返回商品流程状态变更消息，可能会因为审核等原因上架失败 
    */
    ProductId int64 `json:"productId"`
}



func (req *ProductProductSubmitRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ProductProductSubmitResponse, error) {
    resp, err := client.InvokeApi(product_product_submit_url, 59, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ProductProductSubmitResponse
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

func (response *ProductProductSubmitResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

