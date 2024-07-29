package ugcquerystar

/**
* 查询门店星级和单项分
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const ugc_querystar_url = "/ddzh/ugc/querystar"

type RqueryShopStarRes struct {
    /**
    * 点评平台和美团平台一致，50代表5星
    */
    ShopStar int32 `json:"shopStar"`
    /**
    * 单项分,部分类目为空
    */
    SubScore []SubScore `json:"subScore"`
}
type UgcQuerystarRequest struct {
    /**
    *  1-大众点评，2-美团（美团的评价与大众点评的评价是分开的，需要单独查） 
    */
    Platform int32 `json:"platform"`
}
type UgcQuerystarResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data RqueryShopStarRes `json:"data"`
    TraceId string `json:"traceId"`
}
type SubScore struct {
    Title string `json:"title"`
    Score float64 `json:"score"`
}



func (req *UgcQuerystarRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*UgcQuerystarResponse, error) {
    resp, err := client.InvokeApi(ugc_querystar_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response UgcQuerystarResponse
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

func (response *UgcQuerystarResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

