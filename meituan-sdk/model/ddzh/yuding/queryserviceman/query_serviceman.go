package queryserviceman

/**
* 查询服务人员信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_serviceman_url = "/ddzh/yuding/query/serviceman"

type ServiceManInfoDTO struct {
    /**
    * 是否实名制认证，0-未实名，1-已实名
    */
    RealNameAuthFlag int32 `json:"realNameAuthFlag"`
}
type QueryServicemanRequest struct {
    /**
    *  手机号 
    */
    Mobile string `json:"mobile"`
    /**
    *  服务人员姓名 
    */
    Name string `json:"name"`
}
type QueryServicemanResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ServiceManInfoDTO `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *QueryServicemanRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryServicemanResponse, error) {
    resp, err := client.InvokeApi(query_serviceman_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryServicemanResponse
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

func (response *QueryServicemanResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

