package personcreate

/**
* 创建服务人员
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const personcreate_url = "/ddzh/yuding/person/create"

type PersoncreateRequest struct {
    /**
    *  修改信息的服务人员id，在美团侧通过接口生成的服务人员id 新建时指定为0 
    */
    ServiceTechId int64 `json:"serviceTechId"`
    /**
    *  身份证号码，新建时必传 
    */
    IdCard string `json:"idCard"`
    /**
    *  人员姓名，新建时必传 
    */
    Name string `json:"name"`
    /**
    *  手机号，新建和修改手机号时必传 
    */
    Phone string `json:"phone"`
}
type Data struct {
    /**
    * 服务人员名称
    */
    Name string `json:"name"`
    /**
    * 服务人员id
    */
    ServiceTechId int64 `json:"serviceTechId"`
}
type PersoncreateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *PersoncreateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*PersoncreateResponse, error) {
    resp, err := client.InvokeApi(personcreate_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response PersoncreateResponse
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

func (response *PersoncreateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

