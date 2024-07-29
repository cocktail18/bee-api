package personquery

/**
* 查询绑定在门店下的服务人员信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const personquery_url = "/ddzh/yuding/person/query"

type PersonqueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Technician struct {
    /**
    * 服务人员id
    */
    ServiceTechId int64 `json:"serviceTechId"`
    /**
    * 人员姓名
    */
    Name string `json:"name"`
    /**
    * 身份证
    */
    IdCard string `json:"idCard"`
    /**
    * 手机号
    */
    Phone string `json:"phone"`
}
type Data struct {
    /**
    * 记录总数
    */
    Total int64 `json:"total"`
    /**
    * 人员信息列表 注：类型为对象类型，详见扩展参数technicians
    */
    Technicians []Technician `json:"technicians"`
}
type PersonqueryRequest struct {
    /**
    *  分页查询页码，从1开始 
    */
    Page int64 `json:"page"`
    /**
    *  分页查询页大小，最大为100 
    */
    Size int32 `json:"size"`
}



func (req *PersonqueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*PersonqueryResponse, error) {
    resp, err := client.InvokeApi(personquery_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response PersonqueryResponse
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

func (response *PersonqueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

