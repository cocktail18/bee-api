package personunbind

/**
* 解绑服务人员与门店关系
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const personunbind_url = "/ddzh/yuding/person/unbind"

type PersonunbindResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Data struct {
    /**
    * 服务人员id
    */
    ServiceTechId int64 `json:"serviceTechId"`
}
type PersonunbindRequest struct {
    /**
    *  要解绑的服务人员id，在美团侧通过接口生成的服务人员id 
    */
    ServiceTechId int64 `json:"serviceTechId"`
}



func (req *PersonunbindRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*PersonunbindResponse, error) {
    resp, err := client.InvokeApi(personunbind_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response PersonunbindResponse
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

func (response *PersonunbindResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

