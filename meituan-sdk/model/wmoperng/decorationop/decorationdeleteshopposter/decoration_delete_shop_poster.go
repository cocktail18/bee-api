package decorationdeleteshopposter

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

const decoration_delete_shop_poster_url = "/wmoper/ng/decorationop/deleteShopPoster"

type DecorationDeleteShopPosterResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 操作结果
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type DecorationDeleteShopPosterRequest struct {
    /**
    *  海报id 
    */
    PosterId int64 `json:"poster_id"`
}



func (req *DecorationDeleteShopPosterRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DecorationDeleteShopPosterResponse, error) {
    resp, err := client.InvokeApi(decoration_delete_shop_poster_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DecorationDeleteShopPosterResponse
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

func (response *DecorationDeleteShopPosterResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

