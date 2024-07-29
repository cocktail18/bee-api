package queryadaccountsons

/**
* 查询账号的子账号
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_adaccount_sons_url = "/ad/launch/queryAdAccountSons"

type QueryAdaccountSonsRequest struct {
    SonAdaccount int32 `json:"sonAdaccount"`
}
type AdAccountInfo struct {
}
type QueryAdaccountSonsResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 商家账号id，登陆名，签约门店，业务，城市
    */
    Data []AdAccountInfo `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *QueryAdaccountSonsRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryAdaccountSonsResponse, error) {
    resp, err := client.InvokeApi(query_adaccount_sons_url, 22, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryAdaccountSonsResponse
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

func (response *QueryAdaccountSonsResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

