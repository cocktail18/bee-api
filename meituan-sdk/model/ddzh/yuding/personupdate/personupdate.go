package personupdate

/**
* 修改服务人员信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const personupdate_url = "/ddzh/yuding/person/update"

type OpenTechDTO struct {
    /**
    * 服务人员id
    */
    ServiceTechId int64 `json:"serviceTechId"`
    /**
    * 人员姓名
    */
    Name string `json:"name"`
    /**
    * 修改后手机号
    */
    Phone string `json:"phone"`
}
type PersonupdateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data OpenTechDTO `json:"data"`
    TraceId string `json:"traceId"`
}
type PersonupdateRequest struct {
    /**
    *  修改信息的服务人员id，在美团侧通过接口生成的服务人员id 
    */
    ServiceTechId int64 `json:"serviceTechId"`
    /**
    *  手机号，新建和修改手机号时必传 
    */
    Phone string `json:"phone"`
}



func (req *PersonupdateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*PersonupdateResponse, error) {
    resp, err := client.InvokeApi(personupdate_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response PersonupdateResponse
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

func (response *PersonupdateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

