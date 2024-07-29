package deleteshopposter

/**
* 商家开放平台删除海报
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const delete_shop_poster_url = "/waimai/ng/decoration/deleteShopPoster"

type DeleteShopPosterResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 操作结果
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type DeleteShopPosterRequest struct {
    /**
    *  海报id 
    */
    PosterId int64 `json:"posterId"`
}



func (req *DeleteShopPosterRequest) DoInvoke(client mtclient.MeituanClient) (*DeleteShopPosterResponse, error) {
    resp, err := client.InvokeApi(delete_shop_poster_url, 2, "", req)

    if err != nil {
        return nil, err
    }
    var response DeleteShopPosterResponse
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

func (response *DeleteShopPosterResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

