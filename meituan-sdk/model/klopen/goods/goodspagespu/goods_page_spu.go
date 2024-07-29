package goodspagespu

/**
* 分页查询spu
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const goods_page_spu_url = "/kl/open/goods/page/spu"

type GoodsPageSpuRequest struct {
    MinIdNotInclude int64 `json:"minIdNotInclude"`
    Size int32 `json:"size"`
    BrandIdentify string `json:"brandIdentify"`
}
type GoodsPageSpuResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Data struct {
    MaxId string `json:"maxId"`
    List string `json:"list"`
}



func (req *GoodsPageSpuRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GoodsPageSpuResponse, error) {
    resp, err := client.InvokeApi(goods_page_spu_url, 27, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GoodsPageSpuResponse
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

func (response *GoodsPageSpuResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

