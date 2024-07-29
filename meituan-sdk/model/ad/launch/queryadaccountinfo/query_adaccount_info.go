package queryadaccountinfo

/**
* 查询账号基本信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_adaccount_info_url = "/ad/launch/queryAdAccountInfo"

type QueryAdaccountInfoRequest struct {
    SonAdaccount int32 `json:"sonAdaccount"`
}
type QueryAdaccountInfoResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 商家账号id，登陆名，签约门店，业务，城市
    */
    Data AdAccountInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type AdAccountInfo struct {
    /**
    * 商家账号
    */
    AdAccountId int32 `json:"adAccountId"`
    /**
    * 登陆名
    */
    Name string `json:"name"`
    /**
    * 签约门店
    */
    ShopId int64 `json:"shopId"`
    /**
    * 业务
    */
    Bu int32 `json:"bu"`
    /**
    * 城市
    */
    City int32 `json:"city"`
    /**
    * 门店名称
    */
    ShopName string `json:"shopName"`
}



func (req *QueryAdaccountInfoRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryAdaccountInfoResponse, error) {
    resp, err := client.InvokeApi(query_adaccount_info_url, 22, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryAdaccountInfoResponse
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

func (response *QueryAdaccountInfoResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

